package filex

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/karrick/godirwalk"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tsingson/fastx/utils"
)

const bufferSize = 65536

func Exists(name string) bool {
	afs := afero.NewOsFs()
	b, e := afero.Exists(afs, name)
	if e != nil {
		return false
	}
	return b
}

func DirExists(name string) bool {
	afs := afero.NewOsFs()
	b, e := afero.DirExists(afs, name)
	if e != nil {
		return false
	}
	return b
}

func WriteToFile(c []byte, filename string) error {
	//将指定内容写入到文件中
	err := ioutil.WriteFile(filename, c, 0666)
	return err
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListFiles(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

func ListDir(dirPth string) (files []string, err error) {
	files = make([]string, 0, 100)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			files = append(files, utils.StrBuilder(dirPth, PthSep, fi.Name()))
		}
	}
	//litter.Dump(files)
	return files, nil
}

func Md5CheckSum(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil {
		return "", err
	} else if info.IsDir() {
		return "", nil
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		hash.Write(buf[:n])
	}

	checksum := fmt.Sprintf("%x", hash.Sum(nil))
	return checksum, nil
}
func ListSubPath(osDirname string) ([]string, error) {

	children, err := godirwalk.ReadDirnames(osDirname, nil)
	if err != nil {
		err1 := errors.Wrap(err, "cannot get list of directory children")
		return nil, err1
	}
	sort.Strings(children)
	var sublist []string
	sublist = make([]string, len(children))
	for _, child := range children {
		pathNode := utils.StrBuilder(osDirname, "/", child, "/")
		//	fmt.Printf("%s\n", pathNode)
		sublist = append(sublist, pathNode)
	}

	return sublist, nil
}
