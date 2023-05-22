package ejercicios

import (
	"fmt"
	"strconv"
)

func ConvertirStrinaEntero(valor string) (int, string) {
	num, err := strconv.Atoi(valor)
	if err != nil {
		fmt.Println("Error al convertir el string a número entero:", err)
	}

	var valorString string
	if num > 100 {
		valorString = "Es mayor a 100"
	} else {
		valorString = "Es menor a 100"
	}

	return num, valorString
}