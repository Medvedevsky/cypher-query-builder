package main

import (
	"fmt"
	"test/neo4j/pkg/cypher"
	"test/neo4j/pkg/pattern"
)

func main() {

	charlie := pattern.NewNode().
		SetVariable("charlie").
		SetLabel("Person").
		SetProps(pattern.Prop{Key: "name", Value: "Martin Sheen"})

	rob := pattern.NewNode().
		SetVariable("rob").
		SetLabel("Person").
		SetProps(pattern.Prop{Key: "name", Value: "Rob Reiner"})

	edge := pattern.NewEdge().
		SetLabel("OLD FRIENDS").SetPath(pattern.Incoming).
		Relationship(pattern.FullRelationship{
			LeftNode:  charlie,
			RightNode: rob,
		})

	res, error := cypher.
		NewQueryBuilder().
		Match(edge).
		With(pattern.WithConfig{Name: "next"}).
		CALL(
			cypher.NewQueryBuilder().
				With(pattern.WithConfig{Name: "next"}).
				Match(pattern.NewNode().
					SetVariable("current").
					SetLabel("ListHead").ToPattern())).
		Return(pattern.ReturnConfig{Name: "charlie", As: "from"}, pattern.ReturnConfig{Name: "next", As: "to"}).
		Execute()

	fmt.Println(res)
	fmt.Println(error)
}
