package gitlink

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type AttachmentsService struct {
	client *Client
}

type AttachmentsData struct {
	Status      *int       `json:"status,omitempty"`
	Message     string     `json:"message,omitempty"`
	Id          string     `json:"id,omitempty"`
	Filesize    string     `json:"filesize,omitempty"`
	Title       string     `json:"title,omitempty"`
	IsPdf       *bool      `json:"is_pdf,omitempty"`
	Url         string     `json:"url,omitempty"`
	CreatedOn   *time.Time `json:"created_on,omitempty"`
	ContentType string     `json:"content_type,omitempty"`
}

func (m *AttachmentsData) UnmarshalJSON(data []byte) error {

	// 处理为 map
	var dataMap map[string]interface{}
	err := json.Unmarshal(data, &dataMap)
	if err != nil {
		return err
	}

	rawStatus, ok := dataMap["status"].(float64)
	if ok {
		rawStatusInt := int(rawStatus)
		m.Status = &rawStatusInt
	}

	rawMessage, ok := dataMap["message"].(string)
	if ok {
		m.Message = rawMessage
	}

	rawId, ok := dataMap["id"].(string)
	if ok {
		m.Id = rawId
	}

	rawFilesize, ok := dataMap["filesize"].(string)
	if ok {
		m.Filesize = rawFilesize
	}

	rawTitle, ok := dataMap["title"].(string)
	if ok {
		m.Title = rawTitle
	}

	rawIsPdf, ok := dataMap["is_pdf"].(bool)
	if ok {
		m.IsPdf = &rawIsPdf
	}

	rawUrl, ok := dataMap["url"].(string)
	if ok {
		m.Url = rawUrl
	}

	rawContentType, ok := dataMap["content_type"].(string)
	if ok {
		m.ContentType = rawContentType
	}

	rawCreatedOn, ok := dataMap["created_on"].(string)
	if ok {
		parsedTime, err := time.Parse("2006-01-02 15:04", rawCreatedOn)
		if err != nil {
			return err
		}

		m.CreatedOn = &parsedTime
	}

	return nil
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
