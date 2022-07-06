package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal
	var b byte = 255
	fmt.Println("Literal inteiro é", reflect.TypeOf(b))

	// Com sinal
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("Literal inteiro é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode
	fmt.Println("O rune é", i2)
	fmt.Println(i2)

	//numeros reais
	var x float32 = 49.99
	fmt.Println("o tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("o tipo de bo é", reflect.TypeOf(bo))
	fmt.Println("o de bo é", bo)

}
