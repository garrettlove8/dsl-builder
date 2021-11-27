package lexer

import (
	"errors"
)

type SymbolTable struct {
	General     map[string]string
	Identifiers map[string]string
	Literals    map[string]string
	Operators   map[string]string
	Delimiters  map[string]string
	Keywords    map[string]string
}

func NewSymbolTable() SymbolTable {
	return SymbolTable{
		General:     make(map[string]string),
		Identifiers: make(map[string]string),
		Literals:    make(map[string]string),
		Operators:   make(map[string]string),
		Delimiters:  make(map[string]string),
		Keywords:    make(map[string]string),
	}
}

func (st *SymbolTable) AddSymbol(tokenCat, tokenType, literal string) error {
	switch tokenCat {
	case "General":
		st.General[literal] = tokenType
	case "Identifiers":
		st.Identifiers[literal] = tokenType
	case "Literals":
		st.Literals[literal] = tokenType
	case "Operators":
		st.Operators[literal] = tokenType
	case "Delimiters":
		st.Delimiters[literal] = tokenType
	case "Keywords":
		st.Keywords[literal] = tokenType
	default:
		return errors.New("Unable to add symbol, not of a supported category")
	}

	return nil
}
