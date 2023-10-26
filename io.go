package main

import (
	"fmt"
	"github.com/bettersun/rain"
	"log"
	"os"
)

// 输出到文件
func writeCode(config Config, subPath string, codeInfoList []CodeInfo) {
	codePath := fmt.Sprintf("%s/%s", config.CodeOutputPath, subPath)
	//递归创建文件夹
	err := os.MkdirAll(codePath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	for _, v := range codeInfoList {
		codeFileFullPath := fmt.Sprintf("%s/%s%s", codePath, v.CodeFile, config.CodeFileEx)

		var code []string
		code = append(code, v.Code)
		err = rain.WriteFile(codeFileFullPath, code)
		if err != nil {
			log.Println(err)
		}

		s := fmt.Sprintf("代码文件： %s", codeFileFullPath)
		log.Println(s)
	}
}
