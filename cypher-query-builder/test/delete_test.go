package test

import (
	"test/neo4j/pkg/pattern"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var cypher string

	//labels and field can not both be defined
	t1 := pattern.RemoveConfig{
		Labels: []string{"test"},
		Field:  "abs",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//labels and field can not both be defined
	t2 := pattern.RemoveConfig{
		Name:   "abc",
		Labels: []string{"test"},
		Field:  "abs",
	}
	_, err = t2.ToString()
	req.NotNil(err)

	//all not defined
	t3 := pattern.RemoveConfig{}
	_, err = t3.ToString()
	req.NotNil(err)

	//pattern name
	t4 := pattern.RemoveConfig{
		Name: "n",
	}
	cypher, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n", cypher)

	//pattern with name and prop
	t5 := pattern.RemoveConfig{
		Name:  "n",
		Field: "prop",
	}
	cypher, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.prop", cypher)

	//pattern with name and label
	t6 := pattern.RemoveConfig{
		Name:   "n",
		Labels: []string{"TEST"},
	}
	cypher, err = t6.ToString()
	req.Nil(err)
	req.EqualValues("n:TEST", cypher)
}
