package script

// script_t represents a complete DSL script and provides an API for running a script.
type script struct {
	// full is a pointer which points to the complete script.
	full string

	// sep is a character array which holds the character on which a developer intends
	// for statements to be terminated.
	sep [2]string
}

// run allows a developer to execute a DSL script.
func run(script *script) {}
