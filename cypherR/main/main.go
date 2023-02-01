package main

// artisan "github.com/rkrasiuk/cypher-artisan"

import (
	"fmt"
	cypher "test/neo4j"
)

func main() {

	node3 := cypher.NewNode().SetVariable("n").SetLabel("TEST").SetProps(
		cypher.Prop{Key: "flag", Value: 12.5})

	node := cypher.NewNode().SetVariable("n").SetLabels(cypher.Or, "PERSON", "PEOPLE").SetProps(
		cypher.Prop{Key: "flag", Value: 12.5})

	node2 := cypher.NewNode().SetVariable("a").SetLabel("ATTENTION").SetProps(
		cypher.Prop{Key: "height", Value: 190})

	pattern := cypher.NewEdge().SetVariable("e").SetLabel("ACTION").
		SetPath(cypher.Outgoing).
		Relationship(cypher.FullRelationship{
			LeftNode:  node,
			RightNode: node2,
		})

	pattern2 := cypher.NewEdge().SetVariable("f").SetLabel("WINNING").
		SetPath(cypher.Incoming).
		Relationship(
			cypher.FullRelationship{
				LeftNode:  node2,
				RightNode: node3,
			})

	res, errors := cypher.NewQueryBuilder().
		Match(pattern).
		Match(node3.ToPattern()).
		Merge(pattern2).
		Where(cypher.ConditionalQuery{
			Name:            "p",
			Field:           "online",
			Check:           false,
			BooleanOperator: cypher.EqualToOperator,
			Condition:       cypher.AND,
		}, cypher.ConditionalQuery{
			Name:            "n",
			Field:           "age",
			Check:           21,
			BooleanOperator: cypher.EqualToOperator,
		}).
		With(cypher.ConditionalQuery{Name: "t"}, cypher.ConditionalQuery{Name: "a"}).
		OrderBy(cypher.ConditionalQuery{Name: "t", Field: "peop", OrderByOperator: cypher.Desc}).
		Delete(true, cypher.ConditionalQuery{Name: "t"}).
		Limit(5).
		Return(cypher.ConditionalQuery{Name: "t", Field: "prop"}).
		Execute()

	fmt.Println(res)
	fmt.Println(errors)

}
