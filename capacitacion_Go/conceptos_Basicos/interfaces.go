package main

import "fmt"

type Animal interface {
	Mover() string
}

type Perro struct {
	nombre string
	patas string
}
type Serpiente struct {
	nombre string
}

func (p Perro) Mover() string{
	return "Caminando... Moviendo " + p.patas + " patas " 
}
func (s Serpiente) Mover() string{
	return "Reptando..."
}

func iniciar(a Animal) {
	fmt.Println(a.Mover())
}

func main(){
  perro := Perro{"sfsdfd", "4"}
  serp := Serpiente{"dfadf una"}
  iniciar(perro)
  iniciar(serp)
}