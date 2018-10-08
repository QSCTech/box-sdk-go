[![Coverage Status](https://coveralls.io/repos/github/QSCTech/box-sdk-go/badge.svg)](https://coveralls.io/github/QSCTech/box-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/QSCTech/box-sdk-go)](https://goreportcard.com/report/github.com/QSCTech/box-sdk-go)
[![Build Status](https://travis-ci.org/QSCTech/box-sdk-go.svg?branch=master)](https://travis-ci.org/QSCTech/box-sdk-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/QSCTech/box-sdk-go/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/QSCTech/box-sdk-go?status.svg)](https://godoc.org/github.com/QSCTech/box-sdk-go)

### QSC BOX SDK

This is a demo project for [gotten](https://github.com/Hexilee/gotten).

### Usage

This package supports upload/download functions of single file.

```go
type (
	Service struct {
		Upload      func(param *UploadParam) (gotten.Response, error) `method:"POST" path:"item/add_item"`
		Change      func(param *ChangeParam) (gotten.Response, error) `method:"POST" path:"item/change_item"`
		Stat        func(param *TokenParam) (gotten.Response, error)  `path:"item/issec/{token}"`              // resp: YES / NO<SP>
		Verify      func(param *SecParam) (gotten.Response, error)    `path:"item/verify/{token}/{sec_token}"` // resp: Y / N
		Download    func(param *TokenParam) (gotten.Response, error)  `path:"item/get/{token}"`
		DownloadSec func(param *SecParam) (gotten.Response, error)    `path:"item/get/{token}/{sec_token}"` // fail: html, no Content-Disposition
	}
)
```

Get service singleton by GetService()

```go
resp, err := box.GetService().Upload(&box.UploadParam{File: gotten.FilePath("testAssets/avatar.jpg")})
if err == nil {
    var result box.UploadResult
    err = resp.Unmarshal(&result)	
}
```

