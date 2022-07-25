package multiplexar

func encaminhar(origem <-chan string, destino chan string) chan string {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	result := make(chan string)
	go encaminhar(entrada1, result)
	go encaminhar(entrada2, result)
	return result
}

func main() {
	// canalUnico := juntar(
	// 	Titulo("https://www.cod3r.com.br", "https://www.google.com"),
	// 	Titulo("https://www.amazon.com", "https://www.youtube.com"),
	// )

	// fmt.Println("Primeiros:", <-canalUnico, "|", <-canalUnico)
	// fmt.Println("Segundos:", <-canalUnico, "|", <-canalUnico)
}
