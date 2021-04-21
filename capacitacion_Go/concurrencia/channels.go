package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	canal := make(chan int, 10)
	salida := make(chan struct{})
	go publicar(canal, salida, wg)
	go escuchar(canal, salida, wg)
	wg.Wait()
}

func publicar(canal chan<- int, salida chan struct{}, wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Publicando...")
	canal <- 10
	canal <- 20
	canal <- 30
	canal <- 40
	canal <- 50
	canal <- 60
	canal <- 70
	canal <- 80
	canal <- 90
	canal <- 100
	fmt.Println("Saliendo...")
	time.Sleep(10 * time.Second)
	salida <- struct{}{}
	wg.Done()
}
func escuchar(canal <-chan int, salida chan struct{}, wg *sync.WaitGroup) {
	fmt.Println("Escuchando...")
	ind := 0
	for ind == 0 {
		select {
		case dato := <-canal:
			fmt.Println("Recibido: ", dato)
		case <-salida:
			fmt.Println("Salida...")
			ind = 1
		}
	}
	wg.Done()
}
