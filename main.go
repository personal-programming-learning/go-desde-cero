package main

import (
	"fmt"
	"Curso-de-go-desde-cero/go-desdecero/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(12)
	fmt.Println(estado)
	fmt.Println(texto)
}