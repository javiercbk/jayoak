package files

import (
	"crypto/md5"
	"hash"
	"io"
	"os"
	"path/filepath"

	"github.com/javiercbk/filetype"
	"github.com/javiercbk/filetype/types"
)

// ReadWriteSeekCloser is a Reader, a Writer, a Seeker and a Closer
type ReadWriteSeekCloser interface {
	io.Reader
	io.Writer
	io.Seeker
	io.Closer
}

// FileMetadata is a bit masking value to enumerate options to extract metadata from writers
type FileMetadata uint32

const (
	// Checksum extracts the file checksum from the reader
	Checksum FileMetadata = 1 << iota
	// MIME extracts the file MIME from the reader
	MIME
)

// Repository is a file management structure
type Repository struct {
	basePath string
}

// WriteInfo modes all the information that can be extracted when writing a Writer
type WriteInfo struct {
	Written  int64
	Checksum []byte
	MimeType string
}

// NewRepository creates a new file repository
func NewRepository(basePath string) Repository {
	return Repository{
		basePath: basePath,
	}
}

// BasePath return the base path of the file repository
func (r Repository) BasePath() string {
	return r.basePath
}

func (r Repository) soundFolderPath(userID, instrumentID string) string {
	return filepath.Join(r.basePath, userID, instrumentID)
}

func (r Repository) soundFilePath(userID, instrumentID, soundUUID, extension string) string {
	filename := soundUUID + extension
	return filepath.Join(r.soundFolderPath(userID, instrumentID), filename)
}

// SoundFile returns the sound file
func (r Repository) SoundFile(userID, instrumentID, soundUUID, extension string, flag int) (ReadWriteSeekCloser, error) {
	folderPath := r.soundFolderPath(userID, instrumentID)
	err := os.MkdirAll(folderPath, 0777)
	if err != nil {
		return nil, err
	}
	fileLocation := r.soundFilePath(userID, instrumentID, soundUUID, extension)
	return os.OpenFile(fileLocation, flag, 0644)
}

// RemoveSound removes a sound
func (r Repository) RemoveSound(userID, instrumentID, soundUUID, extension string) error {
	fileLocation := r.soundFilePath(userID, instrumentID, soundUUID, extension)
	return os.Remove(fileLocation)
}

// WriteWithMetadata writes a reader to a writer and extract the required metadata
func WriteWithMetadata(writer io.Writer, reader io.Reader, meta FileMetadata) (WriteInfo, error) {
	var h hash.Hash
	var t types.Type
	var mw *filetype.MatcherWriter
	wi := WriteInfo{}
	currentReader := reader
	if meta&Checksum != 0 {
		h = md5.New()
		teeReader := io.TeeReader(currentReader, h)
		currentReader = teeReader
	}
	if meta&MIME != 0 {
		mw = filetype.NewMatcherWriter()
		teeReader := io.TeeReader(currentReader, mw)
		currentReader = teeReader
	}
	written, err := io.Copy(writer, currentReader)
	wi.Written = written
	if err != nil {
		return wi, err
	}
	if meta&Checksum != 0 {
		wi.Checksum = h.Sum(nil)
	}
	if meta&MIME != 0 {
		// if error it will be returned on the last return
		t, err = mw.Match()
		wi.MimeType = t.MIME.Value
	}
	return wi, err
}
