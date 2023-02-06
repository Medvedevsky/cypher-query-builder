package cypher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveConfig_ToString(t *testing.T) {
	req := require.New(t)
	var err error
	var res string

	//labels and field can not both be defined
	t1 := RemoveConfig{
		Labels: []string{"test"},
		Field:  "abs",
	}
	_, err = t1.ToString()
	req.NotNil(err)

	//labels and field can not both be defined
	t2 := RemoveConfig{
		Name:   "abc",
		Labels: []string{"test"},
		Field:  "abs",
	}
	_, err = t2.ToString()
	req.NotNil(err)

	//all not defined
	t3 := RemoveConfig{}
	_, err = t3.ToString()
	req.NotNil(err)

	//pattern name
	t4 := RemoveConfig{
		Name: "n",
	}
	res, err = t4.ToString()
	req.Nil(err)
	req.EqualValues("n", res)

	//pattern with name and prop
	t5 := RemoveConfig{
		Name:  "n",
		Field: "prop",
	}
	res, err = t5.ToString()
	req.Nil(err)
	req.EqualValues("n.prop", res)

	//pattern with name and label
	t6 := RemoveConfig{
		Name:   "n",
		Labels: []string{"TEST"},
	}
	res, err = t6.ToString()
	req.Nil(err)
	req.EqualValues("n:TEST", res)

	// clause DELETE
	t7, err := NewQueryBuilder().Delete(false, RemoveConfig{Name: "n"}).Execute()
	req.Nil(err)
	req.EqualValues("DELETE n", t7)

	// clause DETACH DELETE
	t8, err := NewQueryBuilder().Delete(true, RemoveConfig{Name: "n"}).Execute()
	req.Nil(err)
	req.EqualValues("DETACH DELETE n", t8)

	// clause REMOVE
	t9, err := NewQueryBuilder().Remove(RemoveConfig{Name: "n"}).Execute()
	req.Nil(err)
	req.EqualValues("REMOVE n", t9)
}
