package gitlink

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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

type GetTagsRequest struct {
	Owner string // 必需
	Repo  string // 必需
	Name  string // 搜索关键词，可选
	ListOptions
}

type Tagger struct {
	Id       int32  `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ImageUrl string `json:"image_url"`
}

type Committer struct {
	Id       int32  `json:"id"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ImageUrl string `json:"image_url"`
}

type Commit struct {
	Sha           string    `json:"sha"`
	Message       string    `json:"message"`
	TimeAgo       string    `json:"time_ago"`
	CreatedAtUnix int       `json:"created_at_unix"`
	Committer     Committer `json:"committer"`
	Author        Author    `json:"author"`
}

type Tag struct {
	Name          string `json:"name"`
	Id            string `json:"id"`
	ZipballUrl    string `json:"zipball_url"`
	TarballUrl    string `json:"tarball_url"`
	Tagger        Tagger `json:"tagger"`
	TimeAgo       string `json:"time_ago"`
	CreatedAtUnix int    `json:"created_at_unix"`
	Message       string `json:"message"`
	Commit        Commit `json:"commit"`
}

type GetTagsData struct {
	Status     int    `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	TotalCount int    `json:"total_count"`
	Tags       []Tag  `json:"tags"`
}

// GetTags 获取仓库标签列表 https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-118749619
func (s *TagService) GetTags(request *GetTagsRequest) (*GetTagsData, *Response, error) {

	u := fmt.Sprintf("/v1/%s/%s/tags.json", request.Owner, request.Repo)

	parsedUrl, err := url.Parse(u)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, nil, err
	}

	queryParams := parsedUrl.Query()

	if request.Page != 0 {
		queryParams.Set("page", strconv.Itoa(request.Page))
	}
	if request.Limit != 0 {
		queryParams.Set("limit", strconv.Itoa(request.Limit))
	}
	if request.Name != "" {
		queryParams.Set("name", request.Name)
	}

	parsedUrl.RawQuery = queryParams.Encode()

	req, err := s.client.NewRequest(http.MethodGet, parsedUrl.String(), request)
	if err != nil {
		return nil, nil, err
	}

	var getTagsData *GetTagsData
	resp, err := s.client.Do(req, &getTagsData)
	if err != nil {
		return nil, resp, err
	}

	return getTagsData, resp, nil
}
