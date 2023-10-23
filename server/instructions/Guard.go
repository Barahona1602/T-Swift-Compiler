package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"fmt"
	"strings"
)

type Guard struct {
	Lin       int
	Col       int
	Condition interfaces.Expression
	Bloque    []interface{}
}

func NewGuard(lin int, col int, condition interfaces.Expression, bloque []interface{}) Guard {
	Inst := Guard{lin, col, condition, bloque}
	return Inst
}

func (p Guard) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generando Guard")
	var condicion, result environment.Value
	RetLvl := gen.NewLabel() // Etiqueta de retorno
	gen.AddLabel(RetLvl)

	// Etiquetas para controlar break y continue
	ContinueLabel := gen.NewLabel()
	BreakLabel := gen.NewLabel()
	gen.AddContinue(ContinueLabel)
	gen.AddBreak(BreakLabel)

	condicion = p.Condition.Ejecutar(ast, env, gen)

	// Agregar etiquetas TrueLabel como sea necesario
	for _, lvl := range condicion.TrueLabel {
		gen.AddLabel(lvl.(string))
	}

	// Instrucciones del bucle guard
	for _, s := range p.Bloque {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				if value, ok := resInst.(environment.Value); ok {
					// Manejar Break y Continue
					if value.BreakFlag {
						gen.AddGoto(BreakLabel)
					}
					if value.ContinueFlag {
						gen.AddGoto(ContinueLabel)
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
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				if value, ok := resInst.(environment.Value); ok {
					// Manejar Break y Continue
					if value.BreakFlag {
						gen.AddGoto(BreakLabel)
					}
					if value.ContinueFlag {
						gen.AddGoto(ContinueLabel)
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

	// Retorno
	gen.AddGoto(RetLvl)

	// Agregar false labels
	for _, lvl := range condicion.FalseLabel {
		gen.AddLabel(lvl.(string))
	}

	return result
}
