package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEdges(t *testing.T) {
	tree := NewTree()
	tree.AddEdge("A", "B")
	tree.AddEdge("A", "C")
	tree.AddEdge("B", "D")

	assert.Equal(t, []string{"B", "C"}, tree.Children("A"))
	assert.Equal(t, []string{"D"}, tree.Children("B"))
	assert.Nil(t, tree.Children("C"))
	assert.Nil(t, tree.Children("D"))
}

const example = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

func TestTreeFromString(t *testing.T) {
	tree, err := TreeFromString(example)
	assert.NoError(t, err)

	assert.Equal(t, []string{"B"}, tree.Children("COM"))
	assert.Equal(t, []string{"C", "G"}, tree.Children("B"))
	assert.Equal(t, []string{"F", "J"}, tree.Children("E"))
	assert.Nil(t, tree.Children("H"))
}

func TestTotalDepths(t *testing.T) {
	tree, err := TreeFromString(example)
	assert.NoError(t, err)

	assert.Equal(t, 42, tree.TotalDepths("COM"))
}

const example2 = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`

func TestDistance(t *testing.T) {
	tree, err := TreeFromString(example2)
	assert.NoError(t, err)

	assert.Equal(t, 6, tree.Distance("YOU", "SAN"))
	assert.Equal(t, 5, tree.Distance("L", "C"))
}
