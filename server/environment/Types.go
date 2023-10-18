package environment

type TipoExpresion int

const (
	INTEGER TipoExpresion = iota //0
	FLOAT                        //1
	STRING                       //2
	BOOLEAN                      //3
	NIL                          //4
	ARRAY                        //5
	RANGE                        //6
	STR                          //7
	UNKNOWN                      //8
	STRUCT
)
