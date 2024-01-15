package gitlink

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// 创建发行版
func TestPostReleases(t *testing.T) {
	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	assert.NoError(t, err)

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
		AttachmentIds: []string{"a3277d52-8d8a-4492-ada1-77bfa047facb"},
	}

	postReleases, response, err := gitClient.Releases.PostReleases(requestPath, requestBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 0, postReleases.Status)
	assert.Equal(t, "发布成功", postReleases.Message)
}
