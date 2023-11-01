package instructions

import (
	"Server2/environment"
	"Server2/generator"
	"Server2/interfaces"
)

type ArrayRemoveAt struct {
	Lin   int
	Col   int
	Array interfaces.Expression
	Index interfaces.Expression
}

func NewRemoveAt(lin int, col int, array interfaces.Expression, index interfaces.Expression) ArrayRemoveAt {
	exp := ArrayRemoveAt{lin, col, array, index}
	return exp
}

func (p ArrayRemoveAt) Ejecutar(ast *environment.AST, env interface{}, gen *generator.Generator) interface{} {
	var result environment.Value
	tempArray := p.Array.Ejecutar(ast, env, gen)
	index := p.Index.Ejecutar(ast, env, gen)

	// Asegúrate de verificar que el índice sea válido y que el arreglo no esté vacío antes de intentar eliminar
	sizePtr := tempArray.Value
	count := gen.NewTemp()
	gen.AddGetHeap(count, "(int)"+sizePtr)

	validIndexCheckLabel := gen.NewLabel()
	emptyCheckLabel := gen.NewLabel()
	endLabel := gen.NewLabel()

	// Verifica que el índice esté dentro de los límites del arreglo
	gen.AddIf(index.Value, ">=", "0", validIndexCheckLabel)
	gen.AddComment("Índice no válido para eliminar")
	gen.AddGoto(endLabel)

	gen.AddLabel(validIndexCheckLabel)

	gen.AddIf(index.Value, "<", count, emptyCheckLabel)
	gen.AddComment("Índice no válido para eliminar")
	gen.AddGoto(endLabel)

	// Lógica para eliminar el elemento en la posición 'index'
	// Aquí debes implementar la eliminación del elemento en la posición indicada

	gen.AddComment("Eliminar elemento en la posición 'index'")

	// Actualiza el tamaño del arreglo
	gen.AddExpression(count, count, "1", "-")
	gen.AddSetHeap("(int)"+sizePtr, count)

	// Agrega aquí la lógica específica para eliminar el elemento

	gen.AddGoto(endLabel)

	// Etiqueta para verificar si el arreglo está vacío
	gen.AddLabel(emptyCheckLabel)
	gen.AddComment("No se puede eliminar en un arreglo vacío")

	// Etiqueta de final
	gen.AddLabel(endLabel)

	// Agrega aquí la lógica para manejar el caso de un arreglo vacío y la eliminación del elemento

	result = environment.NewValue(count, true, environment.INTEGER, false, false, false)
	return result
}
