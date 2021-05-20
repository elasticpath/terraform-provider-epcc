package epcc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"path"
	"strconv"
)

var Files files

type files struct{}

type File struct {
	Id       string    `json:"id"`
	Type     string    `json:"type"`
	FileName string    `json:"file_name,omitempty"`
	Link     *FileLink `json:"link,omitempty"`
	MimeType string    `json:"mime_type,omitempty"`
	FileSize int       `json:"file_size,omitempty"`
	Public   bool      `json:"public"`
}

type FileLink struct {
	Href string `json:"href,omitempty"`
}

func (files) Get(ctx *context.Context, client *Client, fileId string) (*FileData, ApiErrors) {
	path := fmt.Sprintf("/v2/files/%s", fileId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var files FileData
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, FromError(err)
	}

	return &files, nil
}

// CreateFromFile creates a file
func (files) CreateFromFile(ctx *context.Context, client *Client, filePath string, public bool) (*FileData, ApiErrors) {

	myUrl, err := url.Parse(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fileName := path.Base(myUrl.Path)

	path := fmt.Sprintf("/v2/files")

	//prepare the reader instances to encode
	values := map[string]string{
		"public": strconv.FormatBool(public),
	}

	fileContents, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, FromError(err)
	}

	byteBuf, contentType, err := EncodeForm(values, fileName, "file", fileContents)

	if err != nil {
		return nil, FromError(err)
	}

	body, apiError := client.DoFileRequest(ctx, path, byteBuf, contentType)
	if apiError != nil {
		return nil, apiError
	}
	var newFile FileData
	if err := json.Unmarshal(body, &newFile); err != nil {
		return nil, FromError(err)
	}

	return &newFile, nil
}

// Delete deletes a file.
func (files) Delete(ctx *context.Context, client *Client, fileID string) ApiErrors {
	path := fmt.Sprintf("/v2/files/%s", fileID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

type FileData struct {
	Data File `json:"data"`
}

// FileMeta contains extra data for an file
type FileMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type FileDataList struct {
}

type FileList struct {
	Data []File
}
