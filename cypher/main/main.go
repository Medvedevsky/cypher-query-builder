package main

// artisan "github.com/rkrasiuk/cypher-artisan"

import (
	"fmt"
	cypher "test/neo4j"
)

func main() {

	charlie := cypher.NewNode().
		SetVariable("charlie").
		SetLabel("Person").
		SetProps(cypher.Prop{Key: "name", Value: "Martin Sheen"})

	rob := cypher.NewNode().
		SetVariable("rob").
		SetLabel("Person").
		SetProps(cypher.Prop{Key: "name", Value: "Rob Reiner"})

	edge := cypher.NewEdge().
		SetLabel("OLD FRIENDS").SetPath(cypher.Incoming).
		Relationship(cypher.FullRelationship{
			LeftNode:  charlie,
			RightNode: rob,
		})

	res, error := cypher.
		NewQueryBuilder().
		Match(edge).
		With(cypher.WithConfig{Name: "next"}).
		CALL(
			cypher.NewQueryBuilder().
				With(cypher.WithConfig{Name: "next"}).
				Match(cypher.NewNode().
					SetVariable("current").
					SetLabel("ListHead").ToPattern())).
		Return(cypher.ReturnConfig{Name: "charlie", As: "from"}, cypher.ReturnConfig{Name: "next", As: "to"}).
		Execute()

	fmt.Println(res)
	fmt.Println(error)
}
