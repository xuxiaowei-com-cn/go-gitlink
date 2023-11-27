package gitlink

import (
	"net/http"
)

type SortByValue string // 排序方式

const (
	SortByDesc SortByValue = "desc" // 倒序
	SortByAsc  SortByValue = "asc"  // 正序
)

type SortDirectionValue string // 排序字段

const (
	SortDirectionUpdatedOn    SortDirectionValue = "updated_on"   // 公开
	SortDirectionCreatedOn    SortDirectionValue = "created_on"   // 私有
	SortDirectionForkedCount  SortDirectionValue = "forked_count" // 所有
	SortDirectionPraisesCount SortDirectionValue = "praises_count"
)

type ProjectTypeValue string // 项目类型

const (
	ProjectTypeCommon     ProjectTypeValue = "common"      // 普通项目
	ProjectTypeMirror     ProjectTypeValue = "mirror"      // 授权用户为仓库成员
	ProjectTypeSyncMirror ProjectTypeValue = "sync_mirror" // 授权用户为仓库所在组织并有访问仓库权限
)

type GetProjectsRequest struct {
	SortBy        SortByValue        `url:"sort_by,omitempty" json:"sort_by,omitempty"`               // 排序方式
	SortDirection SortDirectionValue `url:"sort_direction,omitempty" json:"sort_direction,omitempty"` // 排序字段
	Search        string             `url:"search,omitempty" json:"search,omitempty"`                 // 搜索关键词
	CategoryId    string             `url:"category_id,omitempty" json:"category_id,omitempty"`       // 项目分类id
	LanguageId    string             `url:"language_id,omitempty" json:"language_id,omitempty"`       // 项目语言id
	ProjectType   ProjectTypeValue   `url:"project_type,omitempty" json:"project_type,omitempty"`     // 项目类型
	ListOptions
}

type Language struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Topic struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	Id                  int         `json:"id"`
	RepoId              int         `json:"repo_id"`
	Identifier          string      `json:"identifier"`
	Name                string      `json:"name"`
	Description         string      `json:"description"`
	Visits              int         `json:"visits"`
	PraisesCount        int         `json:"praises_count"`
	ForkedCount         int         `json:"forked_count"`
	IsPublic            bool        `json:"is_public"`
	MirrorUrl           string      `json:"mirror_url"`
	Type                int         `json:"type"`
	LastUpdateTime      int         `json:"last_update_time"`
	TimeAgo             string      `json:"time_ago"`
	ForkedFromProjectId int         `json:"forked_from_project_id"`
	OpenDevops          bool        `json:"open_devops"`
	Platform            string      `json:"platform"`
	Author              Author      `json:"author"`
	Category            interface{} `json:"category"`
	Language            Language    `json:"language"`
	Topics              []Topic     `json:"topics"`
}

type ProjectsData struct {
	Status     int       `json:"status,omitempty"`
	Message    string    `json:"message,omitempty"`
	TotalCount int       `json:"total_count"`
	Projects   []Project `json:"projects"`
}

type ProjectsService struct {
	client *Client
}

// GetProjects 项目列表 https://apifox.com/apidoc/shared-da30afb0-9d2e-429b-a4bc-a83209e06021/api-102299292
func (s *ProjectsService) GetProjects(request *GetProjectsRequest) (*ProjectsData, *Response, error) {

	u := "projects"

	req, err := s.client.NewRequest(http.MethodGet, u, request)
	if err != nil {
		return nil, nil, err
	}

	var projectsData *ProjectsData
	resp, err := s.client.Do(req, &projectsData)
	if err != nil {
		return nil, resp, err
	}

	return projectsData, resp, nil
}
