package environment

type StructType struct {
	Id       string
	Tipo     TipoExpresion
	StructId string
}

func NewStructType(id string, tip TipoExpresion, ids string) StructType {
	exp := StructType{id, tip, ids}
	return exp
}
