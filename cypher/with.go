package cypher

import (
	"errors"
	"fmt"
)

type WithConfig struct {
	Name  string
	Field string
	As    string
}

func (wp *WithConfig) ToString() (string, error) {
	query := ""

	if wp.Name != "" {
		query = wp.Name

		if wp.Field != "" {
			query += fmt.Sprintf(".%s", wp.Field)
		}
	} else {
		return "", errors.New("must define a function or name")
	}

	if wp.As != "" {
		query += fmt.Sprintf(" AS %s", wp.As)
	}
	query += ", "

	return query, nil
}
