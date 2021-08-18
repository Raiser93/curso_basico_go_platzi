package main

import "fmt"

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
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado es:", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicación", result)

	// División
	result = y / x
	fmt.Println("División:", result)

	// Mudulo/residuo
	result = y % x
	fmt.Println("Mudulo:", result)

	// Incrementar
	x++
	fmt.Println("Incrementar:", x)

	// Decrementar
	x--
	fmt.Println("Decrementar:", x)

	// Reto
	// Area rectangulo
	const largo = 6
	const ancho = 2
	areaRectangulo := largo * ancho
	fmt.Println("Area del rectangulo:", areaRectangulo)

	// Area trapecio
	const baseMayor = 9.5
	const baseMenor = 3.5
	const altura = 4
	areaTrapecio := ((baseMayor + baseMenor) / 2) * altura
	fmt.Println("Area de trapecio:", areaTrapecio)

	// Area circulo
	const radio = 3
	const pi = 3.14
	areaCirculo := pi * (radio * radio)
	fmt.Println("Area de circulo:", areaCirculo)
}
