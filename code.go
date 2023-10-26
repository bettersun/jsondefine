package main

import (
	"fmt"
	"github.com/bettersun/rain"
	"strings"
)

// 转换成类代码
func toClassCode(nodeList []*ItemNode) []CodeInfo {
	var codeInfoList []CodeInfo

	for _, v := range nodeList {
		// 文件名
		classFileNm := classFileName(v)

		// Header
		header := codeImport + fmt.Sprintf("part '%s.g.dart';", classFileNm)

		// 类注释
		clsComment := classComment(v)

		// 类名
		classNm := className(v)

		// 成员属性
		var properties []string
		for _, v1 := range v.Children {
			properties = append(properties, propertyDefine(v1))
		}
		props := strings.Join(properties, "\n")

		// 构造函数
		sCon := constructor(v)

		// fromJson/toJson
		fromJson :=
			fmt.Sprintf("  factory %s.fromJson(Map<String, dynamic> json) => _$%sFromJson(json);", classNm, classNm)
		toJson :=
			fmt.Sprintf("  Map<String, dynamic> toJson() => _$%sToJson(this);", classNm)

		// 类的所有代码
		classDef :=
			fmt.Sprintf("%s\n\n%s%sclass %s{\n%s\n\n%s\n%s\n\n%s\n}\n", header, clsComment, codeAnnotation, classNm, props, sCon, fromJson, toJson)

		// 代码输出信息
		var codeInfo CodeInfo
		codeInfo.CodeFile = classFileNm
		codeInfo.Code = classDef
		codeInfoList = append(codeInfoList, codeInfo)
	}

	return codeInfoList
}

// 类文件名
func classFileName(node *ItemNode) string {
	// 驼峰转下划线连接
	if node.ClassName != "" {
		return rain.CamelToSnake(node.ClassName)
	} else {
		return rain.CamelToSnake(node.Name)
	}
}

// 类名
func className(node *ItemNode) string {
	// 首字母大写
	if node.ClassName != "" {
		return rain.UperFirst(node.ClassName)
	} else {
		return rain.UperFirst(node.Name)
	}
}

// 类注释
func classComment(node *ItemNode) string {
	if node.ClassComment != "" {
		return fmt.Sprintf("/// %s \n", node.ClassComment)
	} else {
		return fmt.Sprintf("/// %s \n", node.Comment)
	}
}

// 构造函数
func constructor(node *ItemNode) string {
	var properties []string
	for _, v := range node.Children {

		var property string
		//value := defaultValue(v.NodeType)
		if v.NotNull {
			//property = fmt.Sprintf("    this.%s = %s", v.Name, value)
			property = fmt.Sprintf("    required this.%s", v.Name)
		} else {
			property = fmt.Sprintf("    this.%s", v.Name)
		}

		properties = append(properties, property)
	}

	properties = append(properties, "")
	conProperties := strings.Join(properties, ",\n")

	return fmt.Sprintf("  %s({\n%s  });\n", className(node), conProperties)
}

// 成员定义
func propertyDefine(node *ItemNode) string {
	sTypeName := "%s %s"
	sTypeNullName := "%s? %s"
	sListName := "List<%s> %s"
	sListNullName := "List<%s>? %s"

	property := ""
	typeProp := propertyType(node)
	if node.NotNull {
		property = fmt.Sprintf(sTypeName, typeProp, node.Name)
		if node.NodeType == TypeArray {
			property = fmt.Sprintf(sListName, typeProp, node.Name)
		}
	} else {
		if node.NodeType == TypeArray {
			property = fmt.Sprintf(sListNullName, typeProp, node.Name)
		} else if typeProp == defaultType {
			property = fmt.Sprintf(sTypeName, typeProp, node.Name)
		} else {
			property = fmt.Sprintf(sTypeNullName, typeProp, node.Name)
		}
	}

	property = fmt.Sprintf("  /// %s\n  final %s;", node.Comment, property)
	return property
}

// 成员属性类型
func propertyType(node *ItemNode) string {
	typeProp := defaultType

	if node.NodeType == TypeArray {
		if node.Children != nil {
			typeProp = className(node)
		}
	} else if node.NodeType == TypeObject {
		typeProp = className(node)
	} else {
		typeProp = basicType(node.NodeType)
	}

	return typeProp
}

// 成员属性基本类型
func basicType(nodeType NodeType) string {
	valType := defaultType
	if nodeType == TypeString {
		valType = "String"
	}
	if nodeType == TypeInt || nodeType == TypeLong {
		valType = "int"
	}
	if nodeType == TypeFloat || nodeType == TypeDouble {
		valType = "double"
	}
	if nodeType == TypeBool {
		valType = "bool"
	}
	if nodeType == TypeNone {
		valType = defaultType
	}

	return valType
}

// 默认值
func defaultValue(nodeType NodeType) string {

	value := "''"
	if nodeType == TypeString {
		value = "''"
	}
	if nodeType == TypeInt || nodeType == TypeLong {
		value = "0"
	}
	if nodeType == TypeFloat || nodeType == TypeDouble {
		value = "0"
	}
	if nodeType == TypeBool {
		value = "false"
	}
	if nodeType == TypeArray {
		value = "const []"
	}

	return value
}

// 代码 model.dart
func modelCode(codeInfoList []CodeInfo) []CodeInfo {
	export := ""
	for _, v := range codeInfoList {
		// export 代码
		export = export + fmt.Sprintf("export '%s.dart';\n", v.CodeFile)
	}

	var codeInfo CodeInfo
	codeInfo.CodeFile = "model"
	codeInfo.Code = export

	codeInfoList = append(codeInfoList, codeInfo)
	return codeInfoList
}
