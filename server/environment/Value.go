package environment

type Value struct {
	Value        string
	IsTemp       bool
	Type         TipoExpresion
	ArrType      TipoExpresion
	TrueLabel    []interface{}
	FalseLabel   []interface{}
	OutLabel     []interface{}
	ReturnFlag   bool
	BreakFlag    bool
	ContinueFlag bool
	IntValue     int
}

func NewValue(Val string, tmp bool, typ TipoExpresion, ret, br, cont bool) Value {
	result := Value{
		Value:        Val,
		IsTemp:       tmp,
		Type:         typ,
		IntValue:     0,
		ReturnFlag:   ret,
		BreakFlag:    br,
		ContinueFlag: cont,
	}
	return result
}
