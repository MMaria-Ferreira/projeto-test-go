package main

import ()

const frances = "frances"
const espanhol = "espanhol"
const prefixoOlaFrances = "Bonjour"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return prefixoSausacao(idioma) + nome
}
func prefixoSausacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues

	}
	return
}
