package box

import "net/http"

//go:generate impler Service

/*
@Base https://box.zjuqsc.com/item
*/
type Service interface {
	/*
		@Desc upload a file by file path
		@Post /add_item
		@File(file) {path}
		@Param(PHP_SESSION_UPLOAD_PROGRESS) qscbox
		@Param(filecount) 1
		@Param(callback) handleUploadCallback
		@Param(is_ie9) 0
		@Body multipart
	*/
	Upload(path string) (result *UploadResult, statusCode int, err error)

	/*
		@Desc change file status
		@Post /change_item
		@Param(new_token) {newToken}
		@Param(old_token) {oldToken}
		@Param(secure_id) {secureId}
		@Param(token_sec) {tokenSec}
		@Param(old_sec)   {oldSec}
		@Body form
	*/
	Change(newToken, jiami, oldToken, secureId, tokenSec, oldSec string, expiration int) (result *ChangeResult, statusCode int, err error)

	/*
		@Desc get file status
		@Get issec/{token}
	*/
	Stat(token string) (*http.Response, error)

	/*
		@Desc verify secure token
		@Get verify/{token}/{secToken}
	*/
	Verify(token, secToken string) (*http.Response, error)

	/*
		@Desc download a file
		@Get get/{token}
	*/
	Download(token string) (*http.Response, error)

	/*
		@Desc download a secure file
		@Get get/{token}/{secToken}
	*/
	DownloadSec(token, secToken string) (*http.Response, error)
}

type UploadResult struct {
	Data struct {
		Error      string
		Expiration int
		Filename   string
		Secret     string `json:"secret"`
		SecureId   string `json:"secure_id"`
		Token      string
	}
	Err        int // 0: SUCCESS; -1: Fail
	Expiration int
	Msg        string
}

type ChangeResult struct {
	Message  string
	NewToken string
	Status   int // 0: Success; 1: fail
	Url      string
}
