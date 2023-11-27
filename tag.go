package gitlink

import (
	"fmt"
	"net/http"
)

type DeleteTagRequest struct {
	Owner string // 必需
	Repo  string // 必需
	Tag   string // 标签名称、必需
}

type DeleteTag struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type TagService struct {
	client *Client
}

// DeleteTag 删除一个标签 https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-118749620
func (s *TagService) DeleteTag(request *DeleteTagRequest) (*DeleteTag, *Response, error) {

	u := fmt.Sprintf("/v1/%s/%s/tags/%s.json", request.Owner, request.Repo, request.Tag)

	req, err := s.client.NewRequest(http.MethodDelete, u, request)
	if err != nil {
		return nil, nil, err
	}

	var deleteTag *DeleteTag
	resp, err := s.client.Do(req, &deleteTag)
	if err != nil {
		return nil, resp, err
	}

	return deleteTag, resp, nil
}
