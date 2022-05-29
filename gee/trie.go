package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 待匹配路由，例如 /p/:lang，一个pattern可以理解为trie树的一条路径
	part     string  // 路由中的一部分，例如 :lang，part可以理解为trie树的一个节点
	children []*node // 子节点，例如 [doc, tutorial, intro]
	isWild   bool    // 是否精确匹配，part 含有 : 或 * 时为true，为了匹配节点时使用
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// trie树的插入
// 找到对应深度的节点，则递归插入
// 找不到的话，则新建一个节点，并插入到当前节点的children中
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {  // 此时整个树遍历完了，函数结束，此时才写上node的pattern字段
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)

	// 如果没有找到对应的节点，则创建一个，part实为路径中的一个部分，如/a/b/c中的b
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)  // 将child加入到n的孩子们中
	}
	child.insert(pattern, parts, height+1) // 继续递归进行下一个深度的插入操作
}

// trie树的搜索
// 正常情况，路由表中trie树深度和路径上的节点长度相等，则结束匹配
// 带通配符*的匹配，需要判断节点字符串**前缀**是否含*
// 上述条件都满足则表示找到了。这里需要注意pattern为空的情况
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {  // 说明没有这个pattern，也就是没有搜到因为trie树插入的时候，pattern中所有的part都插完了才写pattern
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
