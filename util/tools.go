package util

import (
	"io/ioutil"
	"os"
)

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
