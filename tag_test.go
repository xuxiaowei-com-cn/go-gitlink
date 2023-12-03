package gitlink

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// 删除一个标签
func TestDeleteTag(t *testing.T) {

	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	assert.NoError(t, err)

	gitClient.Cookie = cookie

	var deleteTagRequest = &DeleteTagRequest{
		Owner: "xuxiaowei-com-cn",
		Repo:  "go-gitlink",
		Tag:   "v0.0.2-test",
	}

	deleteTag, response, err := gitClient.Tag.DeleteTag(deleteTagRequest)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 0, deleteTag.Status)
	assert.Equal(t, "success", deleteTag.Message)
}

// 获取仓库标签列表
func TestGetTags(t *testing.T) {

	gitClient, err := NewClient("")
	assert.NoError(t, err)

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
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 0, getTagsData.Status)
	assert.Equal(t, "", getTagsData.Message)
	assert.NotEqual(t, int64(0), len(getTagsData.Tags))
	assert.NotEqual(t, int64(0), getTagsData.TotalCount)
}
