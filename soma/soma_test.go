package main

import (
	"fmt"
	"testing"
)

func TestSoma(t*testing.T) {
	numeros := []int{1,2,3,4,5}

	resultado := Soma(numeros)
	esperado := 15

	if esperado != resultado {
		t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numeros := []int{1,2,3,4}
		Soma(numeros)
	}
}


func ExampleSoma() {
	numeros := []int{1,8}
	resultadoSoma := Soma(numeros)
	fmt.Println(resultadoSoma)
	//Output: 9
}