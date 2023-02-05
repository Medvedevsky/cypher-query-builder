package main

import (
	"fmt"

	"github.com/Medvedevsky/cypher-query-builder/pkg/cypher"
	"github.com/Medvedevsky/cypher-query-builder/pkg/pattern"
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
		SetLabel("OLD FRIENDS").
		SetPath(pattern.Incoming).
		Relationship(pattern.FullRelationship{
			LeftNode:  charlie,
			RightNode: rob,
		})

	res, err := cypher.
		NewQueryBuilder().
		Match(edge).
		With(pattern.WithConfig{Name: "next"}).
		Call(
			cypher.NewQueryBuilder().
				With(pattern.WithConfig{Name: "next"}).
				Match(pattern.NewNode().
					SetVariable("current").
					SetLabel("ListHead").AsPattern())).
		Where(pattern.ConditionalConfig{
			Name:              "a",
			Field:             "label",
			ConditionFunction: "tfunc"}).
		Return(pattern.ReturnConfig{Name: "charlie", As: "from"}, pattern.ReturnConfig{Name: "next", As: "to"}).
		Execute()

	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()

	pNode := pattern.NewNode().SetVariable("p").SetLabel("Person").AsPattern()

	callCypher, err := cypher.NewQueryBuilder().
		Call(cypher.NewQueryBuilder().
			Match(pNode).
			Return(pattern.ReturnConfig{Name: "p"}).
			OrderBy(pattern.OrderByConfig{Name: "p", Member: "age", Asc: true}).
			Limit(1).
			Union(false).
			Match(pNode).
			Return(pattern.ReturnConfig{Name: "p"}).
			OrderBy(pattern.OrderByConfig{Name: "p", Member: "age", Desc: true}).
			Limit(1)).
		Return(pattern.ReturnConfig{Name: "p", Type: "name"}, pattern.ReturnConfig{Name: "p", Type: "age"}).
		OrderBy(pattern.OrderByConfig{Name: "p", Member: "name"}).
		Execute()

	fmt.Println(callCypher)

	if err != nil {
		fmt.Println(err)
	}

	// github.com/Medvedevsky/cypher-query-builder
	// go get github.com/stretchr/testify
}
