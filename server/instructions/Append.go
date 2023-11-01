package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
)

type Append struct {
	Lin   int
	Col   int
	Array interfaces.Expression
	Expr  interfaces.Expression
}

func NewAppend(lin int, col int, array interfaces.Expression, expr interfaces.Expression) Append {
	exp := Append{lin, col, array, expr}
	return exp
}

func (p Append) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	tempArray := p.Array.Ejecutar(ast, env, gen)
	sizePtr := tempArray.Value
	count := gen.NewTemp()
	gen.AddGetHeap(count, "(int)"+sizePtr)
	gen.AddExpression(count, count, "1", "+")
	gen.AddSetHeap("(int)"+sizePtr, count)
	newElement := p.Expr.Ejecutar(ast, env, gen)
	newElementPtr := gen.NewTemp()
	gen.AddExpression(newElementPtr, sizePtr, count, "+")
	gen.AddSetHeap("(int)"+newElementPtr, newElement.Value)
	result = environment.NewValue(count, true, environment.INTEGER, false, false, false)

	return result
}
