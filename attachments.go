package gitlink

import (
	"net/http"
	"os"
)

type AttachmentsService struct {
	client *Client
}

type AttachmentsData struct {
	Status   int    `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
	Id       int64  `json:"id,omitempty"`
	Filesize int64  `json:"filesize,omitempty"`
}

// PostAttachments 上传文件 https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-128323479
func (s *AttachmentsService) PostAttachments(filePath string) (*AttachmentsData, *Response, error) {

	u := "attachments.json"

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	req, err := s.client.UploadRequest(http.MethodPost, u, file, "file")
	if err != nil {
		return nil, nil, err
	}

	var attachmentsData *AttachmentsData
	resp, err := s.client.Do(req, &attachmentsData)
	if err != nil {
		return nil, resp, err
	}

	return attachmentsData, resp, nil
}
