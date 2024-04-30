package main

import ()

const frances = "frances"
const espanhol = "espanhol"
const prefixoOlaFrances = "Bonjour"
const prefixoOlaPortuges = "Olá, "
const prefixoOlaEspanhol = "Hola, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}
	if idioma == "Espanhol" {
		return prefixoOlaEspanhol + nome
	}
	if idioma == frances {
		return prefixoOlaFrances + nome

	}
	return prefixoOlaEspanhol + nome
}
