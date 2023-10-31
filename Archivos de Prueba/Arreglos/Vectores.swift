print("--------------------------")
print("---------VECTORES---------")
print("----------12 pts----------")
print("--------------------------")

print("")
print("=============================================")
print("================CREACIÓN=====================")
print("=============================================")
var arr1: [Int] = [8, 4, 6, 2]
var arr2: [Int] = [40, 21, 1, 3, 14, 4]
var arr3: [Int] = [90, 3, 40, 10, 8, 5]
print("Se crean los arreglos arr1, arr2, arr3")

print("")
print("=============================================")
print("=================ACCESO======================")
print("=============================================")

print("arr1: ")
print(arr1[1]+4)
print("arr2: ")
print(5+8*5-arr2[2])
print("arr3: ")
print(arr3[4]*8)

print("=============================================")
print("================FUNCIONES====================")
print("=============================================")

// arr1.append(9)
// print("append arr1: ", arr1[4])

// arr2.removeLast()
// arr2.append(6)
// print("append arr2: ", arr2[5])

// arr3.remove(at: 4)
// print("append arr3: ", arr3[4])

print("arr1 vacío: ")
print(arr1.isEmpty)
print("arr1 cantidad: ")
print(arr1.count)

/*
--------------------------
---------VECTORES---------
----------12 pts----------
--------------------------

=============================================
================CREACIÓN=====================
=============================================
Se crean los arreglos arr1, arr2, arr3

=============================================
=================ACCESO======================
=============================================
arr1:  8
arr2:  44
arr3:  64
=============================================
================FUNCIONES====================
=============================================
append arr1:  9
append arr2:  6
append arr3:  5
arr1 vacío:  false
arr1 cantidad:  5
*/