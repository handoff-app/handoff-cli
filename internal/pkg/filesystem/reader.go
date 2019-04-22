package filesystem

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// I got this code from somewhere, but I'm not sure where! Will apply attribution when I find it.
// TODO: Refactor to stream upload
func ReadToFormFile(path string, fieldName string) (io.Reader, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(part, file)

	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}