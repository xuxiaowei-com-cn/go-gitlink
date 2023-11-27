package gitlink

import (
	"encoding/json"
	"os"
	"testing"
)

// 项目列表
func TestGetProjects(t *testing.T) {

	var token = os.Getenv("GO_GITLINK_TOKEN")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	request := &GetProjectsRequest{
		ListOptions: ListOptions{
			Page:  1,
			Limit: 12,
		},
	}

	projectsData, response, err := gitClient.Projects.GetProjects(request)
	if err != nil {
		t.Fatalf("列出授权用户的所有仓库 异常：%s", err)
	}

	t.Log(response.Status)
	t.Log(len(projectsData.Projects))
	t.Log(projectsData.Status)
	t.Log(projectsData.Message)
	t.Log(projectsData.TotalCount)
	t.Log(projectsData.Projects[0].Name)

	jsonData, err := json.Marshal(projectsData)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log(string(jsonData))
}
