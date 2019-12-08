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

		if err := t.AddEdge(nodes[0], nodes[1]); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func (t *Tree) AddEdge(from, to string) error {
	nf := t.getNode(from)
	nt := t.getNode(to)
	if nt.Parent != nil {
		return fmt.Errorf("node %q cannot have parent %q, as it already has parent %q", to, from, nt.Parent.Name)
	}
	nf.Children = append(nf.Children, nt)
	nt.Parent = nf
	return nil
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

func (t *Tree) Distance(from, to string) int {
	nf, ok := t.nodes[from]
	if !ok {
		return -1
	}
	nt, ok := t.nodes[to]
	if !ok {
		return -1
	}

	var depth int
	parents := make(map[string]int)
	for {
		if nf != nil {
			if dparent, ok := parents[nf.Name]; ok {
				return dparent + depth
			}
			parents[nf.Name] = depth
			nf = nf.Parent
		}
		if nt != nil {
			if dparent, ok := parents[nt.Name]; ok {
				return dparent + depth
			}
			parents[nt.Name] = depth
			nt = nt.Parent
		}
		depth++
	}
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
	Parent   *Node
	Children []*Node
}

func (n *Node) totalDepths(curDepth int) int {
	total := curDepth
	for _, child := range n.Children {
		total += child.totalDepths(curDepth + 1)
	}
	return total
}
