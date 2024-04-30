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
	prefixo := prefixoOlaPortugues

	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol

	}
	return prefixo + nome

}
