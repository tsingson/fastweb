package utils

import (
	"reflect"
	"strings"
	"unsafe"
)

// BytesToStringUnsafe converts a byte slice to a string.
// It's fasthttpgx, but not safe. Use it only if you know what you're doing.
func BytesToStringUnsafe(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{
		Data: bytesHeader.Data,
		Len:  bytesHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&strHeader))
}

// StringToBytesUnsafe converts a string to a byte slice.
// It's fasthttpgx, but not safe. Use it only if you know what you're doing.
func StringToBytesUnsafe(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := reflect.SliceHeader{
		Data: strHeader.Data,
		Len:  strHeader.Len,
		Cap:  strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bytesHeader))
}

func StrBuilder(args ...string) string {
	var str strings.Builder

	for _, v := range args {
		str.WriteString(v)
	}
	return str.String()
}
