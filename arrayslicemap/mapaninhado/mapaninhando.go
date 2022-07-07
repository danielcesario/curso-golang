package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.32,
			"Guga Pereira":   9583.45,
		},
		"J": {
			"José da Silva": 4424.54,
		},
		"M": {
			"Maria dos Santos": 4246.64,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
