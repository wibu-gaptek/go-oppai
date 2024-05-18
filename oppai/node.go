package oppai

import (
	"fmt"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, wild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChildren(part)

	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || n.isWild {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	for _, child := range n.children {
		if child.part == part || child.isWild {
			if result := child.search(parts, height+1); result != nil {
				return result
			}
		}
	}

	return nil
}

func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) matchChildren(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func parsePattern(pattern string) []string {
	var parts []string
	start := 0
	isWild := false

	for i := 0; i < len(pattern); i++ {
		if pattern[i] == '/' {
			if start != i {
				parts = append(parts, pattern[start:i])
			}
			start = i + 1
		} else if pattern[i] == '*' {
			if start != i {
				parts = append(parts, pattern[start:i])
			}
			parts = append(parts, pattern[i:])
			isWild = true
			break
		}
	}

	if !isWild && start < len(pattern) {
		parts = append(parts, pattern[start:])
	}

	return parts
}
