package gitlink

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

// 项目列表
func TestGetProjects(t *testing.T) {

	gitClient, err := NewClient("")
	assert.NoError(t, err)

	request := &GetProjectsRequest{
		ListOptions: ListOptions{
			Page:  1,
			Limit: 12,
		},
	}

	projectsData, response, err := gitClient.Projects.GetProjects(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 0, projectsData.Status)
	assert.Equal(t, "", projectsData.Message)
	assert.NotEqual(t, int64(0), len(projectsData.Projects))
	assert.NotEqual(t, int64(0), projectsData.TotalCount)
}
