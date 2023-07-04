package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperSpy struct {
	Chamadas int
}

type SleeperPadrao struct{}

const ultimaPlavra = "Vai!"
const inicioContagem = 3

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}
func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(saida, i)
	}
	sleeper.Sleep()
	fmt.Fprint(saida, ultimaPlavra)

}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
