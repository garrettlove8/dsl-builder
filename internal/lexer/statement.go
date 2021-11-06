package lexer

import "bytes"

// statement_t is a single statement within a larger DSL script.
type Statement struct {
	// full is a pointer to the complete string of text making up the statement.
	Full string

	// obj is a pointer to the statements memory region.
	// As the statement is executed, this is where the results will be stored.
	Obj interface{}
}

// NewStatement constructs a Statement object, inluding the allocation of
// its memory region.
func NewStatement(text string) (*Statement, error) {
	// FIXME: Preemtively adding a FIXME note because I don't think this will work without
	// more thinking and designing.
	statement := Statement{
		Full: text,
		Obj:  new(bytes.Buffer),
	}

	return &statement, nil
}
