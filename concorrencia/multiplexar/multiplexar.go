package main

import (
	"fmt"

	"github.com/acmasters/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar juntar mensagens em um canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
''
	return c
}

func main() {
	c := juntar(
		html.Titulo("http://www.google.com", "http://www.cod3r.com.br"),
		html.Titulo("http://www.youtube.com", "http://www.uol.com.br"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)

}
