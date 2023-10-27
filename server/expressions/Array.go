package expressions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"strconv"
)

type Array struct {
	Lin     int
	Col     int
	ListExp []interface{}
}

func NewArray(lin int, col int, list []interface{}) Array {
	exp := Array{lin, col, list}
	return exp
}

func (p Array) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var result environment.Value
	var val interfaces.Expression
	var arrType environment.TipoExpresion
	size := len(p.ListExp)
	gen.AddComment("Generando array")
	newtmp1 := gen.NewTemp()
	newtmp2 := gen.NewTemp()
	gen.AddAssign(newtmp1, "H")
	gen.AddExpression(newtmp2, newtmp1, "1", "+")
	gen.AddSetHeap("(int)H", strconv.Itoa(size))
	gen.AddExpression("H", "H", strconv.Itoa(size+1), "+")
	for _, exp := range p.ListExp {
		val = exp.(interfaces.Expression)
		arrType = val.Ejecutar(ast, env, gen).Type
		gen.AddSetHeap("(int)"+newtmp2, val.Ejecutar(ast, env, gen).Value)
		gen.AddExpression(newtmp2, newtmp2, "1", "+")
	}
	result = environment.NewValue(newtmp1, true, environment.ARRAY, false, false, false)
	result.ArrType = arrType
	return result
}
