package main

import ()

const espanhol = "espanhol"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if idioma == "Espanhol" {
		return prefixoOlaEspanhol + nome

	}
	return prefixoOlaEspanhol + nome
}
