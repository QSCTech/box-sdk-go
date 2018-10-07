package box

import (
	"github.com/Hexilee/gotten"
)

type (
	Service struct {
		Upload      func(param *UploadParam) (gotten.Response, error) `method:"POST" path:"item/add_item"`
		Change      func(param *ChangeParam) (gotten.Response, error) `method:"POST" path:"item/change_item"`
		Stat        func(param *TokenParam) (gotten.Response, error)  `path:"item/issec/{token}"`              // resp: YES / NO
		Verify      func(param *SecParam) (gotten.Response, error)    `path:"item/verify/{token}/{sec_token}"` // resp: Y / N
		Download    func(param *TokenParam) (gotten.Response, error)  `path:"item/get/{token}"`
		DownloadSec func(param *SecParam) (gotten.Response, error)    `path:"item/get/{token}/{sec_token}"` // fail: html, no Content-Disposition
	}
)
type (
	// Upload
	UploadParam struct {
		PhpSession string          `type:"part" key:"PHP_SESSION_UPLOAD_PROGRESS" default:"qscbox"`
		Filecount  int             `type:"part" default:"1"`
		File       gotten.FilePath `type:"part" require:"true"`
		Callback   string          `type:"part" default:"handleUploadCallback"`
		IsIe9      int             `type:"part" default:"0"`
	}

	UploadResult struct {
		Data struct {
			Error      string
			Expiration int
			Filename   string
			Secret     string
			SecureId   string
			Token      string
		}
		Err        int // 0: SUCCESS; -1: Fail
		Expiration int
		Msg        string
	}
)

type (
	// change
	ChangeParam struct {
		NewToken   string `type:"form"`
		Jiami      string `type:"form"`
		OldToken   string `type:"form" require:"true"`
		SecureId   string `type:"form" require:"true"`
		TokenSec   string `type:"form"`
		OrdSrc     string `type:"form"`
		Expiration int    `type:"form" require:"true"`
	}

	ChangeResult struct {
		Message  string
		NewToken string
		Status   int // 0: Success; 1: fail
		Url      string
	}
)

type (
	// get status
	TokenParam struct {
		Token string `type:"path"`
	}
)

type (
	// verify security token
	SecParam struct {
		Token    string `type:"path"`
		SecToken string `type:"path"`
	}
)
