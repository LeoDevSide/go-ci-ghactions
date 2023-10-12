package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	
	if total != 30 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 30)
	}
}

func TestSubtracao(t *testing.T) {
	total := Subtracao(15, 15)
	
	if total != 0 {
		t.Errorf("Resultado da subtração é inválido: Resultado %d. Esperado: %d", total, 0)
	}
}

func TestMultiplicacao(t *testing.T) {
	total := Multiplicacao(15, 15)
	
	if total != 225 {
		t.Errorf("Resultado da multiplicação é inválido: Resultado %d. Esperado: %d", total, 225)
	}
}

func TestDivisao(t *testing.T) {
	total := Divisao(15, 15)
	
	if total != 1 {
		t.Errorf("Resultado da divisão é inválido: Resultado %d. Esperado: %d", total, 1)
	}
}
