package main

// ItemNode JSON项目节点
type ItemNode struct {
	Level        int         // 层级
	Name         string      // JSON项目名
	Comment      string      // JSON项目说明
	ClassName    string      // JSON项目类名（需要定义成类的项目的类名，用于列表的元素对象类的定义）
	ClassComment string      // JSON项目类说明（需要定义成类的项目的类说明，用于列表的元素对象类的定义）
	NodeType     NodeType    // 节点类型
	NotNull      bool        // 不可空
	Children     []*ItemNode // 子节点
}

// JsonDefineInfo JSON项目定义的行列信息
type JsonDefineInfo struct {
	ColIdxItem        int // JSON项目 列
	ColIdxItemType    int // 项目类型 列
	ColIdxItemNotNull int // 不可空 列
	ColIdxItemComment int // 说明/注释 列
	RowIdxItemStart   int // JSON项目开始行号 行
}

// CodeInfo 代码输出信息
type CodeInfo struct {
	CodeFile string
	Code     string
}

// Config 程序运行配置
type Config struct {
	FileName string `yaml:"fileName"`

	KeywordResponse    string `yaml:"keywordResponse"`
	KeywordItem        string `yaml:"keywordItem"`
	KeywordItemType    string `yaml:"keywordItemType"`
	KeywordItemNotNull string `yaml:"keywordItemNotNull"`
	KeywordItemComment string `yaml:"keywordItemComment"`

	NotNullFlag string `yaml:"notNullFlag"`

	KeywordItemTypeArray  string `yaml:"keywordItemTypeArray"`
	KeywordItemTypeObject string `yaml:"keywordItemTypeObject"`
	KeywordItemTypeString string `yaml:"keywordItemTypeString"`
	KeywordItemTypeNumber string `yaml:"keywordItemTypeNumber"`
	KeywordItemTypeBool   string `yaml:"keywordItemTypeBool"`

	RootNodeName string `yaml:"rootNodeName"`

	CodeFileEx     string `yaml:"codeFileEx"`
	CodeOutputPath string `yaml:"codeOutputPath"`
}
