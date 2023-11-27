package gitlink

import (
	"encoding/json"
	"os"
	"testing"
)

// 上传文件
func TestPostAttachments(t *testing.T) {

	var token = os.Getenv("GO_GITLINK_TOKEN")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	attachmentsData, response, err := gitClient.Attachments.PostAttachments("LICENSE")
	if err != nil {
		t.Fatalf("列出授权用户的所有仓库 异常：%s", err)
	}

	t.Log(response.Status)
	t.Log(attachmentsData.Filesize)
	t.Log(attachmentsData.Id)
	t.Log(attachmentsData.Status)
	t.Log(attachmentsData.Message)

	jsonData, err := json.Marshal(attachmentsData)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log(string(jsonData))
}
