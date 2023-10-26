package main

// 获取JSON项目定义相关的行列信息
func getJsonDefineInfo(rows [][]string, rowDefineStart int, config Config) JsonDefineInfo {
	var jsonDefineInfo JsonDefineInfo
	// JSON项目开始行号
	var rowIdxItemStart int

	maxRowCount := len(rows)
	for i := rowDefineStart; i < maxRowCount; i++ {
		row := rows[i]

		for j := 0; j < len(row); j++ {
			// JSON项目列
			if row[j] == config.KeywordItem {
				jsonDefineInfo.ColIdxItem = j
				rowIdxItemStart = i
			}
			// 类型列
			if row[j] == config.KeywordItemType {
				jsonDefineInfo.ColIdxItemType = j
				rowIdxItemStart = i
			}
			// 不可空列
			if row[j] == config.KeywordItemNotNull {
				jsonDefineInfo.ColIdxItemNotNull = j
				rowIdxItemStart = i
			}
			// 说明列
			if row[j] == config.KeywordItemComment {
				jsonDefineInfo.ColIdxItemComment = j
				rowIdxItemStart = i
			}
		}

		if rowIdxItemStart > 0 {
			// JSON项目开始行号为 JSON定义关键字 所在行 + 1
			jsonDefineInfo.RowIdxItemStart = rowIdxItemStart + 1
			break
		}
	}

	return jsonDefineInfo
}

// / 读取JSON项目定义
func itemList(rows [][]string, jsonDefineInfo JsonDefineInfo, config Config) []*ItemNode {
	maxRowCount := len(rows)
	var itemNodeList []*ItemNode

	for i := jsonDefineInfo.RowIdxItemStart; i < maxRowCount; i++ {
		row := rows[i]
		var itemNode ItemNode
		colIdxTmp := len(row)

		// JSON项目定义最大列
		if colIdxTmp >= jsonDefineInfo.ColIdxItemType {
			colIdxTmp = jsonDefineInfo.ColIdxItemType
		}

		// JSON 项目
		for j := jsonDefineInfo.ColIdxItem; j < colIdxTmp; j++ {
			if row[j] != "" {
				itemNode.Name = row[j]
				// 缩进一个单元格为一个层级
				// 如果一次缩进两个单元格，则不能正常解析
				itemNode.Level = j - jsonDefineInfo.ColIdxItem
				break
			}
		}

		// 类型
		var itemType string
		if len(row) >= jsonDefineInfo.ColIdxItemType {
			itemType = row[jsonDefineInfo.ColIdxItemType]
			itemNode.NodeType = getNodeType(itemType, config)
		}

		// 不可空
		var sNotNull string
		if len(row) >= jsonDefineInfo.ColIdxItemNotNull {
			sNotNull = row[jsonDefineInfo.ColIdxItemNotNull]
		}
		itemNode.NotNull = sNotNull == config.NotNullFlag

		// 说明
		if len(row) >= jsonDefineInfo.ColIdxItemComment {
			itemNode.Comment = row[jsonDefineInfo.ColIdxItemComment]
		}

		// 全为空，则认为是JSON项目定义结束
		if itemNode.Name == "" && itemType == "" && sNotNull == "" && itemNode.Comment == "" {
			break
		}

		itemNodeList = append(itemNodeList, &itemNode)
	}

	return itemNodeList
}

// 根据JSON项目的格式判断类型
func getNodeType(tp string, config Config) NodeType {
	tpItem := TypeNone

	if tp == config.KeywordItemTypeArray {
		tpItem = TypeArray
	}

	if tp == config.KeywordItemTypeObject {
		tpItem = TypeObject
	}

	if tp == config.KeywordItemTypeString {
		tpItem = TypeString
	}

	if tp == config.KeywordItemTypeNumber {
		tpItem = TypeInt
	}

	if tp == config.KeywordItemTypeBool {
		tpItem = TypeBool
	}

	return tpItem
}
