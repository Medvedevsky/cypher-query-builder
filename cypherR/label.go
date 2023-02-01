package cypherr

type Label struct {
	Names     []string
	Condition Condition
}

func NewLabel() *Label {
	return &Label{}
}
