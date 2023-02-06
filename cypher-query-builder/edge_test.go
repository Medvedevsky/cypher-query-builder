package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEdge_ToString(t *testing.T) {
	req := require.New(t)
	var cypher string
	var err error

	//pattern
	cypher = NewEdge().SetVariable("v").SetLabel("TEST").ToCypher()
	req.EqualValues("-[v:TEST]-", cypher)

	//pattern
	cypher = NewEdge().
		SetVariable("v").
		SetLabel("TEST").
		SetProps(Prop{Key: "t", Value: 15}).
		ToCypher()
	req.EqualValues("-[v:TEST {t: 15}]-", cypher)

	//Relationship
	t1 := NewEdge().
		SetVariable("v").
		SetLabel("TEST").
		SetProps(Prop{Key: "t", Value: 15}).
		SetPath(Outgoing).
		Relationship(FullRelationship{
			LeftNode:  NewNode().SetVariable("a"),
			RightNode: NewNode().SetVariable("b"),
		})

	cypher, err = t1.Edge.RelationshipBuild(t1.FullRelationship)
	req.NoError(err)
	req.EqualValues("(a)-[v:TEST {t: 15}]->(b)", cypher)

	//PartialRelationship
	t2 := NewEdge().
		SetVariable("v").
		SetLabel("TEST").
		SetProps(Prop{Key: "t", Value: 15}).
		SetPath(Outgoing).
		PartialRelationship(PartialRelationship{
			RightDirection: true,
			Node:           NewNode().SetVariable("c"),
		})

	cypher, err = t2.Edge.PartialRelationshipBuild(t2.PartialRelationship)
	req.NoError(err)
	req.EqualValues("-[v:TEST {t: 15}]->(c)", cypher)
}
