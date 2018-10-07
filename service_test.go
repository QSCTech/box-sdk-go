package box_test

import (
	"box"
	"github.com/Hexilee/gotten"
	"github.com/Hexilee/gotten/unmarshalers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpload(t *testing.T) {
	resp, err := box.GetService().Upload(&box.UploadParam{File: gotten.FilePath("go.sum")})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var result box.UploadResult
	assert.Nil(t, resp.Unmarshal(&result))
	assert.Equal(t, "go.sum", result.Data.Filename)
}

func TestDownload(t *testing.T) {
	resp, err := box.GetService().Upload(&box.UploadParam{File: gotten.FilePath("testAssets/avatar.jpg")})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var result box.UploadResult
	assert.Nil(t, resp.Unmarshal(&result))
	assert.Equal(t, "avatar.jpg", result.Data.Filename)

	resp, err = box.GetService().Download(&box.TokenParam{Token: result.Data.Token})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	var info unmarshalers.FileInfo
	assert.Nil(t, resp.Unmarshal(&info))
}
