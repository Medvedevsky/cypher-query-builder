package cypher

import (
	"fmt"
	"strings"
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

func (n Node) ToPattern() QueryPattern {
	return QueryPattern{OnlyNode: OnlyNode{Node: &n}}
}

func (n Node) ToCypher() string {
	node := "("

	if n.variable != "" {
		node += n.variable
	}

	if len(n.label.Names) > 0 {
		condition := ""
		if n.label.Condition != "" {
			condition = fmt.Sprintf("%v", n.label.Condition)
		}
		node += fmt.Sprintf(":%v", strings.Join(n.label.Names, condition)) + " "
	}

	if len(n.properties) > 0 {
		node += n.properties.ToCypher()
	}

	return node + ")"
}
