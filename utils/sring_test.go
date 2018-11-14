package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesToStringUnsafe(t *testing.T) {
	out := "hello"
	in := []byte(out)

	assert.Equal(t, out, BytesToStringUnsafe(in))
}

func TestStringToBytesUnsafe(t *testing.T) {
	in := "hello"
	out := []byte(in)

	assert.Equal(t, out, StringToBytesUnsafe(in))
}
