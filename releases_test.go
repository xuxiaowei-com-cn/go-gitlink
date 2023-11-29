package gitlink

import (
	"encoding/json"
	"os"
	"testing"
)

// 创建发行版
func TestPostReleases(t *testing.T) {
	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	if err != nil {
		t.Fatalf("创建客户端异常：%s", err)
	}

	gitClient.Cookie = cookie

	requestPath := &PostReleasesRequestPath{
		Owner: "xuxiaowei-com-cn",
		Repo:  "go-gitlink",
	}

	requestBody := &PostReleasesRequestBody{
		Body:    "v0.0.2-test 测试发布描述，无需关注",
		Name:    "v0.0.2-test 测试发布标题，无需关注",
		TagName: "v0.0.2-test",
		//Prerelease: true,
		Draft:         true,
		AttachmentIds: []int64{442010},
	}

	postReleases, response, err := gitClient.Releases.PostReleases(requestPath, requestBody)

	t.Log("response.Status：", response.Status)
	t.Log("Status：", postReleases.Status)
	t.Log("Message：", postReleases.Message)

	jsonData, err := json.Marshal(postReleases)
	if err != nil {
		t.Log("转换失败：", err)
		return
	}
	t.Log("JSON：", string(jsonData))
}
