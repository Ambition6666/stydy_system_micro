package test

import (
	"fmt"
	"path"
	"runtime"

	"testing"
)

func TestXwd(t *testing.T) {
	// 获取当前文件的路径
	_, filename, _, _ := runtime.Caller(0)
	fmt.Printf("%s/../\n", filename)
	root := path.Dir(path.Dir(filename))
	fmt.Println(root)
}
