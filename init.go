package box

import (
	"errors"
	"github.com/Hexilee/gotten"
	"github.com/Hexilee/gotten/headers"
	"github.com/Hexilee/gotten/unmarshalers"
	"net/http"
	"strconv"
)

var (
	service *Service

	mediaChecker gotten.CheckerFunc = func(response *http.Response) bool {
		return response.Header.Get(headers.HeaderContentDisposition) != ""
	}

	textChecker gotten.CheckerFunc = func(response *http.Response) (yes bool) {
		contentLength := response.Header.Get(headers.HeaderContentLength)
		contentType := response.Header.Get(headers.HeaderContentType)
		length, err := strconv.Atoi(contentLength)
		if err != nil {
			yes = false
		}

		if yes {
			yes = length < 10 && (contentType == headers.MIMETextPlain || contentType == headers.MIMETextHTMLCharsetUTF8)
		}
		return
	}
)

func handlePlainText(data []byte, v interface{}) (err error) {
	yes, ok := v.(*bool)
	if !ok {
		err = errors.New("param kind must be *bool")
	}

	if err != nil {
		switch string(data) {
		case "Y":
			fallthrough
		case "YES":
			*yes = true
		case "NO":
			fallthrough
		case "N":
			*yes = false
		default:
			err = errors.New("data content is unsupported: " + string(data))
		}
	}
	return
}

func init() {
	fileCtr, err := unmarshalers.NewFileCtr()
	if err != nil {
		panic(err)
	}
	creator, err := gotten.NewBuilder().
		SetBaseUrl("https://box.zjuqsc.com/item").
		AddReaderUnmarshaler(fileCtr, mediaChecker).
		AddUnmarshalFunc(handlePlainText, textChecker).
		Build()
	if err != nil {
		panic(err)
	}

	err = creator.Impl(service)
	if err != nil {
		panic(err)
	}
}

func GetService() *Service {
	return service
}
