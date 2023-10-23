package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"fmt"
	"strconv"
	"strings"
)

type ForIn struct {
	Lin  int
	Col  int
	Id   string
	Exp  interfaces.Expression
	Inst []interface{}
}

func NewForIn(lin int, col int, id string, exp interfaces.Expression, ins []interface{}) ForIn {
	loopInstr := ForIn{lin, col, id, exp, ins}
	return loopInstr
}

func (p ForIn) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	gen.AddComment("Generando For")
	var result environment.Value
	var OutLvls []interface{}
	tmpArr := p.Exp.Ejecutar(ast, env, gen)
	//forin
	tmp1 := gen.NewTemp()
	tmp2 := gen.NewTemp()
	tmp3 := gen.NewTemp()
	tmp4 := gen.NewTemp()
	tmp5 := gen.NewTemp()
	tmp6 := gen.NewTemp()
	lvl1 := gen.NewLabel()
	lvl2 := gen.NewLabel()
	lvl3 := gen.NewLabel()
	gen.AddGetHeap(tmp1, "(int)"+tmpArr.Value)
	gen.AddExpression(tmp2, "0", "", "")
	gen.AddExpression(tmp3, tmp2, "1", "+")
	gen.AddExpression(tmp4, tmpArr.Value, tmp3, "+")
	gen.AddGetHeap(tmp5, "(int)"+tmp4)
	symSave := env.(environment.Environment).SaveVariable(p.Id, environment.ARRAY, tmpArr.ArrType)
	gen.AddExpression(tmp6, "P", strconv.Itoa(symSave.Posicion), "+")
	gen.AddSetStack("(int)"+tmp6, tmp5)
	gen.AddLabel(lvl1)
	gen.AddIf(tmp2, tmp1, "==", lvl3)
	for _, s := range p.Inst {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				if result, ok := resInst.(environment.Value); ok {
					OutLvls = append(OutLvls, result.OutLabel...)

				}
				fmt.Println("tipo de dato de resInst", resInst)
			}

			if result.BreakFlag {
				gen.AddGoto(gen.BreakLabel)
				result.BreakFlag = false
			}
			if result.ContinueFlag {
				gen.AddGoto(gen.ContinueLabel)
				result.ContinueFlag = false
			}
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			resInst := s.(interfaces.Instruction).Ejecutar(ast, env, gen)
			if resInst != nil {
				if result, ok := resInst.(environment.Value); ok {
					OutLvls = append(OutLvls, result.OutLabel...)
				}
			}
			if result.BreakFlag {
				gen.AddGoto(gen.BreakLabel)
				result.BreakFlag = false
			}
			if result.ContinueFlag {
				gen.AddGoto(gen.ContinueLabel)
				result.ContinueFlag = false
			}
			for _, lvl := range result.OutLabel {
				gen.AddLabel(lvl.(string))
			}
		} else {
			fmt.Println("Error en bloque")
		}
	}

	gen.AddLabel(lvl2)
	gen.AddExpression(tmp2, tmp2, "1", "+")
	gen.AddExpression(tmp3, tmp2, "1", "+")
	gen.AddExpression(tmp4, tmpArr.Value, tmp3, "+")
	gen.AddGetHeap(tmp5, "(int)"+tmp4)
	gen.AddSetStack("(int)"+tmp6, tmp5)
	gen.AddGoto(lvl1)
	gen.AddLabel(lvl3)

	return result
}
