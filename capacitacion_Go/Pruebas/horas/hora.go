package main

import (
	"errors"
	"time"
)

func main() {
	//calculaHora(1)
}

func calculaHora(cntHoras int) (nuevaHora int, err error) {
	if cntHoras < 0 {
		err = errors.New("Error, no se sumaran horas negativas")
	} else {
		horaActual := time.Now()
		nuevaFecha := horaActual.Add(time.Hour * time.Duration(cntHoras))
		nuevaHora = nuevaFecha.Hour()
	}
	return
}
