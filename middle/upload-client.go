package middle

import (
	"io/ioutil"
	"path/filepath"

	"github.com/parnurzeal/gorequest"
)

func UploadClient(uploadUri, filename, fieldname string) (string, error) {

	f, _ := filepath.Abs(filename)
	fb := filepath.Base(filename)
	bytesOfFile, _ := ioutil.ReadFile(f)
	client := gorequest.New()

	client.Post(uploadUri).
		Type("multipart").
		SendFile(bytesOfFile, fb, fieldname)

	response, payload, err := client.End()
	if len(err) > 0 {
		return payload, err[0]
	}
	if response.StatusCode == 200 {
		return payload, nil
	}
	return "", nil
}
