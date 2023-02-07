package cypher

import (
	"errors"
	"fmt"
)

type ConditionalConfig struct {
	ConditionOperator BooleanOperator
	Condition         Condition

	ConditionFunction string
	Name              string
	Field             string
	Label             string
	Check             interface{}
}

func (condition *ConditionalConfig) ToString() (string, error) {
	if condition.Name == "" {
		return "", errors.New("ConditionalConfig - var name can not be empty")
	}

	if condition.Field != "" && condition.Label != "" {
		return "", errors.New("ConditionalConfig - only one of field or label can be set")
	}

	if condition.Check != nil && condition.ConditionOperator == "" {
		return "", errors.New("ConditionalConfig - condition operator can not be empty with var check")
	}

	if condition.Label != "" && condition.ConditionFunction != "" {
		return "", errors.New("ConditionalConfig - only one of label or condition function can be set")
	}

	query := ""

	//build the fields
	if condition.Field != "" {
		query += fmt.Sprintf("%s.%s", condition.Name, condition.Field)
	} else if condition.Label != "" {
		// or label
		return fmt.Sprintf("%s:%s", condition.Name, condition.Label), nil
	} else {
		query = condition.Name
	}

	//build the operators
	if condition.ConditionOperator != "" {
		query += fmt.Sprintf(" %s", condition.ConditionOperator)
	} else if condition.ConditionFunction != "" {
		//if its a condition function
		return fmt.Sprintf("%s(%s)", condition.ConditionFunction, query), nil
	}

	if condition.Check != nil {
		switch condition.Check.(type) {
		case string:
			query += fmt.Sprintf(" '%s'", condition.Check)
		default:
			query += fmt.Sprintf(" %v", condition.Check)
		}
	}

	// if condition config not one
	if condition.Condition != "" {
		query += fmt.Sprintf(" %s ", condition.Condition)
	}

	return query, nil
}
