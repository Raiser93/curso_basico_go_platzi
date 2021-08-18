package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// Primera linea de GO
	// fmt.Println("Hola mundo")

	// Declarar constante
	// const pi float64 = 3.14
	// const pi2 = 3.1415
	// fmt.Println("pi:", pi)
	// fmt.Println("pi2:", pi2)

	// Declaracion de variables enteras
	// base := 12
	// var altura int = 14
	// var area int

	// fmt.Println(base, altura, area)

	// Zero values
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a, b, c, d)

	// Area cuadrado
	// const baseCuadrado = 10
	// areaCuadrado := baseCuadrado * baseCuadrado
	// fmt.Println("Area cuadrado es:", areaCuadrado)

	// x := 10
	// y := 50

	// Suma
	// result := x + y
	// fmt.Println("Suma:", result)

	// Resta
	// result = y - x
	// fmt.Println("Resta:", result)

	// Multiplicacion
	// result = x * y
	// fmt.Println("Multiplicación", result)

	// División
	// result = y / x
	// fmt.Println("División:", result)

	// Mudulo/residuo
	// result = y % x
	// fmt.Println("Mudulo:", result)

	// Incrementar
	// x++
	// fmt.Println("Incrementar:", x)

	// Decrementar
	// x--
	// fmt.Println("Decrementar:", x)

	// Reto
	// Area rectangulo
	// const largo = 6
	// const ancho = 2
	// areaRectangulo := largo * ancho
	// fmt.Println("Area del rectangulo:", areaRectangulo)

	// Area trapecio
	// const baseMayor = 9.5
	// const baseMenor = 3.5
	// const altura = 4
	// areaTrapecio := ((baseMayor + baseMenor) / 2) * altura
	// fmt.Println("Area de trapecio:", areaTrapecio)

	// Area circulo
	// const radio = 3
	// const pi = 3.14
	// areaCirculo := pi * (radio * radio)
	// fmt.Println("Area de circulo:", areaCirculo)

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	// Declaracion de variables
	// helloMessage := "Hello"
	// worldMessage := "World"

	// // Println
	// fmt.Println(helloMessage, worldMessage)
	// fmt.Println(helloMessage, worldMessage)

	// // Printf
	// nombre := "Platzi"
	// cursos := 500

	// // Por buenas practicas poner la letra inicial del tipo de dato si  se conoce
	// fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	// // Si no poner v para los datos que se desconosca el tipo de dato
	// fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	// // Sprintf
	// message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	// fmt.Println(message)

	// // Tipos de datos
	// fmt.Printf("helloMessage: %T\n", helloMessage)
	// fmt.Printf("Cursos: %T\n", cursos)

	normalFunction("Hola Mundo")
	normalFunction("Hola Mundo 2")
	normalFunction("Hola Mundo 3")

	tripleArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	println("Value 1 y value 2:", value1, value2)

	value1, _ = doubleReturn(3)
	println("Value 1 y value 2:", value1)
}
