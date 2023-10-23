package environment

import (
	"fmt"
	"os"
)

type Environment struct {
	Anterior interface{}
	Tabla    map[string]Symbol
	Id       string
	Size     map[string]int
}

func NewEnvironment(ant interface{}, id string) Environment {
	env := Environment{
		Anterior: ant,
		Tabla:    make(map[string]Symbol),
		Id:       id,
		Size:     make(map[string]int),
	}
	env.Size["size"] = 0
	return env
}

func (env Environment) SaveVariable(id string, tipo TipoExpresion, tipoarr TipoExpresion) Symbol {
	if variable, ok := env.Tabla[id]; ok {
		fmt.Println("La variable "+id+" ya existe ", variable)
		return env.Tabla[id]
	}
	env.Tabla[id] = Symbol{Lin: 0, Col: 0, Id: id, Tipo: tipo, Posicion: env.Size["size"], TipoArr: tipoarr}
	env.Size["size"] = env.Size["size"] + 1
	return env.Tabla[id]
}

func (env Environment) GetVariable(id string) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NIL, Posicion: 0}
}

func (env Environment) SetVariable(id string, value Symbol) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			value.Mutable = variable.Mutable
			if value.Mutable == true {
				if tmpEnv.Tabla[id].Tipo == value.Tipo {
					tmpEnv.Tabla[id] = value
					return variable
				} else {
					fmt.Println("El tipo de dato es incorrecto")
					return Symbol{Lin: 0, Col: 0, Tipo: NIL, Valor: 0}
				}
			} else {
				fmt.Println("La declaracion ", id, " es let y no se puede modificar")
				return Symbol{Lin: 0, Col: 0, Tipo: NIL, Valor: 0}
			}
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Tipo: NIL, Valor: 0}
}

func (env Environment) PrintVariablesToFile() (string, error) {
	htmlContent := generateHTMLTable(env)
	fileName := "SymbolTable.html"

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

func GetTable(env Environment) map[string]Symbol {
	return env.Tabla
}

func generateHTMLTable(env Environment) string {
	html := `
        <style>
            body {
                // ... (estilos anteriores)
            }
            h1 {
                // ... (estilos anteriores)
            }
            table {
                width: 800px; /* Doble del ancho original */
                margin: auto;
                border-collapse: collapse;
                border: 1px solid white;
                margin-top: 20px;
                font-size: calc(8px + 1vmin);
            }
            th, td {
                padding: 12px;
                text-align: left;
                font-size: calc(9px + 1vmin);
            }
            tr:nth-child(even) {
                background-color: transparent;
            }
            th {
                background-color: #00BFFF;
                color: white;
            }
        </style>
        <body>
            <table>
                <tr>
                    <th>Variable</th>
                    <th>Tipo</th>
                    <th>Ámbito</th>
                    <th>Línea</th>
                    <th>Columna</th>
                </tr>
    `

	for id, symbol := range env.Tabla {
		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%d</td><td>%d</td></tr>", id, getTypeName(symbol.Tipo), env.Id, symbol.Lin, symbol.Col)
	}

	html += `
            </table>
        </body>
    `

	return html
}

func getTypeName(t TipoExpresion) string {
	switch t {
	case INTEGER:
		return "INTEGER"
	case FLOAT:
		return "FLOAT"
	case STRING:
		return "STRING"
	case BOOLEAN:
		return "BOOLEAN"
	case NIL:
		return "NIL"
	case ARRAY:
		return "ARRAY"
	case RANGE:
		return "RANGE"
	case STR:
		return "STR"
	default:
		return "UNKNOWN"
	}
}
