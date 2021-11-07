package lexer

import (
	"bytes"
	"encoding/json"
)

// statement_t is a single statement within a larger DSL script.
type Statement struct {
	// Full is a pointer to the complete string of text making up the statement.
	Full string

	// Obj is a pointer to the statements memory region.
	// As the statement is executed, this is where the results will be stored.
	Obj []*bytes.Buffer
}

// NewStatement constructs a Statement object, inluding the allocation of
// its memory region.
func NewStatement(text string) (*Statement, error) {
	// FIXME: Preemtively adding a FIXME note because I don't think this will work without
	// more thinking and designing.
	statement := Statement{
		Full: text,
		Obj:  make([]*bytes.Buffer, 0),
	}

	return &statement, nil
}

// NewStatementObj adds a new buffer to a statement's Obj slice.
func (s *Statement) NewObj() error {
	s.Obj = append(s.Obj, new(bytes.Buffer))
	return nil
}

// WriteObj overwrites the data in buffer at index i.
func (s *Statement) WriteObj(i int, data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	s.Obj[i].Write(dataBytes)
	return nil
}
