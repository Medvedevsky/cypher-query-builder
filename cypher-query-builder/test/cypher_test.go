package test

import (
	"testing"

	"github.com/Medvedevsky/cypher-query-builder/pkg/cypher"
	"github.com/Medvedevsky/cypher-query-builder/pkg/pattern"

	"github.com/stretchr/testify/require"
)

func TestCypher_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	// var res string

	/*
		MATCH (n)
		RETURN n
	*/
	t1, err := cypher.NewQueryBuilder().
		Match(pattern.NewNode().SetVariable("n").AsPattern()).
		Return(pattern.ReturnConfig{Name: "n"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n)\nRETURN n", t1)

	// pattern for t2
	edge := pattern.NewEdge().
		SetVariable("r").
		SetLabel("ACTED_IN").
		SetPath(pattern.Incoming).
		Relationship(pattern.FullRelationship{
			LeftNode:  pattern.NewNode().SetVariable("wallstreet").SetProps(pattern.Prop{Key: "title", Value: "Wall Street"}),
			RightNode: pattern.NewNode().SetVariable("actor"),
		})

	/*
		MATCH (wallstreet {title: 'Wall Street'})<-[r:ACTED_IN]-(actor)
		RETURN r.role
	*/
	t2, err := cypher.NewQueryBuilder().
		Match(edge).
		Return(pattern.ReturnConfig{Name: "r", Type: "role"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (wallstreet {title: 'Wall Street'})<-[r:ACTED_IN]-(actor)\nRETURN r.role", t2)

	// pattern for t3
	edge = pattern.NewEdge().
		SetVariable("r").
		SetLabel("ACTS_IN").
		SetPath(pattern.Outgoing).
		Relationship(pattern.FullRelationship{
			LeftNode:  pattern.NewNode().SetVariable("a"),
			RightNode: pattern.NewNode(),
		})

	/*
		MATCH (a:Movie {title: 'Wall Street'})
		OPTIONAL MATCH (a)-[r:ACTS_IN]->()
		RETURN a.title, r
	*/
	t3, err := cypher.NewQueryBuilder().
		Match(pattern.NewNode().SetVariable("a").SetLabel("Movie").
			SetProps(pattern.Prop{Key: "title", Value: "Wall Street"}).
			AsPattern()).
		OptionlMath(edge).
		Return(pattern.ReturnConfig{Name: "a", Type: "title"}, pattern.ReturnConfig{Name: "r"}).
		Execute()

	req.NoError(err)
	req.EqualValues("MATCH (a:Movie {title: 'Wall Street'})\nOPTIONAL MATCH (a)-[r:ACTS_IN]->()\nRETURN a.title, r", t3)

	// CREATE (a)
	t4, err := cypher.NewQueryBuilder().Create(
		pattern.NewNode().SetVariable("a").AsPattern()).
		Execute()

	req.NoError(err)
	req.EqualValues("CREATE (a)", t4)

	// MERGE (a)
	t5, err := cypher.NewQueryBuilder().Merge(
		pattern.NewNode().SetVariable("a").AsPattern()).
		Execute()

	req.NoError(err)
	req.EqualValues("MERGE (a)", t5)

	// LIMIT 5
	t6, err := cypher.NewQueryBuilder().Limit(5).Execute()

	req.NoError(err)
	req.EqualValues("LIMIT 5", t6)

	// CALL {subquery} clause
	/*
		CALL {
		  MATCH (p)
		  RETURN p
		}
	*/
	t7, err := cypher.NewQueryBuilder().
		Call(cypher.NewQueryBuilder().
			Match(pattern.NewNode().SetVariable("p").AsPattern()).
			Return(pattern.ReturnConfig{Name: "p"})).
		Execute()

	req.NoError(err)
	req.EqualValues("CALL {\n  MATCH (p)\n  RETURN p\n}", t7)
}
