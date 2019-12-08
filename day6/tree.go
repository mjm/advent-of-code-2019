package day6

import (
	"fmt"
	"strings"
)

type Tree struct {
	nodes map[string]*Node
}

func NewTree() *Tree {
	return &Tree{
		nodes: make(map[string]*Node),
	}
}

func TreeFromString(s string) (*Tree, error) {
	t := NewTree()
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		nodes := strings.Split(line, ")")
		if len(nodes) != 2 {
			return nil, fmt.Errorf("line %q does not have two nodes separated by ')'", line)
		}

		t.AddEdge(nodes[0], nodes[1])
	}
	return t, nil
}

func (t *Tree) AddEdge(from, to string) {
	nf := t.getNode(from)
	nt := t.getNode(to)
	nf.Children = append(nf.Children, nt)
}

func (t *Tree) Children(name string) []string {
	var names []string
	n, ok := t.nodes[name]
	if !ok {
		return names
	}

	for _, child := range n.Children {
		names = append(names, child.Name)
	}
	return names
}

func (t *Tree) TotalDepths(from string) int {
	n, ok := t.nodes[from]
	if !ok {
		return 0
	}

	return n.totalDepths(0)
}

func (t *Tree) getNode(name string) *Node {
	n, ok := t.nodes[name]
	if !ok {
		n = &Node{Name: name}
		t.nodes[name] = n
	}
	return n
}

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) totalDepths(curDepth int) int {
	total := curDepth
	for _, child := range n.Children {
		total += child.totalDepths(curDepth + 1)
	}
	return total
}
