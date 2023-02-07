# Cypher Query Builder

Cypher Query Builder for Neo4j

## Example usage

``` go
query, err := NewQueryBuilder().
       Match(NewNode().SetVariable("n").AsPattern()).
       Return(ReturnConfig{Name: "n"}).
       Execute()
```

```
MATCH (n)
RETURN n
```

``` go
node := NewNode().SetVariable("n").SetLabel("My Label").AsPattern()
query, err := NewQueryBuilder().
       Match(node).
       Return(ReturnConfig{Name: "n"}).
       Execute()
```

```
MATCH (n:`My Label`)
RETURN n
```
``` go
node := NewNode().SetVariable("n").SetLabels(Сolon, "My Label", "Our Label").AsPattern()
query, err := NewQueryBuilder().
    Match(node).
    Where(ConditionalConfig{
    	Name:              "n",
    	Field:             "attr1",
    	ConditionOperator: EqualToOperator,
    	Check:             "value 1",
    	Condition:         AND}, ConditionalConfig{
    	Name:              "n",
    	Field:             "attr2",
    	ConditionOperator: EqualToOperator,
    	Check:             "value 2",
    }).
    Return(ReturnConfig{Name: "n"}).
    Execute()
```

```
MATCH (n:`My Label`:`Our Label`)
WHERE n.attr1 = 'value 1' AND n.attr2 = 'value 2'
RETURN n
```

``` go
edge := NewEdge().SetPath(Outgoing).Relationship(FullRelationship{
    LeftNode:  NewNode().SetVariable("n"),
    RightNode: NewNode().SetVariable("m"),
})
query, err := NewQueryBuilder().
    Match(NewNode().SetVariable("n").SetLabel("My Label").AsPattern()).
    OptionlMath(edge).
    Return(ReturnConfig{Name: "n"}, ReturnConfig{Name: "m"}).
    Execute()
```

```
MATCH (n:`My Label`)
OPTIONAL MATCH (n)-[]->(m)
RETURN n, m
```

```go
node := NewNode().SetVariable("n").SetLabel("My Label").AsPattern()
query, err := NewQueryBuilder().
       Match(node).
       Return(ReturnConfig{Name: "n", Type: "fitstProperty"}).
       Execute()
```

```
MATCH (n:`My Label`)
RETURN n.fitstProperty
```

## Implemented Query Clauses
    + Match
    + Optional Match
    + Merge
    + Where
    + With
    + Return
    + OrderBy
    + Limit
    + Create
    + Delete
    + Remove
    + Union
    + Call {subquery}
