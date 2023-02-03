package pattern

import (
	"errors"
	"fmt"
)

type ReturnConfig struct {
	Name string
	Type string
	As   string
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

	if r.As != "" {
		query += fmt.Sprintf(" AS %s", r.As)
	}
	query += ", "

	return query, nil
}
