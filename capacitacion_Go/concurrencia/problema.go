package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

type Estructura struct {
	mapa  map[string]int
	mutex sync.Mutex
}

var estructura = Estructura{mapa: make(map[string]int)}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go guardar(fmt.Sprintf("obj_%d", i), r.Intn(100), wg)
	}
	go leerDato(fmt.Sprintf("obj_%d", r.Intn(100)))
	wg.Wait()
	fmt.Println(estructura.mapa)
}

func guardar(llave string, valor int, wg *sync.WaitGroup) {
	estructura.mutex.Lock()
	estructura.mapa[llave] = valor
	estructura.mutex.Unlock()
	wg.Done()
}

func leerDato(clave string) {
	estructura.mutex.Lock()
	fmt.Printf(" Leyendo mapa en [%s] --> %d \n", clave, estructura.mapa[clave])
	estructura.mutex.Unlock()
}
