package cypher

import (
	"errors"
	"fmt"
	"strings"
)

type Edge struct {
	variable   string
	label      Label
	properties Props
	path       Path
	condition  Condition
}

func NewEdge() *Edge {
	return &Edge{
		variable:   "",
		label:      Label{},
		properties: Props{},
		path:       "",
	}
}

func (e *Edge) SetVariable(name string) *Edge {
	e.variable = name
	return e
}

func (e *Edge) SetLabel(label string) *Edge {
	e.label.Names = append(e.label.Names, label)
	return e
}

func (e *Edge) SetLabels(condition Condition, labels ...string) *Edge {
	e.label.Names = append(e.label.Names, labels...)
	e.condition = condition
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

func (e Edge) Relationship(f FullRelationship) QueryPattern {
	return QueryPattern{
		FullRelationship: f,
		Edge:             e,
	}
}

func (e Edge) PartialRelationship(p PartialRelationship) QueryPattern {
	return QueryPattern{
		PartialRelationship: p,
		Edge:                e,
	}
}

func (e Edge) PartialRelationshipBuild(p PartialRelationship) (string, error) {

	if p.LeftDirection {
		leftNode, err := p.Node.ToCypher()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v%v", leftNode, e.ToCypher()), nil
	}

	if p.RightDirection {
		rightNode, err := p.Node.ToCypher()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v%v", e.ToCypher(), rightNode), nil
	}

	return "", errors.New("PartialRelationshipBuild - not set type direction")
}

func (e Edge) RelationshipBuild(f FullRelationship) (string, error) {

	if f.LeftNode == nil || f.RightNode == nil {
		return "", errors.New("RelationshipBuild - not have nodes")
	}

	leftNode, err := f.LeftNode.ToCypher()
	if err != nil {
		return "", err
	}

	rightNode, err := f.RightNode.ToCypher()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v%v%v", leftNode, e.ToCypher(), rightNode), nil
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

	if e.variable != "" {
		edge += e.variable
	}

	if len(e.label.Names) > 0 {
		condition := ""
		if e.label.Condition != "" {
			condition = fmt.Sprintf("%v", e.label.Condition)
		}
		edge += fmt.Sprintf(":%v", strings.Join(e.label.Names, condition))
	}

	if len(e.properties) > 0 {
		edge += fmt.Sprintf(" %s", e.properties.ToCypher())
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
