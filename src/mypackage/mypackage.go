package mypackage

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimir mensaje
func PrintMessage(text string) {
	println(text)
}

func printMessage1(text string) {
	println(text)
}
