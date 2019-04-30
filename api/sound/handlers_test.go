package sound

import (
	"context"
	"database/sql"
	"fmt"
	"mime/multipart"
	"os"
	"strings"
	"testing"

	"github.com/ericlagergren/decimal"
	"github.com/javiercbk/jayoak/api/sound/spectrum"
	"github.com/javiercbk/jayoak/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"

	"github.com/javiercbk/jayoak/http/session"
	testHelper "github.com/javiercbk/jayoak/testing"
)

func TestMain(m *testing.M) {
	testHelper.InitializeDB(m)
}

type mockFile struct {
	path string
}

func (m mockFile) Open() (multipart.File, error) {
	return os.Open(m.path)
}

func setUp(ctx context.Context) (*Handlers, session.User, error) {
	nullLogger := testHelper.NullLogger()
	repository := testHelper.Repository()
	var sessUser session.User
	db, err := testHelper.DB()
	if err != nil {
		return nil, sessUser, err
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, sessUser, err
	}
	org := &models.Organization{
		Name: "Test",
	}
	err = org.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, sessUser, err
	}
	user := &models.User{
		Name:           "Test",
		Email:          "test@test.com",
		Password:       "unhashedPassword",
		OrganizationID: org.ID,
	}
	err = user.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, sessUser, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, sessUser, err
	}
	sessUser.ID = user.ID
	sessUser.Name = user.Name
	sessUser.OrganizationID = user.OrganizationID
	h := NewHandlers(nullLogger, db, repository)
	return h, sessUser, nil
}

func testProcessSoundStrategy(ctx context.Context, t testing.TB, h *Handlers, user session.User, strategy soundProcessingStrategy) int64 {
	var err error
	var sound *models.Sound
	mf := mockFile{
		path: "testdata/piano-c-1.wav",
	}
	name := "Piano C 1"
	pSound := prospectSound{
		AudioFileName: "piano-c-1.wav",
		FileFactory:   mf,
		InstrumentID:  nil,
		Name:          &name,
		Note:          nil,
	}
	sound, err = h.CreateSound(context.Background(), user, pSound, strategy)
	if err != nil {
		t.Fatalf("error processing sound %s\n", err)
	}
	if sound.ID <= 0 {
		t.Fatal("sound should have an ID ")
	}
	if sound.Name.IsZero() {
		t.Fatal("sound should have a name")
	}
	soundName := sound.Name.Ptr()
	if soundName == nil {
		t.Fatalf("expected sound name %s but was nil\n", name)
	}
	if *soundName != name {
		t.Fatalf("expected sound name %s but was %s\n", name, *soundName)
	}
	return sound.ID
}

func addFrequenciesArray(ctx context.Context, db *sql.DB, soundID int64, freqAnalysis spectrum.FrequenciesAnalysis) error {
	decimalArr := make(types.DecimalArray, len(freqAnalysis.Spectrum))
	for i, v := range freqAnalysis.Spectrum {
		d := decimal.New(0, 0)
		d.SetFloat64(v)
		decimalArr[i] = types.Decimal{Big: d}
	}
	updateCols := models.M{
		models.SoundColumns.MaxFrequency:   null.IntFrom(freqAnalysis.MaxFreq),
		models.SoundColumns.MinFrequency:   null.IntFrom(freqAnalysis.MinFreq),
		models.SoundColumns.MaxPowerFreq:   null.IntFrom(freqAnalysis.MaxPower.Freq),
		models.SoundColumns.MaxPowerValue:  null.Float64From(freqAnalysis.MaxPower.Value),
		models.SoundColumns.FrequenciesArr: decimalArr,
	}
	_, err := models.Sounds(qm.Where("id = ?", soundID)).UpdateAll(ctx, db, updateCols)
	return err
}

func addFrequenciesRows(ctx context.Context, db *sql.DB, soundID int64, freqAnalysis spectrum.FrequenciesAnalysis) error {
	var bigInsert strings.Builder
	updateCols := models.M{
		models.SoundColumns.MaxFrequency:  null.IntFrom(freqAnalysis.MaxFreq),
		models.SoundColumns.MinFrequency:  null.IntFrom(freqAnalysis.MinFreq),
		models.SoundColumns.MaxPowerFreq:  null.IntFrom(freqAnalysis.MaxPower.Freq),
		models.SoundColumns.MaxPowerValue: null.Float64From(freqAnalysis.MaxPower.Value),
	}
	_, err := models.Sounds(qm.Where("id = ?", soundID)).UpdateAll(ctx, db, updateCols)
	if err != nil {
		return err
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
	_, err = queries.Raw(bigInsert.String()).ExecContext(ctx, db)
	return err
}

func createTestSound() error {
	updateCols := models.M{
		models.SoundColumns.MaxFrequency:   null.IntFrom(freqAnalysis.MaxFreq),
		models.SoundColumns.MinFrequency:   null.IntFrom(freqAnalysis.MinFreq),
		models.SoundColumns.MaxPowerFreq:   null.IntFrom(freqAnalysis.MaxPower.Freq),
		models.SoundColumns.MaxPowerValue:  null.Float64From(freqAnalysis.MaxPower.Value),
		models.SoundColumns.FrequenciesArr: decimalArr,
	}
	rowsAff, err := models.Sounds(qm.Where("id = ?", soundID)).UpdateAll(ctx, h.db, updateCols)
	if err != nil {
		h.logger.Printf("error updating sound %s: %s\n", soundUUID, err)
		return err
	}
}

func TestProcessSound(t *testing.T) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		t.Fatalf("error setting up test: %s\n", err)
	}
	testProcessSoundStrategy(ctx, t, h, user, h.ProcessSound)
}

func TestProcessSoundNoFK(t *testing.T) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		t.Fatalf("error setting up test: %s\n", err)
	}
	testProcessSoundStrategy(ctx, t, h, user, h.ProcessSoundNoFK)
}

func TestProcessSoundArray(t *testing.T) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		t.Fatalf("error setting up test: %s\n", err)
	}
	testProcessSoundStrategy(ctx, t, h, user, h.ProcessSoundArray)
}

// BenchmarkProcessSound-12         	      20	4072970382 ns/op	29874251 B/op	  312960 allocs/op
func BenchmarkProcessSound(b *testing.B) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		b.Fatalf("error setting up test: %s\n", err)
	}
	// Reset the timer so that the benchmark does not consider values from
	// initialization
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testProcessSoundStrategy(ctx, b, h, user, h.ProcessSound)
	}
}

// BenchmarkProcessSoundNoFK-12     	      10	4705490930 ns/op	29870240 B/op	  312959 allocs/op
func BenchmarkProcessSoundNoFK(b *testing.B) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		b.Fatalf("error setting up test: %s\n", err)
	}
	// Reset the timer so that the benchmark does not consider values from
	// initialization
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testProcessSoundStrategy(ctx, b, h, user, h.ProcessSoundNoFK)
	}
}

// BenchmarkProcessSoundArray-12    	     100	 206125523 ns/op	43997902 B/op	  564153 allocs/op
func BenchmarkProcessSoundArray(b *testing.B) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		b.Fatalf("error setting up test: %s\n", err)
	}
	// Reset the timer so that the benchmark does not consider values from
	// initialization
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testProcessSoundStrategy(ctx, b, h, user, h.ProcessSoundArray)
	}
}

func BenchmarkFindFrequencies(b *testing.B) {
	ctx := context.Background()
	h, user, err := setUp(ctx)
	if err != nil {
		b.Fatalf("error setting up test: %s\n", err)
	}
	testProcessSoundStrategy(ctx, b, h, user, h.ProcessSound)
	b.ResetTimer()

}
