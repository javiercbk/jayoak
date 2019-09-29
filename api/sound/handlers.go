package sound

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/volatiletech/null"

	"github.com/ericlagergren/decimal"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"

	"github.com/javiercbk/jayoak/api/sound/spectrum"
	"github.com/javiercbk/jayoak/files"
	"github.com/javiercbk/jayoak/http/response"
	"github.com/javiercbk/jayoak/http/session"
	"github.com/javiercbk/jayoak/models"
)

// ErrInstrumentNotFound is an error thrown when an instrument was not found in the database
var ErrInstrumentNotFound = errors.New("instrument was not found in the database")

// ErrUnexpectedRowsAffectedCount is an error thrown when updating a sound and rows affected count is not expected
var ErrUnexpectedRowsAffectedCount = errors.New("unexpected rows affected count")

// soundProcessingStrategy is a function that process an audio files and stores the frequencies in the database
type soundProcessingStrategy func(userID, instrumentID, soundUUID, extension string, soundID int64) error

type open interface {
	Open() (multipart.File, error)
}

type soundSearch struct {
	ID int64 `uri:"soundID" binding:"required,numeric"`
}

type prospectSound struct {
	Name          *string `json:"name,omitempty" binding:"length(1,256),optional"`
	AudioFileName string  `json:"audioFileName" binding:"length(1,256)"`
	FileFactory   open    `json:"-"`
	InstrumentID  *int64  `json:"instrumentId,omitempty" binding:"numeric,optional"`
	Note          *string `json:"note,omitempty" binding:"in(a,b,c,d,e,f,g,a#,b#,c#,d#,e#,f#,g#,ab,bb,cb,db,eb,fb,gb),optional"`
}

// Handlers defines a handler for sound
type Handlers struct {
	logger     *log.Logger
	db         *sql.DB
	repository files.Repository
}

// NewHandlers creates a new handler for sound
func NewHandlers(logger *log.Logger, db *sql.DB, repository files.Repository) *Handlers {
	return &Handlers{
		logger:     logger,
		db:         db,
		repository: repository,
	}
}

// FindSound find an audio clip in the database
func (h *Handlers) Retrieve(c *gin.Context) {
	var search soundSearch
	user := session.RetrieveUser(c)
	err := c.ShouldBind(&search)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "error binding sound search params")
		return
	}
	sound, err := models.Sounds(qm.Where("id = ? AND organization_id", search.ID, user.OrganizationID)).One(c, h.db)
	if err != nil {
		h.logger.Printf("error retriving sound %d: %v\n", search.ID, err)
		response.NewErrorResponse(c, http.StatusInternalServerError, "error retrieving sound from the database")
		return
	}
	response.NewSuccessResponse(c, *sound)
}

// PostSound creates an audio, analizes it and stores it in the database
func (h *Handlers) PostSound(c *gin.Context) {
	var pSound prospectSound
	user := session.RetrieveUser(c)
	err := c.ShouldBindJSON(&pSound)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if pSound.InstrumentID == nil && pSound.Note != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "note was provided with no instrument")
		return
	}
	if pSound.InstrumentID != nil && pSound.Note == nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "instrument was provided with no note")
		return
	}
	file, err := c.FormFile("audio")
	if err != nil {
		h.logger.Printf("error retrieving uploaded file: %s\n", err)
		response.NewErrorResponse(c, http.StatusBadRequest, "no sound file provided")
		return
	}
	pSound.FileFactory = file
	sound, err := h.CreateSound(c, user, pSound, h.processSound)
	if err != nil {
		h.logger.Printf("error creating sound: %s\n", err)
		response.NewErrorResponse(c, http.StatusBadRequest, "no sound file provided")
		return
	}
	response.NewSuccessResponse(c, *sound)
}

// Routes initializes the routes for the audio handlers
func (h *Handlers) Routes(rg *gin.RouterGroup) {
	rg.GET("/sound/:soundId", h.Retrieve)
	rg.POST("/sound", h.PostSound)
}

// non HTTP Handlers

