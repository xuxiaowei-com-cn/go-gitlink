package gitlink

import (
	"encoding/json"
	"os"
	"testing"
)

// 删除一个标签
func TestDeleteTag(t *testing.T) {

	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	gitClient.cookie = cookie

	var deleteTagRequest = &DeleteTagRequest{
		Owner: "xuxiaowei-com-cn",
		Repo:  "go-gitlink",
		Tag:   "v0.0.2-test",
	}

	deleteTag, response, err := gitClient.Tag.DeleteTag(deleteTagRequest)
	if err != nil {
		t.Fatalf("删除一个标签 异常：%s", err)
	}

	t.Log(response.Status)
	t.Log(deleteTag.Status)
	t.Log(deleteTag.Message)

	jsonData, err := json.Marshal(deleteTag)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log(string(jsonData))
}

// 获取仓库标签列表
func TestGetTags(t *testing.T) {
	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	gitClient.cookie = cookie

	var getTagsRequest = &GetTagsRequest{
		Owner: "xuxiaowei-com-cn",
		Repo:  "go-gitlink",
	}

	getTagsData, response, err := gitClient.Tag.GetTags(getTagsRequest)
	if err != nil {
		t.Fatalf("获取仓库标签列表 异常：%s", err)
	}

	t.Log(response.Status)
	t.Log(getTagsData.Status)
	t.Log(getTagsData.Message)
	t.Log(getTagsData.TotalCount)

	jsonData, err := json.Marshal(getTagsData)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log(string(jsonData))
}
