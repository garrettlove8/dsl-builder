package lexer

// script_t represents a complete DSL script and provides an API for running a script.
type Script struct {
	// full represents the complete script.
	Full string

	// sep represents the character on which a developer intends for statements
	// to be terminated.
	Sep string
}

func New(full, sep string) *Script {
	return &Script{
		Full: full,
		Sep:  sep,
	}
}

func (s *Script) Parse() {}
