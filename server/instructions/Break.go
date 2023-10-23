package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
)

type Break struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewBreak(lin int, col int, exp interfaces.Expression) Break {
	breakInstr := Break{lin, col, exp}
	return breakInstr
}

func (p Break) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	if p.Expresion == nil {
		result = environment.NewValue("", false, environment.NIL, false, true, false)

	} else {
		tmpVal := p.Expresion.Ejecutar(ast, env, gen)
		tmpVal.BreakFlag = true
		result = tmpVal
	}

	return result
}
