package cypher

import (
	"fmt"
	"strings"
)

// Path ...
type Path string

const (
	// Plain --
	Plain Path = "--"

	// Outgoing -->
	Outgoing Path = "-->"

	// Incoming <--
	Incoming Path = "<--"

	// Bidirectional <-->
	Bidirectional Path = "<-->"
)

// Edge ...
type Edge struct {
	Variable string
	*Label
	properties Props
	path       Path
	Condition  Condition
}

// NewEdge ...
func NewEdge() *Edge {
	return &Edge{
		Variable:   "",
		Label:      NewLabel(),
		properties: Props{},
		path:       "",
	}
}

func (e *Edge) SetVariable(name string) *Edge {
	e.Variable = name
	return e
}

func (e *Edge) SetLabel(label string) *Edge {
	e.Names = append(e.Names, label)
	return e
}

func (e *Edge) SetLabels(condition Condition, labels ...string) *Edge {
	e.Names = append(e.Names, labels...)
	e.Condition = condition
	return e
}

func (e *Edge) SetProps(props ...Prop) *Edge {
	for _, p := range props {
		e.properties[p.Key] = p.Value
	}
	return e
}

func (e *Edge) SetPath(path Path) *Edge {
	e.path = path
	return e
}

// Relationship ...
func (e Edge) Relationship(f FullRelationship) QueryPattern {
	// f.Edge = e
	return QueryPattern{
		FullRelationship: f,
		Edge:             e,
	}
}

func (e Edge) PartialRelationship(p PartialRelationship) QueryPattern {
	// p.Edge = e
	return QueryPattern{
		PartialRelationship: p,
		Edge:                e,
	}
}

func (e Edge) PartialRelationshipBuild(p PartialRelationship) string {
	//switch
	if p.LeftDirection {
		return fmt.Sprintf("%v%v", p.Node.ToCypher(), e.ToCypher())
	}

	if p.RightDirection {
		return fmt.Sprintf("%v%v", e.ToCypher(), p.Node.ToCypher())
	}

	fmt.Println("error not set type direction")
	return ""
}

func (e Edge) RelationshipBuild(f FullRelationship) string {

	// if f.Nodes == nil && len(f.Nodes) == 2 {
	// 	//error
	// 	fmt.Println("error RelationshipBuild not have nodes")
	// 	return ""
	// }

	if f.LeftNode == nil || f.RightNode == nil {
		//error
		fmt.Println("error RelationshipBuild not have nodes")
		return ""
	}

	leftNode := f.LeftNode
	rightNode := f.RightNode

	return fmt.Sprintf("%v%v%v", leftNode.ToCypher(), e.ToCypher(), rightNode.ToCypher())
}

func (e Edge) RelationshipMulti(edges ...string) string {
	res := ""
	for _, edge := range edges {
		res += fmt.Sprintf("%v", edge)
	}
	return res
}

func (e Edge) ToCypher() string {
	edge := ""

	if e.Variable != "" {
		edge += e.Variable
	}

	if len(e.Label.Names) > 0 {
		condition := ""
		if e.Label.Condition != "" {
			condition = fmt.Sprintf("%v", e.Label.Condition)
		}
		edge += fmt.Sprintf(":%v", strings.Join(e.Label.Names, condition)) + " "
	}

	if len(e.properties) > 0 {
		edge += e.properties.ToCypher()
	}

	edge = fmt.Sprintf("-[%v]-", edge)

	switch e.path {
	case Outgoing:
		edge += ">"
	case Incoming:
		edge = "<" + edge
	case Bidirectional:
		edge = "<" + edge + ">"
	case Plain:
	default:
	}

	return edge
}
