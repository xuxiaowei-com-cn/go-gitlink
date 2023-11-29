package gitlink

import (
	"fmt"
	"net/http"
)

type PostReleasesRequestPath struct {
	Owner string // 必需
	Repo  string // 必需
}

type PostReleasesRequestBody struct {
	AttachmentIds   []int  `json:"attachment_ids"`   // 可选，附件ID数组
	Body            string `json:"body"`             // 必需，发行版描述
	Name            string `json:"name"`             // 必需，发行版标题
	TagName         string `json:"tag_name"`         // 必需，标签
	Draft           bool   `json:"draft"`            // 必需，是否为草稿
	TargetCommitish string `json:"target_commitish"` // 可选，分支
	Prerelease      bool   `json:"prerelease"`       // 可选，是否为预发行版本
}

type PostReleases struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type ReleasesService struct {
	client *Client
}

// PostReleases 创建发行版 https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-128319361
//
// 接口不需要凭证
func (s *ReleasesService) PostReleases(requestPath *PostReleasesRequestPath, requestBody *PostReleasesRequestBody) (*PostReleases, *Response, error) {

	u := fmt.Sprintf("%s/%s/releases", requestPath.Owner, requestPath.Repo)

	req, err := s.client.NewRequest(http.MethodPost, u, nil, requestBody)
	if err != nil {
		return nil, nil, err
	}

	var postReleases *PostReleases
	resp, err := s.client.Do(req, &postReleases)
	if err != nil {
		return nil, resp, err
	}

	return postReleases, resp, nil
}
