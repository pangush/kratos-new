package fileutil

import (
	"io"
	"net/http"
	"os"
)

func Download(url string, filename string) error {
	//下载
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	io.Copy(f, res.Body)

	return nil
}