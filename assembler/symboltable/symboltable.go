// Package symboltable defines symbol table with predefined symbols and helper methods
package symboltable

type SymbolTable struct {
	Symbols map[string]int
	currentVarLocation int
}

// Gets new SymbolTable with predefined symbols
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		map[string]int{
			"SP": 0,
			"LCL": 1,
			"ARG": 2,
			"THIS": 3,
			"THAT": 4,
			"R0": 0,
			"R1": 1,
			"R2": 2,
			"R3": 3,
			"R4": 4,
			"R5": 5,
			"R6": 6,
			"R7": 7,
			"R8": 8,
			"R9": 9,
			"R10": 10,
			"R11": 11,
			"R12": 12,
			"R13": 13,
			"R14": 14,
			"R15": 15,
			"SCREEN": 16384,
			"KBD": 24576,
		},
		16,
	}
}

// Adds new label to the table(used in first pass)
func (table *SymbolTable) AddLabelToTable(symbol string, address int) {
	table.Symbols[symbol] = address
}

// Adds new variable to the table(used in second pass)
func (table *SymbolTable) AddVariableToTable(symbol string) {
	table.Symbols[symbol] = table.currentVarLocation
	table.currentVarLocation++
}

// Helper wrapper to get(possibly after adding) symbols address
func (table *SymbolTable) GetSymbol(symbol string) int {
	value, isPresent := table.Symbols[symbol]
	if isPresent {
		return value
	}

	table.AddVariableToTable(symbol)
	return table.Symbols[symbol]
}