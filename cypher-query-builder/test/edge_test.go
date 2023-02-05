package test

import (
	"testing"

	"github.com/Medvedevsky/cypher-query-builder/pkg/pattern"

	"github.com/stretchr/testify/require"
)

func TestEdge_ToString(t *testing.T) {
	req := require.New(t)
	var cypher string
	var err error

	//pattern
	cypher = pattern.NewEdge().SetVariable("v").SetLabel("TEST").ToCypher()
	req.EqualValues("-[v:TEST]-", cypher)

	//pattern
	cypher = pattern.NewEdge().
		SetVariable("v").
		SetLabel("TEST").
		SetProps(pattern.Prop{Key: "t", Value: 15}).
		ToCypher()
	req.EqualValues("-[v:TEST {t: 15}]-", cypher)

	//Relationship
	t1 := pattern.NewEdge().
		SetVariable("v").
		SetLabel("TEST").
		SetProps(pattern.Prop{Key: "t", Value: 15}).
		SetPath(pattern.Outgoing).
		Relationship(pattern.FullRelationship{
			LeftNode:  pattern.NewNode().SetVariable("a"),
			RightNode: pattern.NewNode().SetVariable("b"),
		})

	cypher, err = t1.Edge.RelationshipBuild(t1.FullRelationship)
	req.NoError(err)
	req.EqualValues("(a)-[v:TEST {t: 15}]->(b)", cypher)

	//PartialRelationship
	t2 := pattern.NewEdge().
		SetVariable("v").
		SetLabel("TEST").
		SetProps(pattern.Prop{Key: "t", Value: 15}).
		SetPath(pattern.Outgoing).
		PartialRelationship(pattern.PartialRelationship{
			RightDirection: true,
			Node:           pattern.NewNode().SetVariable("c"),
		})

	cypher, err = t2.Edge.PartialRelationshipBuild(t2.PartialRelationship)
	req.NoError(err)
	req.EqualValues("-[v:TEST {t: 15}]->(c)", cypher)
}
