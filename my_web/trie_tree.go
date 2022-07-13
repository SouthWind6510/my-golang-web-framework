package my_web

type node struct {
	path     string  // 待匹配路由
	part     string  // 路由节点
	children []*node // 子节点
	isWild   bool    // 是否模糊匹配，part 含有 : 或 * 时 true
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
	result := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			result = append(result, child)
		}
	}
	return result
}

func (n *node) insert(path string, parts []string, height int) {
	if len(parts) == height {
		n.path = path
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	// 新建节点
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(path, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height {
		if n.path == "" {
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
