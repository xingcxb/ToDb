package lib

import (
	"os"
	goruntime "runtime"
	"strings"
)

// GetProgramSafePath 获取程序存放连接信息的路径
func GetProgramSafePath() string {
	var dirBuild strings.Builder
	dir, _ := os.Getwd()
	dirBuild.WriteString(dir)
	if goruntime.GOOS == "windows" {
		//windows下存放配置文件路径
		dirBuild.WriteString("\\safe\\")
	} else if goruntime.GOOS == "darwin" {
		//macOS下存放配置文件路径
		dirBuild.WriteString("/safe/")
	}
	return dirBuild.String()
}
