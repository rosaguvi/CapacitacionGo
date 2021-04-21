package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Inversionista struct {
	identificacion int
	nombre         string
	inversiones    []int
}

type Control struct {
	inv   map[int]Inversionista
	mutex sync.Mutex
}

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)
var control = Control{inv: make(map[int]Inversionista)}

func main() {
	cargarDatos()
	fmt.Println(control)
	wg := &sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go invertir(i, wg)
	}
	wg.Wait()
	fmt.Println(control)
}

func invertir(inv int, wg *sync.WaitGroup) {
	wg1 := &sync.WaitGroup{}
	wg1.Add(5)
	for i := 0; i < 5; i++ {
		canal := make(chan int, 10)
		go genInversio(control.inv[inv].inversiones[i], canal, wg1)
		dato := <-canal
		control.mutex.Lock()
		control.inv[inv].inversiones[i] = dato
		control.mutex.Unlock()
	}
	wg1.Wait()
	wg.Done()
}

func genInversio(vlr int, canal chan<- int, wg1 *sync.WaitGroup) {
	ir := IntRange{-10, 10}
	//fmt.Println("valor inicial: ", vlr)
	for i := 0; i < 100; i++ {
		valor := ir.NextRandom(r)
		var vlr_gan float32
		vlr_gan = float32(valor) / 20
		vlr_gan = vlr_gan * float32(vlr)
		vlr = vlr + int(vlr_gan)
		if vlr <= 0 {
			i = 100
			vlr = 0
		}
		i++
	}
	canal <- vlr
	//fmt.Println("valor Final: ", vlr)
	//time.Sleep(3 * time.Second)
	wg1.Done()
}

func cargarDatos() {
	arreglo := []int{30000, 25000, 18000, 30000, 20000}
	inversio := Inversionista{identificacion: 12345, nombre: "Juan", inversiones: arreglo}
	control.inv[0] = inversio
	arreglo = []int{40000, 28000, 22000, 20000, 26000}
	inversio = Inversionista{identificacion: 45678, nombre: "Tony", inversiones: arreglo}
	control.inv[1] = inversio
	arreglo = []int{35000, 15000, 17000, 26000, 30000}
	inversio = Inversionista{identificacion: 34567, nombre: "Ana", inversiones: arreglo}
	control.inv[2] = inversio
	arreglo = []int{21000, 35000, 20000, 18000, 28000}
	inversio = Inversionista{identificacion: 56789, nombre: "Lina", inversiones: arreglo}
	control.inv[3] = inversio
	arreglo = []int{31000, 23000, 16000, 25000, 22000}
	inversio = Inversionista{identificacion: 23456, nombre: "Ricardo", inversiones: arreglo}
	control.inv[4] = inversio
}
