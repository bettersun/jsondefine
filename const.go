package main

// import
const codeImport = "import 'package:json_annotation/json_annotation.dart';\nimport 'model.dart';\n\n"

// 注解
const codeAnnotation = "@JsonSerializable()\n"

// 默认类型
const defaultType = "dynamic"

// NodeType 节点类型
type NodeType int

const (
	TypeNone NodeType = iota
	TypeArray
	TypeObject
	TypeBool
	TypeNum
	TypeInt
	TypeLong
	TypeFloat
	TypeDouble
	TypeString
)
