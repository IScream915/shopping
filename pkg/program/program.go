package program

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ExitWithError 错误退出的辅助函数，遇到错误时，将接收的err参数输出到标准错误流stderr，并终止程序
// 获取程序的名称，便于在错误信息中标识是哪个程序发生了错误
func ExitWithError(err error) {
	// os.Args[0]通常是程序执行时的路径
	// filepath.Base用来获取路径中的最后一部分，也就是程序的实际名称
	progName := filepath.Base(os.Args[0])
	_, _ = fmt.Fprintf(os.Stderr, "%s exit -1: %+v\n", progName, err)
	os.Exit(-1)
}

func GetProcessName() string {
	args := os.Args
	if len(args) > 0 {
		segments := strings.Split(args[0], "/")
		if len(segments) > 0 {
			return segments[len(segments)-1]
		}
	}
	return ""
}
