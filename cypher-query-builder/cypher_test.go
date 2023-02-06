package cypher

import (
	"testing"

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
	t1, err := NewQueryBuilder().
		Match(NewNode().SetVariable("n").AsPattern()).
		Return(ReturnConfig{Name: "n"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (n)\nRETURN n", t1)

	// pattern for t2
	edge := NewEdge().
		SetVariable("r").
		SetLabel("ACTED_IN").
		SetPath(Incoming).
		Relationship(FullRelationship{
			LeftNode:  NewNode().SetVariable("wallstreet").SetProps(Prop{Key: "title", Value: "Wall Street"}),
			RightNode: NewNode().SetVariable("actor"),
		})

	/*
		MATCH (wallstreet {title: 'Wall Street'})<-[r:ACTED_IN]-(actor)
		RETURN r.role
	*/
	t2, err := NewQueryBuilder().
		Match(edge).
		Return(ReturnConfig{Name: "r", Type: "role"}).
		Execute()
	req.NoError(err)
	req.EqualValues("MATCH (wallstreet {title: 'Wall Street'})<-[r:ACTED_IN]-(actor)\nRETURN r.role", t2)

	// pattern for t3
	edge = NewEdge().
		SetVariable("r").
		SetLabel("ACTS_IN").
		SetPath(Outgoing).
		Relationship(FullRelationship{
			LeftNode:  NewNode().SetVariable("a"),
			RightNode: NewNode(),
		})

	/*
		MATCH (a:Movie {title: 'Wall Street'})
		OPTIONAL MATCH (a)-[r:ACTS_IN]->()
		RETURN a.title, r
	*/
	t3, err := NewQueryBuilder().
		Match(NewNode().SetVariable("a").SetLabel("Movie").
			SetProps(Prop{Key: "title", Value: "Wall Street"}).
			AsPattern()).
		OptionlMath(edge).
		Return(ReturnConfig{Name: "a", Type: "title"}, ReturnConfig{Name: "r"}).
		Execute()

	req.NoError(err)
	req.EqualValues("MATCH (a:Movie {title: 'Wall Street'})\nOPTIONAL MATCH (a)-[r:ACTS_IN]->()\nRETURN a.title, r", t3)

	// CREATE (a)
	t4, err := NewQueryBuilder().Create(
		NewNode().SetVariable("a").AsPattern()).
		Execute()

	req.NoError(err)
	req.EqualValues("CREATE (a)", t4)

	// MERGE (a)
	t5, err := NewQueryBuilder().Merge(
		NewNode().SetVariable("a").AsPattern()).
		Execute()

	req.NoError(err)
	req.EqualValues("MERGE (a)", t5)

	// LIMIT 5
	t6, err := NewQueryBuilder().Limit(5).Execute()

	req.NoError(err)
	req.EqualValues("LIMIT 5", t6)

	// CALL {subquery} clause
	/*
		CALL {
		  MATCH (p)
		  RETURN p
		}
	*/
	t7, err := NewQueryBuilder().
		Call(NewQueryBuilder().
			Match(NewNode().SetVariable("p").AsPattern()).
			Return(ReturnConfig{Name: "p"})).
		Execute()

	req.NoError(err)
	req.EqualValues("CALL {\n  MATCH (p)\n  RETURN p\n}", t7)
}
