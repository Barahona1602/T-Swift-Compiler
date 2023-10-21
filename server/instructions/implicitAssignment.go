package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"strconv"
)

type ImplicitAssignment struct {
	Lin       int
	Col       int
	Id        string
	Operator  string
	Expresion interfaces.Expression
}

func NewImplicitAssignment(lin int, col int, id string, operator string, val interfaces.Expression) ImplicitAssignment {
	instr := ImplicitAssignment{lin, col, id, operator, val}
	return instr
}

func (p ImplicitAssignment) Ejecutar(ast *environment.AST, env interface{}, g *generator.Generator) interface{} {
	var result environment.Value
	g.AddComment("Generando asignación implícita")
	// Buscar la variable en el entorno
	variable := env.(environment.Environment).GetVariable(p.Id)
	// Ejecutar la expresión
	result = p.Expresion.Ejecutar(ast, env, g)

	// Obtener el valor actual de la variable
	varName := g.NewTemp() // Nombre temporal para almacenar el valor
	g.AddGetStack(varName, strconv.Itoa(variable.Posicion))

	// Realizar la operación dependiendo del operador
	switch p.Operator {
	case "+=":
		g.AddExpression(varName, varName, result.Value, "+")
	case "-=":
		g.AddExpression(varName, varName, result.Value, "-")
	default:
		return nil // Operador no soportado
	}

	// Asignar el resultado a la variable
	g.AddSetStack(strconv.Itoa(variable.Posicion), varName)
	g.AddBr()
	return result
}
