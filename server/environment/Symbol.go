package environment

type Symbol struct {
	Lin          int
	Col          int
	Id           string
	Tipo         TipoExpresion
	Valor        interface{}
	Mutable      bool
	Externo      string
	Inout        bool
	ReturnFlag   bool
	BreakFlag    bool
	ContinueFlag bool
	Posicion     int
}
