package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Persona struct {
	nombre map[int]string
	mutex  sync.Mutex
}

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)
var persona = Persona{nombre: make(map[int]string)}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 200; i++ {
		opc := r.Intn(4)
		fmt.Println("iteracion ", i, "opcion: ", opc)
		llave := r.Intn(20)
		switch opc {
		case 0:
			go crear(llave, fmt.Sprintf("Persona_%d", r.Intn(100)), wg)
		case 1:
			go listar(llave, wg)
		case 2:
			go actualizar(llave, fmt.Sprintf("Persona_%d", r.Intn(100)), wg)
		case 3:
			go eliminar(llave, wg)
		default:
			fmt.Println("Opcion no valdia", opc)
			time.Sleep(2 * time.Second)
		}

	}
	wg.Wait()
	fmt.Println(persona)
}

func crear(llave int, valor string, wg *sync.WaitGroup) {
	fmt.Println("Insertando ", llave)
	persona.mutex.Lock()
	persona.nombre[llave] = valor
	persona.mutex.Unlock()
	wg.Done()
}
func listar(llave int, wg *sync.WaitGroup) {
	persona.mutex.Lock()
	_, existe := persona.nombre[llave]
	if existe {
		fmt.Println("Listando .. Dato Encontrado.. ", persona.nombre[llave])
	} else {
		fmt.Println("Listando .. Dato no encontrado o.. ", llave)
	}
	persona.mutex.Unlock()
	wg.Done()
}

func actualizar(llave int, valor string, wg *sync.WaitGroup) {
	persona.mutex.Lock()
	_, existe := persona.nombre[llave]
	if existe {
		fmt.Println("Actualizando .. Dato Encontrado..", persona.nombre[llave])
		persona.nombre[llave] = valor
		fmt.Println("Actualizando ..Dato Actualizaro..", persona.nombre[llave])
	} else {
		fmt.Println("Actualizando ..Dato no encontrado o.. ", llave)
	}
	persona.mutex.Unlock()
	wg.Done()
}
func eliminar(llave int, wg *sync.WaitGroup) {
	persona.mutex.Lock()
	_, existe := persona.nombre[llave]
	if existe {
		fmt.Println("Eliminando ... Dato Encontrado..", persona.nombre[llave])
		delete(persona.nombre, llave)
		fmt.Println("Eliminando ... Dato Eliminado..", llave)
	} else {
		fmt.Println("Eliminando ... Dato no encontrado o.. ", llave)
	}
	persona.mutex.Unlock()
	wg.Done()
}
