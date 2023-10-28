package expressions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
	"strconv"
)

type Operation struct {
	Lin      int
	Col      int
	Op_izq   interfaces.Expression
	Operador string
	Op_der   interfaces.Expression
}

func NewOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) Operation {
	exp := Operation{Lin: lin, Col: col, Op_izq: Op1, Operador: Operador, Op_der: Op2}
	return exp
}

func (o Operation) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) environment.Value {
	var dominante environment.TipoExpresion

	tabla_dominante := [5][5]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				BOOLEAN				NIL
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NIL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NIL, environment.NIL},
		//STRING
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NIL},
		//BOOLEAN
		{environment.BOOLEAN, environment.NIL, environment.STRING, environment.BOOLEAN, environment.NIL},
		//NIL
		{environment.NIL, environment.NIL, environment.NIL, environment.NIL, environment.NIL},
	}

	var op1, op2, result environment.Value

	newTemp := gen.NewTemp()

	switch o.Operador {
	case "+":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if op1.Type == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.STRING {
				//llamar a generar concatstring
				gen.GenerateConcatString()
				//concat
				gen.AddComment("Concatenando strings")
				envSize := strconv.Itoa(env.(environment.Environment).Size["size"])
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("dbrust_concatString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				result = environment.NewValue(tmp2, true, dominante, false, false, false)
				return result
			} else {
				ast.SetError(" No es posible sumar", o.Lin, o.Col)
				return result
			}
		}

	case "-":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "-")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue - op2.IntValue
				return result
			} else {
				ast.SetError(" No es posible restar", o.Lin, o.Col)
			}
		}
	case "*":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "*")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue * op2.IntValue
				return result
			} else {
				ast.SetError(" No es posible Multiplicar", o.Lin, o.Col)
			}
		}
	case "/":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				lvl1 := gen.NewLabel()
				lvl2 := gen.NewLabel()

				gen.AddIf(op2.Value, "0", "!=", lvl1)
				gen.AddPrintf("c", "77")
				gen.AddPrintf("c", "97")
				gen.AddPrintf("c", "116")
				gen.AddPrintf("c", "104")
				gen.AddPrintf("c", "69")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "114")
				gen.AddPrintf("c", "111")
				gen.AddPrintf("c", "114")
				gen.AddExpression(newTemp, "0", "", "")
				gen.AddGoto(lvl2)
				gen.AddLabel(lvl1)
				gen.AddExpression(newTemp, op1.Value, op2.Value, "/")
				gen.AddLabel(lvl2)
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				return result
			} else {
				ast.SetError(" No es posible Dividir", o.Lin, o.Col)
			}

		}
	case "%":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]

			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, "fmod("+op1.Value, op2.Value+")", ",")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				return result
			} else {
				ast.SetError(" No es posible hacer modulo", o.Lin, o.Col)
			}
		}
	case "<":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {
				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result

			} else {
				ast.SetError(" No es posible comparar <", o.Lin, o.Col)
			}
		}
	case ">":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(" No es posible comparar >", o.Lin, o.Col)
			}
		}
	case "<=":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "<=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(" No es posible comparar <=", o.Lin, o.Col)
			}
		}
	case ">=":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, ">=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(" No es posible comparar >=", o.Lin, o.Col)
			}
		}
	case "==":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "==", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(" No es posible comparar == ", o.Lin, o.Col)
			}
		}
	case "!=":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if dominante == environment.INTEGER || dominante == environment.FLOAT {

				trueLabel := gen.NewLabel()
				falseLabel := gen.NewLabel()

				gen.AddIf(op1.Value, op2.Value, "!=", trueLabel)
				gen.AddGoto(falseLabel)

				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = append(result.TrueLabel, trueLabel)
				result.FalseLabel = append(result.FalseLabel, falseLabel)
				return result
			} else {
				ast.SetError(" No es posible comparar !=", o.Lin, o.Col)
			}
		}
	case "&&":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			//add op1 labels
			for _, lvl := range op1.TrueLabel {
				gen.AddLabel(lvl.(string))
			}

			op2 = o.Op_der.Ejecutar(ast, env, gen)

			result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op1.FalseLabel, result.FalseLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		}
	case "||":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)

			for _, lvl := range op1.FalseLabel {
				gen.AddLabel(lvl.(string))
			}
			op2 = o.Op_der.Ejecutar(ast, env, gen)

			result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)

			result.TrueLabel = append(op1.TrueLabel, result.TrueLabel...)
			result.TrueLabel = append(op2.TrueLabel, result.TrueLabel...)
			result.FalseLabel = append(op2.FalseLabel, result.FalseLabel...)
			return result
		}
	case "!":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			if op1.Type == environment.BOOLEAN {
				// Invertir las etiquetas de verdadero y falso
				result = environment.NewValue("", false, environment.BOOLEAN, false, false, false)
				result.TrueLabel = op1.FalseLabel
				result.FalseLabel = op1.TrueLabel
				return result
			} else {
				ast.SetError("No es posible realizar ! con el tipo de operando proporcionado", o.Lin, o.Col)
			}
		}

	case "NEGACION":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			if op1.Type == environment.INTEGER {
				gen.AddExpression(newTemp, "0", op1.Value, "-")
				result = environment.NewValue(newTemp, true, environment.INTEGER, false, false, false)
				return result
			} else if op1.Type == environment.FLOAT {
				gen.AddExpression(newTemp, "0", op1.Value, "-")
				result = environment.NewValue(newTemp, true, environment.FLOAT, false, false, false)
				return result
			} else {
				ast.SetError(" No es posible realizar la negación", o.Lin, o.Col)
			}
		}
	case ",":
		{
			op1 = o.Op_izq.Ejecutar(ast, env, gen)
			op2 = o.Op_der.Ejecutar(ast, env, gen)
			dominante = tabla_dominante[op1.Type][op2.Type]
			if op1.Type == environment.INTEGER || dominante == environment.FLOAT {
				gen.AddExpression(newTemp, op1.Value, op2.Value, "+")
				result = environment.NewValue(newTemp, true, dominante, false, false, false)
				result.IntValue = op1.IntValue + op2.IntValue
				return result
			} else if dominante == environment.STRING {
				//llamar a generar concatstring
				gen.GenerateConcatString()
				//concat
				gen.AddComment("Concatenando strings")
				envSize := strconv.Itoa(env.(environment.Environment).Size["size"])
				tmp1 := gen.NewTemp()
				tmp2 := gen.NewTemp()
				gen.AddExpression(tmp1, "P", envSize, "+")
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op1.Value)
				gen.AddExpression(tmp1, tmp1, "1", "+")
				gen.AddSetStack("(int)"+tmp1, op2.Value)
				gen.AddExpression("P", "P", envSize, "+")
				gen.AddCall("dbrust_concatString")
				gen.AddGetStack(tmp2, "(int)P")
				gen.AddExpression("P", "P", envSize, "-")
				gen.AddBr()
				result = environment.NewValue(tmp2, true, dominante, false, false, false)
				return result
			} else {
				ast.SetError(" No es posible concatenar", o.Lin, o.Col)
				return result
			}
		}
	}
	gen.AddBr()
	return environment.Value{}
}
