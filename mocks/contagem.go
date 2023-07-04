package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Pausa()
}

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa func(time.Duration)
}

func (c *SleeperConfiguravel) Pausa() {
	c.pausa(c.duracao)
}


type TempoSpy struct {
	duracaoPausa time.Duration
}


func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.duracaoPausa = duracao
}



const ultimaPlavra = "Vai!"
const inicioContagem = 3
const escrita = "escrita"
const pausa = "pausa"


func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	for i := inicioContagem; i > 0; i -- {
		fmt.Fprintln(saida, i)
	}
	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPlavra)

}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
