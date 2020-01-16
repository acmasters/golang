package main

import ( 
	"html"
)

func OmaisRapido(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	//estrutura de concorrencia 
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case -< time.After(1000 * time.Millisecond):
		return "Todos Perderam!!"
	default:
		return "Sem Resposta!"

		}
}
func main() {
	campeao := OmaisRapido(
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.terra.com.br",
	)
	fmt.Println(campeao)
}
