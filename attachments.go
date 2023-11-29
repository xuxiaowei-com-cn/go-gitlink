package gitlink

import (
	"net/http"
	"os"
	"path/filepath"
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
//
// 接口需要凭证
//
// filePath 文件路径，包含文件名
//
// fileName 上传到服务器上的文件名，如果为空，则截取 filePath 的文件名
func (s *AttachmentsService) PostAttachments(filePath string, fileName string) (*AttachmentsData, *Response, error) {

	u := "attachments.json"

	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	if fileName == "" {
		fileName = filepath.Base(filePath)
	}

	req, err := s.client.UploadRequest(http.MethodPost, u, file, fileName)
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
