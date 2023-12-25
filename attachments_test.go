package gitlink

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

// 上传文件
func TestPostAttachments(t *testing.T) {

	var token = os.Getenv("GO_GITLINK_TOKEN")
	var cookie = os.Getenv("GO_GITLINK_COOKIE")

	gitClient, err := NewClient(token)
	assert.NoError(t, err)

	gitClient.Cookie = cookie

	attachmentsData, response, err := gitClient.Attachments.PostAttachments("LICENSE", "")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	//t.Log(attachmentsData)
	assert.Nil(t, attachmentsData.Status)
	assert.Equal(t, "", attachmentsData.Message)
	assert.NotEqual(t, "", attachmentsData.Id)
	assert.NotEqual(t, "", attachmentsData.Filesize)
	assert.NotEqual(t, "", attachmentsData.Title)
	assert.NotEqual(t, "", attachmentsData.Url)
}
