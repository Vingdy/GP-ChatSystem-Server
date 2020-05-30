package utils

import (
	"os/exec"
	"os"
	"path/filepath"
	"strings"
)

// GetProDir 用于获取项目根目录
func GetProDir() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	end := strings.LastIndex(path, string(os.PathSeparator))
	proPath := path[:end]
	return proPath, nil
}