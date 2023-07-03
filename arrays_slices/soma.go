package soma 

func Soma(numeros []int) (resultado int) {
	
	for _, numeroAtual := range numeros {
		resultado += numeroAtual
	}
	
	return
}

func SomaTudo(numeros ...[]int) (resultado []int){

	for _, numeroAtual := range numeros {
		resultado = append(resultado, Soma(numeroAtual))
	}

	return
}

func SomaTodoOResto(numeros ...[]int)(resultado []int) {
	for _, numeroAtual := range numeros {
		if len(numeroAtual) == 0 {
			resultado = append(resultado, 0)
		} else {
			final := numeroAtual[1:]
			resultado = append(resultado, Soma(final))
		}

	}
	return
}