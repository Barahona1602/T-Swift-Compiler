package expressions

import (
	"Server2/environment"
	"Server2/generator"
)

type Count struct {
	Lin     int
	Col     int
	ArrayID string
}

func NewCount(lin int, col int, arrayID string) Count {
	count := Count{Lin: lin, Col: col, ArrayID: arrayID}
	return count
}

func (c Count) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	return result
}
