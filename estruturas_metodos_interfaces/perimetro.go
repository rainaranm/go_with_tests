package perimetro

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

type Triangulo struct {
	Base float64
	Altura float64
}

type Forma interface {
	Area() float64
}

func Perimetro(retangulo Retangulo) (resultado float64) {

	resultado = (retangulo.Largura + retangulo.Altura) * 2

	return
}

func (r Retangulo) Area() (resultado float64) {
	resultado = r.Largura * r.Altura
	return
}

func (c Circulo) Area() (resultado float64) {
	resultado = math.Pi * c.Raio * c.Raio
	return
}

func (t Triangulo) Area() float64 {
    return (t.Base * t.Altura) * 0.5
}
