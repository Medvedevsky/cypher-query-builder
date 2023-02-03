package cypher

import (
	"errors"
	"fmt"
)

type ReturnConfig struct {
	Name string
	Type string
}

func (r *ReturnConfig) ToString() (string, error) {
	if r.Name == "" {
		return "", errors.New("error Return clause: name must be defined")
	}

	query := ""

	if r.Type != "" {
		query += fmt.Sprintf("%s.%s", r.Name, r.Type)
	} else {
		query += r.Name
	}
	query += ", "

	return query, nil
}
