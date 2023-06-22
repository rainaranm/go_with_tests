package main

import (
	"fmt"
	"testing"
)

func TestOla(t *testing.T) {

	verificarMensagemCorreta := func(t testing.TB, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %q, esperado %q", resultado, esperado)
		} else{
			fmt.Println(resultado)
		}

	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Sthefany","")
		esperado := "Olá, Sthefany"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para string vazia", func(t *testing.T) {
		resultado := Ola("","")
		esperado := "Olá, Mundo"

		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Rainara", "espanhol")
		esperado := "Hola, Rainara"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em francês", func(t *testing.T) {
		resultado := Ola("Millena", "francês")
		esperado := "Bonjour, Millena"
		
		verificarMensagemCorreta(t, resultado, esperado)
	})

	
}
