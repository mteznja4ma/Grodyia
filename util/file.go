package util

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/**
 * 获取当前系统路径
 *
 * @return (string)
 **/
func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New("error: Can't find '/' or '\\'")
	}
	return string(path[0 : i+1]), nil
}
