package environment

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AST struct {
	Instructions []interface{}
	Print        string
	Errors       string
}

func NewAST(inst []interface{}, print string, err string) AST {
	ast := AST{Instructions: inst, Print: print, Errors: err}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

func (a *AST) GetErrors() string {
	return a.Errors
}

func (a *AST) SetError(ToErr string, lin int, col int) {
	errMsg := fmt.Sprintf("Error en línea %d, columna %d: %s", lin, col, ToErr)
	a.Errors = a.Errors + errMsg + "\n"
}

func (ast AST) PrintErrorsToFile() (string, error) {
	htmlContent := generateHTMLErrorTable(ast)
	fileName := "ErrorTable.html"

	file, err := os.Create(fileName)
	if err != nil {
		return htmlContent, err
	}
	defer file.Close()

	_, err = file.WriteString(htmlContent)
	if err != nil {
		return htmlContent, err
	}

	return htmlContent, nil
}

func generateHTMLErrorTable(ast AST) string {
	html := `
	<style>
		body {
			font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen',
			'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue',
			sans-serif;
			-webkit-font-smoothing: antialiased;
			-moz-osx-font-smoothing: grayscale;
			background-color: #282c34; /* Color de fondo personalizado */
			color: white; /* Color de texto blanco */
			font-size: calc(8px + 1vmin); /* Tamaño de fuente reducido */
		}
		h1 {
			text-align: center;
			margin-top: 20px;
		}
		table {
			width: 600px; /* Hacer la tabla más ancha */
			margin: auto;
			border-collapse: collapse;
			border: 1px solid white; /* Borde blanco */
			margin-top: 20px;
			font-size: calc(8px + 1vmin); /* Tamaño de fuente reducido */
		}
		th, td {
			padding: 12px; /* Mayor espacio alrededor de los elementos */
			text-align: left;
			font-size: calc(9px + 1vmin); /* Tamaño de fuente reducido */
		}
		tr:nth-child(even) {
			background-color: transparent; /* Celda transparente */
		}
		th {
			background-color: #FF5733; /* Color naranja */
			color: white;
		}
	</style>
		<body>
			<table>
				<tr>
					<th>No.</th>
					<th>Descripción</th>
					<th>Línea</th>
					<th>Columna</th>
				</tr>
	`

	errors := strings.Split(ast.Errors, "\n")
	for idx, errMsg := range errors {
		if errMsg != "" {
			parts := strings.SplitN(errMsg, ": ", 3)
			if len(parts) >= 2 {
				description := parts[1]
				positionParts := strings.SplitN(parts[0], ", ", 2)
				if len(positionParts) == 2 {
					lineStr := strings.TrimPrefix(positionParts[0], "Error en línea ")
					columnStr := strings.TrimPrefix(positionParts[1], "columna ")
					line, _ := strconv.Atoi(lineStr)
					column, _ := strconv.Atoi(columnStr)
					html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%d</td><td>%d</td></tr>", idx+1, description, line, column)
				}
			}
		}
	}

	html += `
			</table>
		</body>
	`

	return html
}
