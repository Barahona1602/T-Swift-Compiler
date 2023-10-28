print("--------------------------")
print("---------MATRICES---------")
print("----------14 pts----------")
print("--------------------------")

print("")
print("Creación de Matriz!")
let notas: [[Int]] = [
    [53, 88, 95, 89, 75],
    [81, 51, 57, 67, 93],
    [94, 74, 58, 84, 100]
]

print("")
print("Acceso:")
print("Esta suma debería ser 100: ")
print(notas[1][0]+19)
print("Esta suma debería ser 100: ")
print(notas[0][4]+25)
print("Esta suma debería ser 100: ")
print(notas[2][0]+6)
print("Esta suma debería ser 245: ")
print(notas[0][1]+notas[1][2]+notas[2][4])

print("")
let md: [Int] = [1, 2, 3, 4, 5, 6]
print("Matriz de una dimensión: ")
print(md[0])

print("")
let mdd: [[Int]] = [
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6],
    [1, 2, 3, 4, 5, 6]
]
print("Matriz de dos dimensiones: ")
print(mdd[0][1])

print("")
let mddd: [[[Int]]] = [
    [
        [1, 2, 3, 4, 5, 6],
        [1, 2, 3, 4, 5, 6], 
        [1, 2, 3, 4, 5, 6]
    ],
    [
        [1, 2, 3, 4, 5, 6],
        [1, 2, 3, 4, 5, 6],
        [1, 2, 3, 4, 5, 6]
    ]
]
print("Matriz de tres dimensiones: ")
print(mddd[0][1][2])

print("")
let mdddd: [[[[Int]]]] = [
    [
        [
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6]
        ],
        [
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6]
        ]
    ],
    [
        [
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6]
        ],
        [
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6],
            [1, 2, 3, 4, 5, 6]
        ]
    ]
]
print("Matriz de N dimensiones: ")
print(mdddd[0][1][2][3])

/*
--------------------------
---------MATRICES---------
----------14 pts----------
--------------------------

Creación de Matriz!

Acceso:
Esta suma debería ser 100:  100
Esta suma debería ser 100:  100
Esta suma debería ser 100:  100
Esta suma debería ser 245:  245

Matriz de una dimensión:  1

Matriz de dos dimensiones:  2

Matriz de tres dimensiones:  3

Matriz de N dimensiones:  4
*/