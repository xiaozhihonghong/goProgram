package gee

import "strings"

type node struct {
	pattern string
	children []*node
	part string
	isLast bool
}

func Newnode() *node {
	return &node{
		children: make([]*node, 0),
	}
}

// 第一个成功匹配的节点
func (n *node) matchFirstChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isLast {
			return child
		}
	}
	return nil
}

// 所有匹配的节点
func (n *node) matchAllChild(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isLast {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) travel(list *([]*node))  {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

// 前缀树的插入,添加路由,输入路由路径和划分的路由路径以及高度, 递归算法
func (n *node) insert(pattern string, parts []string, height int)  {
	if len(parts) == height {
		// 此处报错， n是nil，n.patter出现空指针异常
		n.pattern = pattern    //最后一个节点才插入pattern，其他都是空
		return
	}

	part := parts[height]
	child := n.matchFirstChild(part)
	if child == nil {
		//没有匹配到，插入路由节点
		//child := &node{part: part, isLast: part[0]==':' || part[0]=='*'}  //我自己在这写了个:=导致新生成一个变量，空指针异常
		child = &node{part: part, isLast: part[0]==':' || part[0]=='*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// 前缀树的查找，匹配路由,是使用parts去匹配已有的node前缀树
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*"){  //最后一个节点才有pattern
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchAllChild(part)

	for _, child := range children {
		res := child.search(parts, height + 1)
		if res != nil {
			return res
		}
	}
	return nil
}