package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i

	fmt.Printf("endereço do ponteiro: %v - valor do ponteiro: %v\n", p, *p)
	i = 2
	fmt.Printf("endereço do ponteiro: %v - valor do ponteiro: %v", p, *p)

}
