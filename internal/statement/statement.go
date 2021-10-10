package statement

// statement_t is a single statement within a larger DSL script.
type statement struct {
	// full is a pointer to the complete string of text making up the statement.
	full string

	// obj is a pointer to the statements memory region.
	// As the statement is executed, this is where the results will be stored.
	obj interface{}
}

// newStatement constructs a statement object - inluding the allocation of
// its memory region.
func newStatement(text string) (*statement, error) {
	statement := statement{}

	// // strcpy(statement.full, text);
	// statement.obj = stalloc();

	return &statement, nil
}

// deleteStatement destructs a statement - including its memory region.
func deleteStatement(statement *statement) {
	// stfree(statement);
}

// parseStatement creates the token structures found in the statement base
// on the lexical units.
func parseStatement(statement *statement) {}

// executeStatement runs the statement's logic by running each tokens function
// keeping track of the effects of each on the statement's memory region.
func executeStatement(statement *statement) {}

func stfree(statement *statement) {
	// free(statement->obj);
}

func stalloc() {
	// return calloc(1, MODEL_OBJ_CAPACITY);
}
