package instructions

import (
	"Server2/environment"
	"Server2/generator"
)

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	ContinueInstr := Continue{lin, col}
	return ContinueInstr
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	result = environment.NewValue("", false, environment.NIL, false, false, true)
	return result
}