// CreateSound creates a sound with a validated prospect audio.
func (h *Handlers) CreateSound(ctx context.Context, user session.User, pSound prospectSound, strategy soundProcessingStrategy) (*models.Sound, error) {
	var instrumentStr string
	isNote := false
	extension := path.Ext(pSound.AudioFileName)
	if pSound.InstrumentID != nil {
		isNote = true
		instrumentStr = strconv.FormatInt(*pSound.InstrumentID, 10)
	}
	h.logger.Println("will write audio file")
	file, err := pSound.FileFactory.Open()
	if err != nil {
		h.logger.Printf("error opening uploaded file: %s\n", err)
	}
	defer file.Close()
	soundUUID := uuid.NewV4().String()
	soundFile, err := h.repository.SoundFile(strconv.FormatInt(user.ID, 10), instrumentStr, soundUUID, extension, os.O_WRONLY|os.O_CREATE|os.O_TRUNC)
	if err != nil {
		h.logger.Printf("error retrieving sound file: %s\n", err)
		return nil, err
	}
	wi, err := files.WriteWithMetadata(soundFile, file, files.Checksum|files.MIME)
	if err != nil {
		h.logger.Printf("error writing sound with metadata: %s\n", err)
		return nil, err
	}
	// we cannot defer this call in the case we have to delete the file
	// if the sound fails to be stored in the db
	err = soundFile.Close()
	if err != nil {
		h.logger.Printf("error closing sound file: %s\n", err)
		return nil, err
	}
	h.logger.Printf("%d bytes were written to the audio file with MIME %s and checksum %s\n", wi.Written, wi.MimeType, wi.Checksum)
	sound := &models.Sound{
		AudioFileName:  pSound.AudioFileName,
		MD5File:        fmt.Sprintf("%x", wi.Checksum),
		MimeType:       wi.MimeType,
		AudioUUID:      soundUUID,
		OrganizationID: user.OrganizationID,
		CreatorID:      user.ID,
	}
	if isNote {
		instrument, err := models.Instruments(qm.Where("id = ? AND organization_id = ?", pSound.InstrumentID, user.ID)).One(ctx, h.db)
		if err != nil {
			h.logger.Printf("error retrieving instrument %d for user %d: %s\n", pSound.InstrumentID, user.ID, err)
			return nil, err
		}
		if instrument == nil {
			return nil, ErrInstrumentNotFound
		}
		sound.InstrumentID = null.Int64FromPtr(pSound.InstrumentID)
		sound.Note = null.StringFromPtr(pSound.Note)
	} else {
		sound.Name = null.StringFromPtr(pSound.Name)
	}
	err = sound.Insert(ctx, h.db, boil.Infer())
	if err != nil {
		h.logger.Printf("error saving sound: %s\n", err)
		errFile := h.repository.RemoveSound(strconv.FormatInt(user.ID, 10), instrumentStr, soundUUID, extension)
		if errFile != nil {
			h.logger.Printf("error removing sound file after failing to save soun in the db: %s\n", errFile)
		}
		return nil, err
	}
	err = strategy(strconv.FormatInt(user.ID, 10), instrumentStr, soundUUID, extension, sound.ID)
	if err != nil {
		h.logger.Printf("error retrieving instrument %d for user %d: %s\n", pSound.InstrumentID, user.ID, err)
	}
	// if an error was thrown when processing the audio then return both the sound and the error
	// FIXME: implement a transaction commit here
	return sound, err
}

// ProcessSound finds a sound file and reads it to extract the frequencies with the intensity and saves it to the frequency table
func (h *Handlers) ProcessSound(userID, instrumentID, soundUUID, extension string, soundID int64) error {
	return h.processSound(userID, instrumentID, soundUUID, extension, soundID, false)
}

// ProcessSoundNoFK does the same that ProcessSound but disables the foreign key for faster insertion
func (h *Handlers) ProcessSoundNoFK(userID, instrumentID, soundUUID, extension string, soundID int64) error {
	return h.processSound(userID, instrumentID, soundUUID, extension, soundID, false)
}

func (h *Handlers) analyzeFrequencies(userID, instrumentID, soundUUID, extension string) (spectrum.FrequenciesAnalysis, error) {
	var freqAnalysis spectrum.FrequenciesAnalysis
	soundFile, err := h.repository.SoundFile(userID, instrumentID, soundUUID, extension, os.O_RDONLY)
	if err != nil {
		h.logger.Printf("error while processing. Could not retrieve sound file: %s\n", err)
		return freqAnalysis, err
	}
	defer soundFile.Close()
	pcm, sampleRate, err := spectrum.PCMFromWav(soundFile)
	if err != nil {
		h.logger.Printf("error extracting PCM from wav file %s: %s\n", soundUUID, err)
		return freqAnalysis, err
	}
	return spectrum.FrequenciesSpectrumAnalysis(pcm, sampleRate)
}

// processSound process a wav file and stores the frequencies in a separated table
func (h *Handlers) processSound(userID, instrumentID, soundUUID, extension string, soundID int64, disableForeignKey bool) error {
	ctx := context.Background()
	freqAnalysis, err := h.analyzeFrequencies(userID, instrumentID, soundUUID, extension)
	if err != nil {
		h.logger.Printf("error building spectrum for wav file %s: %s\n", soundUUID, err)
		return err
	}
	updateCols := models.M{
		models.SoundColumns.MaxFrequency:  freqAnalysis.MaxFreq,
		models.SoundColumns.MinFrequency:  freqAnalysis.MinFreq,
		models.SoundColumns.MaxPowerFreq:  freqAnalysis.MaxPower.Freq,
		models.SoundColumns.MaxPowerValue: freqAnalysis.MaxPower.Value,
	}
	rowsAff, err := models.Sounds(qm.Where("id = ?", soundID)).UpdateAll(ctx, h.db, updateCols)
	if err != nil {
		h.logger.Printf("error updating sound %s: %s\n", soundUUID, err)
		return err
	}
	if rowsAff != 1 {
		h.logger.Printf("unexpected rows affected count %d for sound %s\n", rowsAff, soundUUID)
		return ErrUnexpectedRowsAffectedCount
	}
	var bigInsert strings.Builder
	if disableForeignKey {
		fmt.Fprintf(&bigInsert, "ALTER TABLE frequencies DISABLE TRIGGER ALL;\n")
	}
	fmt.Fprintf(&bigInsert, "INSERT INTO frequencies (sound_id, frequency, spl) VALUES ")
	first := true
	for freq, val := range freqAnalysis.Spectrum {
		if first {
			first = false
		} else {
			bigInsert.WriteString(",")
		}
		fmt.Fprintf(&bigInsert, "(%d, %d, %f)", soundID, freq, val)
	}
	bigInsert.WriteString(";")
	if disableForeignKey {
		fmt.Fprintf(&bigInsert, "\nALTER TABLE frequencies ENABLE TRIGGER ALL;")
	}
	result, err := queries.Raw(bigInsert.String()).ExecContext(ctx, h.db)
	if err != nil {
		h.logger.Printf("error inserting all frequencies for sound %s: %s\n", soundUUID, err)
		return err
	}
	h.logger.Printf("inserted all frequencies for sound %s:%v\n", soundUUID, result)
	return nil
}
