package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Errorf("Resultado da soma é invalido: Resultado %d, Esperado: %d", total, 30)
	}
}
