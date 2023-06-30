package dicionario

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"
	
		compareStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	
	})
	
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}
 func compareStrings(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona("Intrínseco", "Que se encontra na essência ou na natureza de algo ou alguém")
	
	esperado := "Que se encontra na essência ou na natureza de algo ou alguém"
	resultado, err := dicionario.Busca("Intrínseco")

	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada: ", err)
	}

	if esperado != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
