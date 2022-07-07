package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[1213131313131] = "Maria"
	aprovados[2321324353535] = "Pedro"
	aprovados[4969284840584] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[1213131313131])

	delete(aprovados, 1213131313131)
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

}
