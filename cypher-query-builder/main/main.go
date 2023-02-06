package main

import (
	"fmt"

	"github.com/Medvedevsky/cypher-query-builder"
)

func main() {
	// example usage №1
	pNode := cypher.NewNode().SetVariable("p").SetLabel("Person").AsPattern()

	callCypher, err := cypher.NewQueryBuilder().
		Call(cypher.NewQueryBuilder().
			Match(pNode).
			Return(cypher.ReturnConfig{Name: "p"}).
			OrderBy(cypher.OrderByConfig{Name: "p", Member: "age", Asc: true}).
			Limit(1).
			Union(false).
			Match(pNode).
			Return(cypher.ReturnConfig{Name: "p"}).
			OrderBy(cypher.OrderByConfig{Name: "p", Member: "age", Desc: true}).
			Limit(1)).
		Return(cypher.ReturnConfig{Name: "p", Type: "name"}, cypher.ReturnConfig{Name: "p", Type: "age"}).
		OrderBy(cypher.OrderByConfig{Name: "p", Member: "name"}).
		Execute()

	fmt.Println(callCypher)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println()

	// example usage №2
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
		SetPath(cypher.Incoming).
		Relationship(cypher.FullRelationship{
			LeftNode:  charlie,
			RightNode: rob,
		})

	res, err :=
		cypher.NewQueryBuilder().
			Match(edge).
			With(cypher.WithConfig{Name: "next"}).
			Call(
				cypher.NewQueryBuilder().
					With(cypher.WithConfig{Name: "next"}).
					Match(cypher.NewNode().
						SetVariable("current").
						SetLabel("ListHead").AsPattern())).
			Where(cypher.ConditionalConfig{
				Name:              "rob",
				Field:             "age",
				ConditionOperator: cypher.EqualToOperator,
				Check:             21}).
			Return(cypher.ReturnConfig{Name: "charlie", As: "from"}, cypher.ReturnConfig{Name: "next", As: "to"}).
			Execute()

	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}
