package main

// artisan "github.com/rkrasiuk/cypherr-artisan"

import (
	"fmt"
	cypherr "test/neo4j"
)

func main() {

	// node3 := cypherr.NewNode().SetVariable("n").SetLabel("TEST").SetProps(
	// 	cypherr.Prop{Key: "flag", Value: 12.5})
	// node := cypherr.NewNode().SetVariable("n").SetLabels(cypherr.Or, "PERSON", "PEOPLE").SetProps(
	// 	cypherr.Prop{Key: "flag", Value: 12.5})

	// node2 := cypherr.NewNode().SetVariable("a").SetLabel("ATTENTION").SetProps(
	// 	cypherr.Prop{Key: "height", Value: 190})

	// pattern := cypherr.NewEdge().SetVariable("e").SetLabel("ACTION").
	// 	SetPath(cypherr.Outgoing).
	// 	Relationship(cypherr.FullRelationship{
	// 		LeftNode:  node,
	// 		RightNode: node2,
	// 	})

	// pattern2 := cypherr.NewEdge().SetVariable("f").SetLabel("WINNING").
	// 	SetPath(cypherr.Outgoing).
	// 	PartialRelationship(
	// 		cypherr.PartialRelationship{
	// 			LeftDirection: true,
	// 			Node:          node3})

	res, errors := cypherr.NewQueryBuilder().
		Match().
		Where(cypherr.WhereQuery{
			Name:            "p",
			Field:           "online",
			Check:           false,
			BooleanOperator: cypherr.EqualToOperator,
			Condition:       cypherr.AND,
		}, cypherr.WhereQuery{
			Name:            "n",
			Field:           "age",
			Check:           21,
			BooleanOperator: cypherr.EqualToOperator,
		}).Execute()

	fmt.Println(res)
	fmt.Println(errors)

}
