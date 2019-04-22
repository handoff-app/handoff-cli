package api

import (
	"../filesystem"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type FilesClient struct {
	Client  http.Client
	BaseUri string
}

type UploadSuccessDto struct {
	DownloadUri string `json:"download_uri"`
	DeleteUri   string `json:"delete_uri"`
}

type UploadResponse struct {
	Data UploadSuccessDto `json:"data"`
}

type Uploader interface {
	Upload(filePath string) (*http.Response, error)
}

func createUploadRequest(uri string, payload io.Reader, contentType string) *http.Request {
	req, _ := http.NewRequest("POST", uri, payload)
	req.Header.Set("Content-Type", contentType)
	return req
}

func (fc FilesClient) Upload(filePath string) (*UploadResponse, error) {
	payload, contentType, err := filesystem.ReadToFormFile(filePath, "file")

	body := &UploadResponse{}

	if err != nil {
		return body, err
	}

	req := createUploadRequest(fc.BaseUri, payload, contentType)

	resp, err := fc.Client.Do(req)

	if err != nil {
		return body, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		return body, errors.New(fmt.Sprintf("Upload failed. Received status code %d expected 201", resp.StatusCode))
	}

	err = json.NewDecoder(resp.Body).Decode(body)

	if err != nil {
		return body, err
	}

	return body, nil
}
