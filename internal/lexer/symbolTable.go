package lexer

// type SymbolTable struct {
// 	General     map[string]string
// 	Identifiers map[string]string
// 	Literals    map[string]string
// 	Operators   map[string]string
// 	Delimiters  map[string]string
// 	Keywords    map[string]string
// }

// func NewSymbolTable() SymbolTable {
// 	return SymbolTable{
// 		General:     make(map[string]string),
// 		Identifiers: make(map[string]string),
// 		Literals:    make(map[string]string),
// 		Operators:   make(map[string]string),
// 		Delimiters:  make(map[string]string),
// 		Keywords:    make(map[string]string),
// 	}
// }

var SymbolTable map[string]string

func init() {
	SymbolTable = make(map[string]string)
}

func AddSymbol(tokenType, literal string) {
	SymbolTable[literal] = tokenType
}
