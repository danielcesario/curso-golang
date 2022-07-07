package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José João":      11324.45,
		"Gabriela Silva": 122583.87,
		"Ana Silva":      8765.23,
	}

	funcESalarios["Rafael Filho"] = 1350.42

	delete(funcESalarios, "teste")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
