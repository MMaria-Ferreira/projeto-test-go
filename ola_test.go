package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Mari")
	esperado := "Olá, Mari"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
