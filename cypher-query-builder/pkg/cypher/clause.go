package cypher

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Medvedevsky/cypher-query-builder/pkg/pattern"
)

type QueryBuilder struct {
	query  string
	errors []error
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) mapConfigToString(clauses ...pattern.QueryConfig) string {
	query := ""

	for _, clause := range clauses {
		res, error := clause.ToString()

		if error != nil {
			qb.addError(error)
		}
		query += res
	}

	return query
}

func (qb *QueryBuilder) queryPatternMap(pattern pattern.QueryPattern) string {

	if reflect.ValueOf(pattern).IsZero() {
		qb.addError(errors.New("error match QueryPattern null"))
		return ""
	}

	if !reflect.ValueOf(pattern.OnlyNode).IsZero() {
		query, err := pattern.OnlyNode.Node.ToCypher()

		if err != nil {
			qb.addError(err)
		}

		return query
	}

	if !reflect.ValueOf(pattern.PartialRelationship).IsZero() {
		p := pattern.PartialRelationship
		query, err := pattern.Edge.PartialRelationshipBuild(p)

		if err != nil {
			qb.addError(err)
		}

		return query
	}

	if !reflect.ValueOf(pattern.FullRelationship).IsZero() {
		f := pattern.FullRelationship
		query, err := pattern.Edge.RelationshipBuild(f)

		if err != nil {
			qb.addError(err)
		}

		return query
	}

	return ""
}

func (qb *QueryBuilder) queryPatternUsage(clause string, patterns ...pattern.QueryPattern) string {
	if len(patterns) == 0 {
		error := fmt.Sprintf("error %s patterns null", clause)
		qb.addError(errors.New(error))
		return ""
	}
	query := clause + " "
	for _, pattern := range patterns {
		query += qb.queryPatternMap(pattern)
	}
	query += "\n"

	return query
}

// MATCH clause
func (qb *QueryBuilder) Match(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MATCH", patterns...)
	return qb
}

// OPRIONAL MATCH clause
func (qb *QueryBuilder) OptionlMath(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("OPTIONAL MATCH", patterns...)
	return qb
}

// MERGE clause
func (qb *QueryBuilder) Merge(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("MERGE", patterns...)
	return qb
}

// CREATE clause
func (qb *QueryBuilder) Create(patterns ...pattern.QueryPattern) *QueryBuilder {
	qb.query += qb.queryPatternUsage("CREATE", patterns...)
	return qb
}

// DELETE clause
func (qb *QueryBuilder) Delete(detachDelete bool, deleteClause pattern.RemoveConfig) *QueryBuilder {
	if reflect.ValueOf(deleteClause).IsZero() {
		qb.addError(errors.New("error empty Delete clause"))
		return qb
	}

	if detachDelete {
		qb.query += "DETACH DELETE "
	} else {
		qb.query += "DELETE "
	}

	res := qb.mapConfigToString(&deleteClause)
	qb.query += res

	qb.query += "\n"

	return qb
}

// WHERE clause
func (qb *QueryBuilder) Where(whereClauses ...pattern.ConditionalConfig) *QueryBuilder {
	if len(whereClauses) == 0 {
		qb.addError(errors.New("error empty Where clause"))
		return qb
	}

	qb.query += "WHERE "
	for _, clause := range whereClauses {
		res := qb.mapConfigToString(&clause)
		qb.query += res
	}
	qb.query += "\n"

	return qb
}

// RETURN clause
func (qb *QueryBuilder) Return(returnClauses ...pattern.ReturnConfig) *QueryBuilder {
	if len(returnClauses) == 0 {
		qb.addError(errors.New("error empty Return clause"))
		return qb
	}

	query := "RETURN "
	for _, clause := range returnClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

// REMOVE clause
func (qb *QueryBuilder) Remove(removeClauses pattern.RemoveConfig) *QueryBuilder {
	if reflect.ValueOf(removeClauses).IsZero() {
		qb.addError(errors.New("error empty where clause"))
		return qb
	}

	query := "REMOVE "
	query += qb.mapConfigToString(&removeClauses)
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

func (qb *QueryBuilder) Union(all bool) *QueryBuilder {
	if all {
		qb.query += "UNION ALL\n"
		return qb
	}

	qb.query += "UNION\n"
	return qb
}

// WITH clause
func (qb *QueryBuilder) With(withClauses ...pattern.WithConfig) *QueryBuilder {
	if len(withClauses) == 0 {
		qb.addError(errors.New("error empty WITH clause"))
		return qb
	}

	query := "WITH "
	for _, clause := range withClauses {
		res := qb.mapConfigToString(&clause)
		query += res
		query += ", "
	}
	query = strings.TrimSuffix(query, ", ")
	query += "\n"
	qb.query += query

	return qb
}

// ORDER BY clause
func (qb *QueryBuilder) OrderBy(orderByClause pattern.OrderByConfig) *QueryBuilder {
	if reflect.ValueOf(orderByClause).IsZero() {
		qb.addError(errors.New("error empty OrderBy clause"))
		return qb
	}

	qb.query += "ORDER BY "
	res := qb.mapConfigToString(&orderByClause)
	qb.query += res
	qb.query += "\n"

	return qb
}

// LIMIT clause
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.query += "LIMIT " + strconv.Itoa(limit) + "\n"

	return qb
}

// CALL {subquery} clause
func (qb *QueryBuilder) Call(nqb *QueryBuilder) *QueryBuilder {
	res := "CALL {\n"
	subquery, error := nqb.Execute()
	if error != nil {
		qb.addError(error)
	}

	var buffer bytes.Buffer

	for i, rune := range subquery {
		buffer.WriteRune(rune)
		char := string(rune)

		if char == "\n" {
			if i != len(subquery)-1 {
				// buffer.WriteRune('\t')
				// таб слишком большой, по этому добавляю два пробела
				buffer.WriteRune(' ')
				buffer.WriteRune(' ')
			}
		}
	}

	subquery = buffer.String()
	res += "  " + subquery + "\n}\n"
	qb.query += res

	return qb
}

// return cypher query
func (qb *QueryBuilder) Execute() (string, error) {
	qb.query = strings.TrimSuffix(qb.query, "\n")
	return qb.query, qb.errorBuild()
}

func (q *QueryBuilder) addError(e error) {
	if q.errors == nil {
		q.errors = []error{e}
	} else {
		q.errors = append(q.errors, e)
	}
}

func (qb *QueryBuilder) errorBuild() error {
	if len(qb.errors) > 0 {
		str := "errors found: "
		for _, err := range qb.errors {
			str += err.Error() + ";"
		}

		str = strings.TrimSuffix(str, ";") + fmt.Sprintf(" -- total errors (%v)", len(qb.errors))
		return errors.New(str)
	}

	return nil
}
