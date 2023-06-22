package main

import "fmt"

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFracês = "Bonjour, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
}
func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
		case "espanhol":
			prefixo = prefixoOlaEspanhol
		case "francês":
			prefixo = prefixoOlaFracês
		default:
			prefixo = prefixoOlaPortugues
	}

	return
}

func main() {
	fmt.Println(Ola("Rainara", ""))

}
