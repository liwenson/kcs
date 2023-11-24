package util

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"os/exec"
	"runtime"
)

// PathExists 路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		// 创建文件夹
		err = os.MkdirAll(path, os.ModePerm)
		if err == nil {
			return true, nil
		}
	}

	return false, err
}

// FileExists 文件是否存在，不存在是否需要创建
func FileExists(path string, mk bool) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		if mk {
			// 创建文件夹
			_, err = os.Create(path)
			if err == nil {
				return true, nil
			}
		}
		return false, nil
	}

	return false, err
}

// ListDir 列出目录中的文件
func ListDir(dirname string) ([]string, error) {
	infos, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(infos))
	for i, info := range infos {
		names[i] = info.Name()
	}
	return names, nil
}

// MD5 MD5校验，输入字byte
func MD5(v []byte) string {
	//d := []byte(v)
	m := md5.New()
	m.Write(v)
	return hex.EncodeToString(m.Sum(nil))
}

// MD5Str MD5校验，输入字符串
func MD5Str(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// ClearScreen 清屏
func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
