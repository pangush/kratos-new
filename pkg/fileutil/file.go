package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func CheckFileIsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteFile(filename string, content []byte) error {
	bo, err := CheckFileIsExists(filename)
	if err != nil {
		return err
	}
	if bo {
		err := os.Remove(filename)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(filename) //创建文件
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(content) //写入文件
	if err != nil {
		return err
	}
	return nil
}

//读取到file中，再利用ioutil将file直接读取到[]byte中, 这是最优
func ReadFile(filename string)  ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return fd, nil
}

func GetFileAll(path string) ([]string, error) {
	filesInfo, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	files := make([]string, 0)
	for _, fileInfo := range filesInfo {
		file := filepath.Join(path, fileInfo.Name())
		if fileInfo.IsDir() {
			childFiles, err := GetFileAll(file)
			if err != nil {
				return nil, err
			}
			files = append(files, childFiles...)
		} else {
			files = append(files, file)
		}
	}

	return files, nil
}