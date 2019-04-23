package files

import (
	"bytes"
	"crypto/md5"
	"io"
	"testing"
)

type writeWithMetadataTest struct {
	reader   io.Reader
	checksum string
	mime     string
}

func TestWriteWithMetadata(t *testing.T) {
	magicNumber := []byte{0xFF, 0xD8, 0xFF}
	shouldWrite := int64(len(magicNumber))
	rawChecksum := md5.Sum(magicNumber)
	checksum := string(rawChecksum[0:16])
	table := []writeWithMetadataTest{
		{bytes.NewReader([]byte{0xFF, 0xD8, 0xFF}), checksum, "image/jpeg"},
		{bytes.NewReader([]byte{0xFF, 0xD8, 0xFF}), checksum, ""},
		{bytes.NewReader([]byte{0xFF, 0xD8, 0xFF}), "", "image/jpeg"},
		{bytes.NewReader([]byte{0xFF, 0xD8, 0xFF}), "", ""},
	}
	for _, test := range table {
		var flags FileMetadata
		byteArr := make([]byte, 3)
		buf := bytes.NewBuffer(byteArr)
		if test.checksum != "" {
			flags = flags | Checksum
		}
		if test.mime != "" {
			flags = flags | MIME
		}
		wi, err := WriteWithMetadata(buf, test.reader, flags)
		if err != nil {
			t.Fatalf("failed to execute WriteWithMetadata: %s\n", err)
		}
		if wi.Written != shouldWrite {
			t.Fatalf("expected to write %d but wrote %d\n", shouldWrite, wi.Written)
		}
		if test.checksum == "" && len(wi.Checksum) > 0 {
			t.Fatalf("expected checksum to not be calculated but was %s\n", wi.Checksum)
		} else if test.checksum != "" && len(wi.Checksum) == 0 {
			t.Fatalf("expected checksum be calculated but was %s\n", wi.Checksum)
		} else if test.checksum != string(wi.Checksum) {
			t.Fatalf("expected checksum %s but was %s\n", test.checksum, wi.Checksum)
		}
		if test.mime == "" && len(wi.MimeType) > 0 {
			t.Fatalf("expected mime to not be detected but was %s\n", wi.MimeType)
		} else if test.mime != "" && len(wi.MimeType) == 0 {
			t.Fatalf("expected mime be detected but was %s\n", wi.MimeType)
		} else if test.mime != string(wi.MimeType) {
			t.Fatalf("expected mime %s but was %s\n", test.mime, wi.MimeType)
		}
	}
}
