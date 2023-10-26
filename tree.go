package main

// 将JSON项目定义转换成树
func toTree(node *ItemNode, itemNodeList []*ItemNode) {
	nodeList := make([]*ItemNode, 0)

	for i := 0; i < len(itemNodeList); i++ {
		v := itemNodeList[i]

		if v.Level <= node.Level {
			break
		}

		// 当前元素的下一个元素的 level 更大，则下一个元素为当前元素的子
		if i < len(itemNodeList)-1 && itemNodeList[i+1].Level > v.Level {
			descendant := make([]*ItemNode, 0)
			for j := i + 1; j < len(itemNodeList); j++ {
				v2 := itemNodeList[j]
				if v2.Level > v.Level {
					descendant = append(descendant, v2)
				}
				if v2.Level <= v.Level {
					break
				}
			}

			toTree(v, descendant)
		}
	}

	for i := 0; i < len(itemNodeList); i++ {
		v := itemNodeList[i]

		// 元素的 level 为 参数节点的 leve +1 时，则添加到参数节点的子中
		if v.Level == node.Level+1 {
			nodeList = append(nodeList, v)
		}

		if v.Level < node.Level {
			break
		}
	}

	node.Children = nodeList
}

// 去掉多余的中间节点
func grandToChild(node *ItemNode) {
	for _, child := range node.Children {

		// 当前节点是列表，且子节点是对象，则将子节点的所有子节点（当前节点的孙节点）放到当前节点的子节点中。
		if node.NodeType == TypeArray && child.NodeType == TypeObject {
			node.ClassName = child.Name
			node.ClassComment = child.Comment
			var children []*ItemNode
			for _, grandChild := range child.Children {
				grandChild.Level = child.Level
				children = append(children, grandChild)
			}
			node.Children = children
		}
		grandToChild(child)
	}
}

// 所有需要定义成类的对象
func allObject(node *ItemNode, objectNode *ItemNode) {
	objectNode.Children = append(objectNode.Children, node)
	for _, v := range node.Children {
		// 列表或对象时递归调用
		if v.NodeType == TypeArray {
			allObject(v, objectNode)
		}
		if v.NodeType == TypeObject {
			allObject(v, objectNode)
		}
	}
}
