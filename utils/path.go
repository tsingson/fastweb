package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func GetCurrentPath() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func GetCurrentExecDir() (dir string, err error) {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Printf("exec.LookPath(%s), err: %s\n", os.Args[0], err)
		return "", err
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("filepath.Abs(%s), err: %s\n", path, err)
		return "", err
	}
	dir = filepath.Dir(absPath)
	return dir, nil
}

func Md5CheckSum(filePath string) (result string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return
	}

	result = hex.EncodeToString(hash.Sum(nil))
	return
}
func GetFileSize(fullFilename string) int64 {

	fileInfo, err := os.Stat(fullFilename)
	if err != nil {
		return 0
	}
	fileSize := fileInfo.Size() //获取size
	return fileSize
}
