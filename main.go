package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"kratos-new/pkg/fileutil"
)

const url = "https://codeload.github.com/pangush/kratos-demo/zip/master"

var projectName  string

func init()  {
	flag.StringVar(&projectName, "project", "", "project name")
}

func main()  {
	flag.Parse()
	if projectName == "" {
		panic("Please specify a file name by -project flag.")
	}
	path := projectName

	// 返回所给路径的绝对路径
	pathAbs, _ := filepath.Abs(path)

	// 返回路径最后一个元素
	pathBase := filepath.Base(pathAbs)

	bo , err := fileutil.CheckFileIsExists(pathAbs)
	if err != nil {
		panic(err)
	}

	if bo {
		panic(fmt.Errorf("项目'%v'已经存在 ", pathAbs))
	}

	zipFile := "kratos-demo-master.zip"
	err = fileutil.Download(url, zipFile)
	if err != nil {
		panic(err)
	}

	err = fileutil.UnZip(zipFile, "./")
	if err != nil {
		panic(err)
	}

	err = os.Remove(zipFile)
	if err != nil {
		panic(err)
	}

	// 重命名文件夹
	err = os.Rename("kratos-demo-master", pathBase)
	if err != nil {
		panic(err)
	}

	files, err := fileutil.GetFileAll(pathAbs)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileBytes, err := fileutil.ReadFile(file)
		if err != nil {
			panic(err)
		}
		if !bytes.Contains(fileBytes,[]byte("kratos-demo")) {
			continue
		}
		
		replaceBytes := bytes.Replace(fileBytes, []byte("kratos-demo"), []byte(pathBase), -1)

		err = fileutil.WriteFile(file, replaceBytes)
		if err != nil {
			panic(err)
		}
	}
}
