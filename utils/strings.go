package utils

import (
	"bytes"
	"io/ioutil"

	"github.com/valyala/bytebufferpool"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// Concat some strings
func Concat(a string, b ...string) string {
	if len(b) == 0 {
		return a
	}
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	buf.SetString(a)
	for _, s := range b {
		buf.WriteString(s)
	}

	return string(buf.B)
}
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
