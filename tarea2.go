package main

import (
	"fmt"
	"math"
)

type Figuras interface {
	Area() float64
	Perimetro() float64
	Nombre() string
}

type Circulo struct {
	Radio float64
}

type Rectangulo struct {
	Ancho float64
	Largo float64
}

type Triangulo struct {
	LadoA float64
	LadoB float64
	LadoC float64
}

/* circulo */
func (c Circulo) Nombre() string {
	return "Circulo"
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * (math.Pi * c.Radio)
}

/* triangulo */

func (t Triangulo) Nombre() string {
	return "Triangulo"
}

func (t Triangulo) Area() float64 {
	s := (t.LadoA + t.LadoB + t.LadoC) / 2
	return math.Sqrt(s * (s - t.LadoA) * (s - t.LadoB) * (s - t.LadoC))
}

func (t Triangulo) Perimetro() float64 {
	return (t.LadoA + t.LadoB + t.LadoC)
}

/* rectangulo */

func (r Rectangulo) Nombre() string {
	return "Rectangulo"
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Largo
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Largo)
}

func imprimir_figuras(f Figuras) {
	fmt.Printf("Nombre: %s\n", f.Nombre())
	fmt.Printf("Area: %2.2f\n", f.Area())
	fmt.Printf("Perimetro: %2.2f\n", f.Perimetro())
}

func main() {

	c := Circulo{Radio: 5}
	r := Rectangulo{Ancho: 5, Largo: 3}
	t := Triangulo{LadoA: 3, LadoB: 4, LadoC: 5}

	imprimir_figuras(c)
	imprimir_figuras(r)
	imprimir_figuras(t)
}
