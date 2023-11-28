package gitlink

import (
	"encoding/json"
	"os"
	"testing"
)

// 上传文件
func TestPostAttachments(t *testing.T) {

	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	gitClient.cookie = cookie

	attachmentsData, response, err := gitClient.Attachments.PostAttachments("LICENSE")
	if err != nil {
		t.Fatalf("上传文件 异常：%s", err)
	}

	t.Log("response.Status：", response.Status)
	t.Log("Status：", attachmentsData.Status)
	t.Log("Message：", attachmentsData.Message)
	t.Log("Id：", attachmentsData.Id)
	t.Log("Filesize：", attachmentsData.Filesize)

	jsonData, err := json.Marshal(attachmentsData)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log("JSON：", string(jsonData))
}
