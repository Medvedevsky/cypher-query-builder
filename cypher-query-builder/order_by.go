package cypher

import (
	"errors"
	"fmt"
)

type OrderByConfig struct {
	Name   string
	Member string
	Desc   bool
	Asc    bool
}

func (o *OrderByConfig) ToString() (string, error) {
	if o.Name == "" || o.Member == "" {
		return "", errors.New("OrderByConfig - name and member have to be defined")
	}

	if o.Desc {
		return fmt.Sprintf("%s.%s DESC", o.Name, o.Member), nil
	}

	if o.Asc {
		return fmt.Sprintf("%s.%s ASC", o.Name, o.Member), nil
	}
	return fmt.Sprintf("%s.%s", o.Name, o.Member), nil
}
