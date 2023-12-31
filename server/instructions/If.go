package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"fmt"
	"strings"
)

type If struct {
	Lin        int
	Col        int
	Expresion  interfaces.Expression
	Bloque     []interface{}
	ElseBloque []interface{}
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque []interface{}, elseBloque []interface{}) If {
	ifInstr := If{lin, col, condition, bloque, elseBloque}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("****** Generando If")
	var condicion, result environment.Value
	var OutLvls []interface{}
	condicion = p.Expresion.Ejecutar(ast, env, gen)
	newLabel := gen.NewLabel()
	//*****************************************add true labels
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}
	//instrucciones if
	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				//agregando etiquetas de salida
				for _, lvl := range resInst.(environment.Value).OutLabel {
					OutLvls = append(OutLvls, lvl)
				}
			}
		} else {
			ast.SetError("Error en bloque", p.Lin, p.Col)
		}
	}
	//*****************************************add out labels
	gen.AddGoto(newLabel)
	//*****************************************add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}
	OutLvls = append(OutLvls, newLabel)

	// Ejecutar el bloque "else" si existe
	if len(p.ElseBloque) > 0 {
		gen.AddComment("****** Generando Else")
		for _, s := range p.ElseBloque {
			if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
				resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
				if resInst != nil {
					//agregando etiquetas de salida
					for _, lvl := range resInst.(environment.Value).OutLabel {
						OutLvls = append(OutLvls, lvl)
					}
				}
			} else {
				ast.SetError("Error en bloque else", p.Lin, p.Col)
			}
		}
	}

	copiedSlice := make([]interface{}, len(OutLvls))
	for i, item := range OutLvls {
		copiedSlice[i] = item
	}

	result.OutLabel = copiedSlice
	return result
}
