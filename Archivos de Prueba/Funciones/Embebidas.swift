print("--------------------------")
print("---FUNCIONES EMBEBIDAS----")
print("----------14 pts----------")
print("--------------------------")

print("")
func suma(_ numero1: Int, _ numero2: Int) -> Int {
    let resultado = numero1 + numero2
    print("La suma es: ")
    print(resultado)
}

suma(3,5)

print("")
func saludo3() {
    print("saludos!")
}

func saludo2() {
    print("mundo")
    saludo3()
}

func saludo1() {
    print("hola")
    saludo2()
}

saludo1()

print("")
func ejemplo2(verdura v: Int, v verdura: Int ) {
    print(v)
    print(verdura)
}

let precio1: Int = 251
let precio2: Int = 85
ejemplo2(precio1,precio2)



/*
--------------------------
---FUNCIONES EMBEBIDAS----
----------14 pts----------
--------------------------

La suma es:  8

hola
mundo
saludos!

251
85
*/