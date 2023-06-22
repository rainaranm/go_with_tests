package main 

func Soma(numeros []int) (resultado int) {
	
	for _, numero := range numeros {
		resultado += numero
	}
	
	return
}