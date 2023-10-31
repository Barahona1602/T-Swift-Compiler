package expressions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
)

type Count struct {
	Lin   int
	Col   int
	Array interfaces.Expression
}

func NewCount(lin int, col int, array interfaces.Expression) Count {
	exp := Count{lin, col, array}
	return exp
}

func (p Count) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	tempArray := p.Array.Ejecutar(ast, env, gen)
	size := tempArray.Value
	count := gen.NewTemp()
	gen.AddGetHeap(count, "(int)"+size)
	result = environment.NewValue(count, true, environment.INTEGER, false, false, false)
	return result
}
