package main

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObterPontuacaoJogador(nome string) int
}
type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		s.registrarVitoria(w)
	case http.MethodGet:
		s.mostrarPontuacao(w, r)
	}
}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	pontuacao := s.armazenamento.ObterPontuacaoJogador(jogador)

	if pontuacao == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuacao)
}

func (s *ServidorJogador) registrarVitoria(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

