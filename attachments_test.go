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
	assert.Equal(t, 0, attachmentsData.Status)
	assert.Equal(t, "", attachmentsData.Message)
	assert.NotEqual(t, int64(0), attachmentsData.Id)
}
