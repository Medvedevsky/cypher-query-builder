package cypher

import (
	"fmt"
	"reflect"
	"strconv"
)

// *QueryBuilder ...
type QueryBuilder struct {
	query  string
	errors []error
}

// NewQueryBuilder ...
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) queryPatternMap(pattern QueryPattern) string {
	query := ""

	if reflect.ValueOf(pattern).IsZero() {
		// error
		qb.addError(fmt.Errorf("error match QueryPattern null"))
		return ""
	}

	if !reflect.ValueOf(pattern.OnlyNode).IsZero() {
		query += pattern.OnlyNode.Node.ToCypher()

		return query
	}

	if !reflect.ValueOf(pattern.PartialRelationship).IsZero() {
		p := pattern.PartialRelationship

		query += pattern.Edge.PartialRelationshipBuild(p)
		return query
	}

	if !reflect.ValueOf(pattern.FullRelationship).IsZero() {
		f := pattern.FullRelationship
		query += pattern.Edge.RelationshipBuild(f)

		return query
	}

	return ""
}

func (qb *QueryBuilder) queryPatternUsage(clauses string, patterns ...QueryPattern) string {
	if len(patterns) == 0 {
		// error
		error := fmt.Sprintf("error %s patterns null", clauses)
		qb.addError(fmt.Errorf(error))
		return ""
	}
	query := clauses + " "
	for _, pattern := range patterns {
		query += qb.queryPatternMap(pattern)
	}
	query += "\n"

	return query
}

func (qb *QueryBuilder) Match(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MATCH", patterns...)
	return qb
}

func (qb *QueryBuilder) OptionlMath(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("OPTIONAL MATH", patterns...)
	return qb
}

func (qb *QueryBuilder) Merge(patterns ...QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MERGE", patterns...)
	return qb
}

func (qb *QueryBuilder) Delete(detchDelete bool, pattern ConditionalQuery) *QueryBuilder {
	p := fmt.Sprintf("%v", pattern.Name)
	if detchDelete {
		qb.query += "DETACH DELETE " + p
	} else {
		qb.query += "DELETE " + p
	}
	qb.query += "\n"

	return qb
}

func (qb *QueryBuilder) Where(whereClauses ...ConditionalQuery) *QueryBuilder {
	if len(whereClauses) == 0 {
		qb.addError(fmt.Errorf("error empty where clause"))
		return qb
	}

	qb.query += "WHERE "
	qb.query += whereMap(whereClauses...)
	qb.query += "\n"

	return qb
}

func (qb *QueryBuilder) Return(returnClauses ...ConditionalQuery) *QueryBuilder {
	if len(returnClauses) == 0 {
		qb.addError(fmt.Errorf("error empty where clause"))
		return qb
	}
	qb.query += "RETRUN "
	qb.query += conditionalMap(returnClauses...)
	qb.query += "\n"

	return qb
	// return *QueryBuilder{
	// 	qb.query + "RETURN " + strings.Join(returnClauses, ", "),
	// }
}

func (qb *QueryBuilder) With(withClauses ...ConditionalQuery) *QueryBuilder {
	if len(withClauses) == 0 {
		qb.addError(fmt.Errorf("error empty WITH clause"))
		return qb
	}
	qb.query += "WITH "
	qb.query += withMap(withClauses...)
	qb.query += "\n"

	return qb
}

func (qb *QueryBuilder) OrderBy(orderByClause ConditionalQuery) *QueryBuilder {
	if reflect.ValueOf(orderByClause).IsZero() {
		qb.addError(fmt.Errorf("error empty OrderBy clause"))
		return qb
	}

	qb.query += "ORDER BY "
	if orderByClause.Field != "" {
		qb.query += fmt.Sprintf("%v.%v ", orderByClause.Name, orderByClause.Field) + string(orderByClause.OrderByOperator)
	} else {
		qb.query += fmt.Sprintf("%v ", orderByClause.Name) + string(orderByClause.OrderByOperator)
	}
	qb.query += "\n"

	return qb
}

// Limit ...
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.query += "LIMIT " + strconv.Itoa(limit) + "\n"

	return qb
}

// Execute ...
func (qb *QueryBuilder) Execute() (string, []error) {
	return qb.query, qb.errors
}

// As ...
func As(initial, alias string) string {
	return fmt.Sprintf("%v AS %v", initial, alias)
}

// Assign ...
func Assign(name, pattern string) string {
	return fmt.Sprintf("%v = %v", name, pattern)
}

func whereMap(conditions ...ConditionalQuery) string {
	query := ""
	for _, condition := range conditions {
		query += fmt.Sprintf("%v.%v %v %v %v ", condition.Name, condition.Field, condition.BooleanOperator, condition.Check, condition.Condition)
	}

	return query
}

func withMap(conditions ...ConditionalQuery) string {
	query := ""
	for i, condition := range conditions {
		if i != len(conditions)-1 {
			query += fmt.Sprintf("%v", condition.Name) + ","
			continue
		}
		query += fmt.Sprintf("%v", condition.Name)
	}

	return query
}

func conditionalMap(conditions ...ConditionalQuery) string {
	res := ""
	for i, condition := range conditions {
		if i != len(conditions)-1 {
			if condition.Field != "" {
				res += fmt.Sprintf("%v.%v", condition.Name, condition.Field) + ", "
			} else {
				res += fmt.Sprintf("%v", condition.Name) + ", "
			}
			continue
		}

		if condition.Field != "" {
			res += fmt.Sprintf("%v.%v", condition.Name, condition.Field)
		} else {
			res += fmt.Sprintf("%v", condition.Name)
		}
	}
	return res
}

func (q *QueryBuilder) addError(e error) {
	if q.errors == nil {
		q.errors = []error{e}
	} else {
		q.errors = append(q.errors, e)
	}
}
