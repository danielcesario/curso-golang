package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAntior, _ := fatorial(n - 1)
		return n * fatorialAntior, nil
	}
}

func fatorialSimplificado(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorialSimplificado(n-1)
	}
}

func main() {
	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}

	resultado2 := fatorialSimplificado(8)
	fmt.Println(resultado2)
}
