package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"fmt"
	"strings"
)

type While struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    []interface{}
}

func NewWhile(lin int, col int, condition interfaces.Expression, bloque []interface{}) While {
	Inst := While{lin, col, condition, bloque}
	return Inst
}

func (p While) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generando While")
	var condicion, result environment.Value
	RetLvl := gen.NewLabel() // Etiqueta de retorno
	gen.AddLabel(RetLvl)

	condicion = p.Expresion.Ejecutar(ast, env, gen)

	gen.AddContinue(RetLvl)
	gen.AddBreak(condicion.FalseLabel[0].(string))
	// Agregar etiquetas TrueLabel como sea necesario
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}

	// Instrucciones del bucle while
	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				if value, ok := resInst.(environment.Value); ok {
					if value.BreakFlag {
						result.BreakFlag = false
					}
					if value.ContinueFlag {
						result.ContinueFlag = false
					}
					for _, lvl := range value.OutLabel {
						gen.AddLabel(lvl.(string))
					}
				} else {
					fmt.Println("resInst no se puede convertir a environment.Value")
				}
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				if value, ok := resInst.(environment.Value); ok {
					if value.BreakFlag {
						result.BreakFlag = false
					}
					if value.ContinueFlag {
						result.ContinueFlag = false
					}
					// Agregar etiquetas de salida si es necesario
					for _, lvl := range value.OutLabel {
						gen.AddLabel(lvl.(string))
					}
				} else {
					fmt.Println("resInst no se puede convertir a environment.Value")
				}
			} else {
				fmt.Println("resInst es nulo")
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}

	//retorno
	gen.AddGoto(RetLvl)
	//add false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}
	return result

}
