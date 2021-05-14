package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func (files) Get(client *Client, fileId string) (*FileData, ApiErrors) {
	path := fmt.Sprintf("/v2/files/%s", fileId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var files FileData
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, FromError(err)
	}

	return &files, nil
}

// GetAll fetches all files
func (files) GetAll(client *Client) (*FileList, ApiErrors) {
	path := fmt.Sprintf("/v2/files")

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var files FileList
	if err := json.Unmarshal(body, &files); err != nil {
		return nil, FromError(err)
	}

	return &files, nil
}

// Create creates a file
func (files) CreateFromFile(client *Client, filename string, public bool, reader io.Reader) (*FileData, ApiErrors) {

	path := fmt.Sprintf("/v2/files")

	//prepare the reader instances to encode
	values := map[string]string{
		"public": strconv.FormatBool(public),
	}

	fileContents, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, FromError(err)
	}

	byteBuf, contentType, err := EncodeForm(values, filename, "file", fileContents)

	if err != nil {
		return nil, FromError(err)
	}

	body, apiError := client.DoFileRequest(path, byteBuf, contentType)
	if apiError != nil {
		return nil, apiError
	}
	var newFile FileData
	if err := json.Unmarshal(body, &newFile); err != nil {
		return nil, FromError(err)
	}

	return &newFile, nil
}

// Create creates a file
func (files) CreateFromFileLocation(client *Client, fileLocation string) (*FileData, ApiErrors) {

	path := fmt.Sprintf("/v2/files")

	//prepare the reader instances to encode
	values := map[string]string{
		"file_location": fileLocation,
	}

	byteBuf, contentType, err := EncodeForm(values, "", "", []byte{})

	if err != nil {
		return nil, FromError(err)
	}

	body, apiError := client.DoFileRequest(path, byteBuf, contentType)
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
func (files) Delete(client *Client, fileID string) ApiErrors {
	path := fmt.Sprintf("/v2/files/%s", fileID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a file.
func (files) Update(client *Client, fileID string, file *File) (*FileData, ApiErrors) {

	fileData := FileData{
		Data: *file,
	}

	jsonPayload, err := json.Marshal(fileData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/files/%s", fileID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedFile FileData
	if err := json.Unmarshal(body, &updatedFile); err != nil {
		return nil, FromError(err)
	}

	return &updatedFile, nil
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
