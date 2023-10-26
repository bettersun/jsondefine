package main

import (
	"fmt"
	"github.com/bettersun/rain/yaml"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

func main() {
	// 读取配置文件
	file := "config.yml"
	s := fmt.Sprintf("=====读取配置文件： %s =====", file)
	log.Println(s)

	var config Config
	err := yaml.YamlFileToStruct(file, &config)
	if err != nil {
		log.Println(err)
	}

	s = fmt.Sprintf("=====打开 XLSX 文件： %s =====", config.FileName)
	log.Println(s)
	// 打开 XLSX 文件
	f, err := excelize.OpenFile(config.FileName)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	s = fmt.Sprintf("=====查找包含JSON定义的工作表=====")
	log.Println(s)
	// 包含JSON定义的工作表
	var jsonDefineList []string
	for _, v := range f.GetSheetList() {
		// 响应开始位置
		result, err := f.SearchSheet(v, config.KeywordResponse, true)
		if err != nil {
			log.Println(err)
			continue
		}

		if len(result) > 0 {
			jsonDefineList = append(jsonDefineList, v)
		}
	}

	s = fmt.Sprintf("=====生成代码=====")
	log.Println(s)
	// 循环包含JSON定义的工作表，生成代码
	for i, v := range jsonDefineList {
		sheetName := v

		s = fmt.Sprintf("=====JSON定义工作表：%s =====", v)
		log.Println(s)

		// 响应开始位置
		result, err := f.SearchSheet(sheetName, config.KeywordResponse, true)
		if err != nil {
			log.Println(err)
		}

		// 响应行号
		var rowIdx int
		for _, v := range result {
			_, rowIdx, _ = excelize.SplitCellName(v)
			// 取第一个
			break
		}

		// 所有行
		rows, err := f.GetRows(sheetName)
		jsonDefineInfo := getJsonDefineInfo(rows, rowIdx, config)

		// JSON 项目定义
		itemNodeList := itemList(rows, jsonDefineInfo, config)

		// 根节点
		var rootNode ItemNode
		rootNode.Level = -1
		// 根节点名称 需要定制
		rootNode.Name = config.RootNodeName + strconv.Itoa(i)

		// 列表转成树
		var treeNodeList []*ItemNode
		treeNodeList = append(treeNodeList, &rootNode)
		toTree(&rootNode, itemNodeList)

		// 去掉中间子节点
		grandToChild(&rootNode)

		// 遍历找出需要定义类的节点
		var objectNode ItemNode
		allObject(&rootNode, &objectNode)

		// 转换成 Dart 类代码
		codeInfoList := toClassCode(objectNode.Children)

		// 添加 model.dart
		codeInfoList = modelCode(codeInfoList)

		// 输出到代码文件
		writeCode(config, sheetName, codeInfoList)
	}

	s = fmt.Sprintf("=====代码生成已完成=====")
	log.Println(s)
}
