package script

// script_t represents a complete DSL script and provides an API for running a script.
type script struct {
	// full represents the complete script.
	full string

	// sep represents the character on which a developer intends for statements
	// to be terminated.
	sep string
}

// Run allows a developer to execute a DSL script.
func Run(userScript, sep string) {
	newScript := script{
		full: userScript,
		sep:  sep,
	}

}
