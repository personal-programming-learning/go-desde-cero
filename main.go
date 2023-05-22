package main

import (
	"fmt"
	"Curso-de-go-desde-cero/go-desdecero/ejercicios"
)

func main() {
	num, valorString := ejercicios.ConvertirStrinaEntero("15")
	fmt.Println(num)
	fmt.Println(valorString)
}