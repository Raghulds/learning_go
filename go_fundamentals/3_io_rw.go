package gofundamentals

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func GetFileSign() {
	// fmt.Println(sha1Sig("go.mod"))
	fmt.Println(sha1Sig("http.log.gz"))
}

// Returns Sha1 signature of an uncompressed file
// cat http.log.gz | gunzip | sha1sum
func sha1Sig(filename string) (string, error) {
	if filepath.Ext(filename) != ".gz" {
		return "", fmt.Errorf("%s - Not a valid file extension. Expecting gz file.", filename)
	}

	// cat http.log.gz
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// gunzip
	r, err := gzip.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("%q - gzip: %w", filename, err)
	}
	defer r.Close()

	// sha1sum
	w := sha1.New()
	// Copy from r and write to w
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", filename, err)
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
