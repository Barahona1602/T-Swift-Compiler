package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"strconv"
)

type Declaration struct {
	Lin       int
	Col       int
	Id        string
	Tipo      environment.TipoExpresion
	Expresion interfaces.Expression
	Mutable   bool
}

func NewDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, val interfaces.Expression, mut bool) Declaration {
	instr := Declaration{lin, col, id, tipo, val, mut}
	return instr
}

func (p Declaration) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	var newVar environment.Symbol
	result = p.Expresion.Ejecutar(ast, env, gen)
	gen.AddComment("Agregando una declaracion")
	if p.Expresion == nil {
		// Asignar el valor nil al símbolo en el entorno
		newVar = env.(environment.Environment).SaveVariable(p.Id, environment.NIL, result.ArrType)
	} else if p.Tipo == environment.ARRAY {
		newVar = env.(environment.Environment).SaveVariable(p.Id, result.Type, result.ArrType)
	} else if p.Tipo == environment.UNKNOWN {
		newVar = env.(environment.Environment).SaveVariable(p.Id, result.Type, result.ArrType)
	} else {
		newVar = env.(environment.Environment).SaveVariable(p.Id, p.Tipo, result.ArrType)
	}

	if result.Type == environment.BOOLEAN {
		//si no es temp (boolean)
		newLabel := gen.NewLabel()
		//add labels
		for _, lvl := range result.TrueLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), "1")
		gen.AddGoto(newLabel)
		//add labels
		for _, lvl := range result.FalseLabel {
			gen.AddLabel(lvl.(string))
		}
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), "0")
		gen.AddGoto(newLabel)
		gen.AddLabel(newLabel)
		gen.AddBr()
	} else {
		//si es temp (num,string,etc)
		gen.AddSetStack(strconv.Itoa(newVar.Posicion), result.Value)
		gen.AddBr()
	}

	return result
}

func (p Declaration) ArrayValidation(result environment.Symbol) bool {
	return true
}
