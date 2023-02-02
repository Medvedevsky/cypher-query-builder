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
		SetLabel("OLD FRIENDS").
		Relationship(cypher.FullRelationship{
			LeftNode:  rob,
			RightNode: charlie,
		})

	res, error := cypher.
		NewQueryBuilder().
		Match(charlie.ToPattern()).
		Match(rob.ToPattern()).
		Create(edge).
		Execute()

	fmt.Println(res)
	fmt.Println(error)

}
