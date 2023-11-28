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

	t.Log("response.Status：", response.Status)
	t.Log("Status：", projectsData.Status)
	t.Log("Message：", projectsData.Message)
	t.Log("len：", len(projectsData.Projects))
	t.Log("TotalCount：", projectsData.TotalCount)
	t.Log("Projects[0].Name：", projectsData.Projects[0].Name)

	jsonData, err := json.Marshal(projectsData)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log("JSON：", string(jsonData))
}
