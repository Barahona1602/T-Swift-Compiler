package expressions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
)

type IsEmpty struct {
	Lin   int
	Col   int
	Array interfaces.Expression
}

func NewIsEmpty(lin int, col int, array interfaces.Expression) IsEmpty {
	exp := IsEmpty{lin, col, array}
	return exp
}

func (p IsEmpty) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	tempArray := p.Array.Ejecutar(ast, env, gen)
	zeroLabel := gen.NewTemp()
	gen.AddAssign(zeroLabel, "0")
	trueLabel := gen.NewLabel()
	falseLabel := gen.NewLabel()
	gen.AddIf(tempArray.Value, zeroLabel, "!=", trueLabel)
	gen.AddGoto(falseLabel)
	result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
	result.TrueLabel = append(result.TrueLabel, trueLabel)
	result.FalseLabel = append(result.FalseLabel, falseLabel)
	return result
}
