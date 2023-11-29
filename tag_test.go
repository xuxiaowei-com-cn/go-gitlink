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

	t.Log("response.Status：", response.Status)
	t.Log("Status：", deleteTag.Status)
	t.Log("Message：", deleteTag.Message)

	jsonData, err := json.Marshal(deleteTag)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log("JSON：", string(jsonData))
}

// 获取仓库标签列表
func TestGetTags(t *testing.T) {

	gitClient, err := NewClient("")
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	var getTagsRequest = &GetTagsRequestPath{
		Owner: "xuxiaowei-com-cn",
		Repo:  "go-gitlink",
	}

	var getTagsRequestQuery = &GetTagsRequestQuery{
		Name: "2",
		ListOptions: ListOptions{
			Page:  1,
			Limit: 4,
		},
	}

	getTagsData, response, err := gitClient.Tag.GetTags(getTagsRequest, getTagsRequestQuery)
	if err != nil {
		t.Fatalf("获取仓库标签列表 异常：%s", err)
	}

	t.Log("response.Status：", response.Status)
	t.Log("Status：", getTagsData.Status)
	t.Log("Message：", getTagsData.Message)
	t.Log("len：", len(getTagsData.Tags))
	t.Log("TotalCount：", getTagsData.TotalCount)

	jsonData, err := json.Marshal(getTagsData)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log("JSON：", string(jsonData))
}
