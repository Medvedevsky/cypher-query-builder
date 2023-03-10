package cypher

import (
	"errors"
	"fmt"
)

type Node struct {
	variable   string
	label      Label
	properties Props
}

func NewNode() *Node {
	return &Node{
		label:      Label{},
		properties: Props{},
		variable:   "",
	}
}

func (n *Node) SetVariable(name string) *Node {
	n.variable = name
	return n
}

func (n *Node) SetProps(props ...Prop) *Node {
	for _, p := range props {
		n.properties[p.Key] = p.Value
	}
	return n
}

func (n *Node) SetLabel(label string) *Node {
	n.label.Names = append(n.label.Names, label)
	return n
}

func (n *Node) SetLabels(condition Condition, labels ...string) *Node {
	n.label.Names = append(n.label.Names, labels...)
	n.label.Condition = condition
	return n
}

func (n Node) AsPattern() QueryPattern {
	return QueryPattern{OnlyNode: OnlyNode{Node: &n}}
}

func (n Node) ToCypher() (string, error) {

	if n.variable == "" && len(n.label.Names) > 0 {
		return "", errors.New("Node must have a variable with at least one label")
	}

	node := "("

	if n.variable != "" {
		node += n.variable
	}

	node += n.label.ToCypher()

	if len(n.properties) > 0 {
		node += fmt.Sprintf(" %s", n.properties.ToCypher())
	}
	node += ")"

	return node, nil
}
