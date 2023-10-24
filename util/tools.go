package util

import (
	"io/ioutil"
	"os"
)

//PathExists 判断文件夹是否存在
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


func ListDir(dirname string) ([]string, error) {
	infos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(infos))
	for i, info := range infos {
		names[i] = info.Name()
	}
	return names, nil
}