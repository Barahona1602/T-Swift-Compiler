package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
)

type ArrayRemoveLast struct {
	Lin   int
	Col   int
	Array interfaces.Expression
}

func NewRemoveLast(lin int, col int, array interfaces.Expression) ArrayRemoveLast {
	exp := ArrayRemoveLast{lin, col, array}
	return exp
}

func (p ArrayRemoveLast) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	tempArray := p.Array.Ejecutar(ast, env, gen)
	sizePtr := tempArray.Value
	count := gen.NewTemp()
	gen.AddComment("Eliminar el último elemento del arreglo")
	gen.AddGetHeap(count, "(int)"+sizePtr)
	lastElementPtr := gen.NewTemp()
	gen.AddExpression(lastElementPtr, sizePtr, count, "-")
	gen.AddSetHeap("(int)"+lastElementPtr, "0")
	gen.AddExpression(count, count, "1", "-")
	gen.AddSetHeap("(int)"+sizePtr, count)

	result = environment.NewValue(count, true, environment.INTEGER, false, false, false)
	return result
}
