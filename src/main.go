package main

import (
	"fmt"
	"strings"
)

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

func esPar(a int) string {
	if a%2 == 0 {
		return "Es par"
	} else {
		return "No es par"
	}
}

func validarUsuario(u string, p string) string {
	if u == "raiser93" && p == "1234" {
		return "Usuario valido"
	} else {
		return "Credenciales incorrectas"
	}
}

func isPalindromo(texto string) {
	var textReverse string

	texto = strings.ToLower(texto)

	for i := len(texto) - 1; i >= 0; i-- {
		textReverse += string(texto[i])
	}

	if texto == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
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

	// normalFunction("Hola Mundo")
	// normalFunction("Hola Mundo 2")
	// normalFunction("Hola Mundo 3")

	// tripleArgument(1, 2, "Hola")

	// value := returnValue(2)
	// fmt.Println("Value:", value)

	// value1, value2 := doubleReturn(2)
	// println("Value 1 y value 2:", value1, value2)

	// value1, _ = doubleReturn(3)
	// println("Value 1 y value 2:", value1)

	// // For condicional
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Printf("\n")

	// // For while
	// counter := 0
	// for counter < 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }
	// fmt.Printf("\n")
	// // For forever
	// // counterForever := 0
	// // for {
	// // 	fmt.Println(counterForever)
	// // 	counterForever++
	// // }

	// // For range
	// listaNumeroPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	// for i, par := range listaNumeroPares {
	// 	fmt.Printf("posicion %d numero par %d \n", i, par)
	// }

	// fmt.Printf("\n")

	// // Reto
	// for i := 10; i > 0; i-- {
	// 	fmt.Println(i)
	// }

	// fmt.Printf("\n")

	// counter = 10
	// for counter > 0 {
	// 	fmt.Println(counter)
	// 	counter--
	// }

	// valor1 := 1
	// valor2 := 2

	// if valor1 == 1 {
	// 	fmt.Println("Es 1")
	// } else {
	// 	fmt.Println("No es 1")
	// }

	// // With and

	// if valor1 == 1 && valor2 == 2 {
	// 	fmt.Println("Es verdad")
	// }

	// if valor1 == 0 || valor2 == 2 {
	// 	fmt.Println("Es verdad, OR")
	// }

	// fmt.Println(esPar(6))
	// fmt.Println(esPar(7))

	// fmt.Println(validarUsuario("raiser93", "1234"))
	// fmt.Println(validarUsuario("raiser93", "12345"))

	// switch modulo := 6 % 2; modulo {
	// case 0:
	// 	fmt.Println("Es par")
	// 	break
	// default:
	// 	fmt.Println("Es impar")
	// }

	// // sin codicion
	// value := 101
	// switch {
	// case value > 100:
	// 	fmt.Println("Es mayor a 100")
	// case value < 0:
	// 	fmt.Println("Es menor a 0")
	// default:
	// 	fmt.Println("No condicion")
	// }

	// // Defer
	// defer fmt.Println("Hola")
	// fmt.Println("Mundo")

	// // Continue y  break
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// 	if i == 2 {
	// 		fmt.Println("Es 2")
	// 		continue
	// 	}

	// 	if i == 8 {
	// 		fmt.Println("Es 8")
	// 		break
	// 	}
	// }

	// Array
	// var array [4]int
	// array[0] = 1
	// array[1] = 2
	// fmt.Println(array, len(array), cap(array))

	// // slice
	// slice := []int{0, 1, 2, 3, 4, 5, 6}
	// fmt.Println(slice, len(slice), cap(slice))

	// // Metodos en el slice
	// fmt.Println(slice[0])
	// fmt.Println(slice[:3])
	// fmt.Println(slice[2:4])
	// fmt.Println(slice[4:])

	// // append
	// slice = append(slice, 7)
	// fmt.Println(slice)

	// // append nueva list
	// newSlice := []int{8, 9, 10}
	// slice = append(slice, newSlice...)
	// fmt.Println(slice)

	slice := []string{"Hola", "que", "hace?"}
	for i := range slice {
		fmt.Println(i)
	}

	isPalindromo("Ama")
	isPalindromo("amor a roma")
	isPalindromo("casas")

	var palabra string
	fmt.Scan(&palabra)
	isPalindromo(palabra)
}
