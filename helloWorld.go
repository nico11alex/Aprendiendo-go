package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello, Go!")

	// variables
	var string1 string
	string1 = "Esto es una cadena"
	fmt.Println(string1)

	string1 = "cambie"
	fmt.Println(string1)

	var string2 = "valor2"
	fmt.Println(string2)

	var myInt int = 7
	myInt = myInt + 4

	fmt.Println(myInt)
	fmt.Println(myInt - 1)

	fmt.Println(string1, myInt)

	fmt.Println(reflect.TypeOf(myInt))

	var myFloat = 6.5
	fmt.Println(myFloat)

	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(myFloat + float64(myInt))

	var myBool bool = false
	myBool = true

	fmt.Println(myBool)

	string3 := "Esto es una cadena 3"
	fmt.Println(string3)

	// costantes

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	// control de flujo

	if myInt == 10 {
		fmt.Println("El valor es 10")
	} else if myInt == 11 {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}

	//  Array

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray[1])

	// Map

	myMap := make(map[string]int)
	myMap["Nicollas"] = 18

	// Lista
	myLista := list.New()
	myLista.PushBack(1)
	myLista.PushBack(2)
	myLista.PushBack(3)

	// Bucle

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray {
		fmt.Println(index, value)
	}

	fmt.Println(myFuction())

	// Estructura

	type MyStructure struct {
		name string
		edad int
	}

	myStructure := MyStructure{"Nicolas", 18}
	fmt.Println(myStructure)

}

func myFuction() string {
	return "My función"
}
