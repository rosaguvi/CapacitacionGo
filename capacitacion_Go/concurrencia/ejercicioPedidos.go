package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Elemento struct {
	id       int
	nombre   string
	cantidad int
}

type Producto struct {
	id        int
	nombre    string
	precio    float32
	tiempo    int
	elementos map[int]int
}

type Ingrediente struct {
	elementos []Elemento
	mutex     sync.Mutex
}

type Respuesta struct {
	estado   int
	mensaje  string
	producto Producto
}

func (p Producto) validarIngresientes(canal chan<- Respuesta) {
	var res = Respuesta{estado: 1, producto: p}
	for key, cantidad := range p.elementos {
		cnt, e := valCntIngrediente(key)
		if e != nil {
			res.estado = -1
			res.mensaje = e.Error()
		} else {
			if cnt >= cantidad {
				res.mensaje = "Cantidad Suficiente para preparar el Producto"
			} else {
				res.estado = -1
				res.mensaje = "Cantidad no es Suficiente"
			}
		}
	}
	canal <- res
}
func (p Producto) preparaProducto(canal chan<- Respuesta) {
	var res = Respuesta{estado: 1, producto: p}
	for key, cantidad := range p.elementos {
		cnt, e := valCntIngrediente(key)
		if e != nil {
			res.estado = -1
			res.mensaje = e.Error()
			break
		} else {
			if cnt < cantidad {
				res.estado = -1
				res.mensaje = "Cantidad no es Suficiente"
				break
			}
		}
	}
	if res.estado > 0 {
		for key, cantidad := range p.elementos {
			e := descontarIngrediente(key, cantidad)
			if e != nil {
				res.estado = -1
				res.mensaje = e.Error()
				break
			}
		}
	}
	if res.estado > 0 {
		fmt.Println("Preparando el Producto: ", p.nombre, " Por favor Espere..")
		time.Sleep(time.Duration(p.tiempo) * time.Second)
		res.mensaje = "Producto Preparado Correctamente.."
	}
	canal <- res
}

var carne = Elemento{1, "Carne", 5}
var pan = Elemento{2, "Pan", 10}
var salchicha = Elemento{3, "Salchicha", 3}
var tomate = Elemento{4, "Tomate", 5}
var lechuga = Elemento{5, "Lechuga", 5}
var papa = Elemento{6, "Papa", 10}
var arepa = Elemento{7, "Arepa", 3}

var hamburguesa = Producto{
	id:     1,
	nombre: "Hamburguesa",
	precio: 10000,
	tiempo: 10,
	elementos: map[int]int{
		1: 1,
		2: 2,
		4: 1,
		5: 1,
		6: 3,
	},
}

var perro = Producto{
	id:     2,
	nombre: "Perro Caliente",
	precio: 7000,
	tiempo: 6,
	elementos: map[int]int{
		3: 1,
		2: 1,
		4: 1,
		5: 1,
		6: 3,
	},
}

var arepaB = Producto{
	id:     3,
	nombre: "Arepa burger",
	precio: 8000,
	tiempo: 7,
	elementos: map[int]int{
		1: 1,
		4: 1,
		5: 1,
		6: 3,
	},
}

var elementos = [7]Elemento{carne, pan, salchicha, tomate, lechuga, papa, arepa}
var ingredientes = Ingrediente{elementos: elementos[:]}
var productos = [3]Producto{hamburguesa, perro, arepaB}

func main() {
	// fmt.Println(elementos)
	// fmt.Println(productos)
	var e error
	var pedido []Producto
	defer func() {
		if e != nil {
			fmt.Println("....")
			fmt.Println("....")
			fmt.Println("....")
			fmt.Println("Programa Terminado Con Errores ")
			fmt.Println(e)
			fmt.Println(pedido)
		} else {
			fmt.Println("....")
			fmt.Println("Programa Terminado Correctamente ")
		}
	}()
	pedido, e = menu()
	if e == nil {
		pedido, e = validarIngredientes(pedido)
		if e == nil {
			pedido, e = prepararPedido(pedido)
			if e == nil {
				vrlTotalPedido := float32(0)
				for key, producto := range pedido {
					fmt.Printf(" %d  %s  -->> %f  \n", key+1, producto.nombre, producto.precio)
					vrlTotalPedido += producto.precio
				}
				fmt.Println(" ------------------------------")
				fmt.Printf(" Total a Pagar    --->>  %f  \n", vrlTotalPedido)
			}
		}
	}
}

func menu() (seleccionados []Producto, e error) {
	salir := 0
	opcc := 0
	for salir == 0 {
		fmt.Println("***** Listado de Productos **** ")
		i := 0
		for ; i < len(productos); i++ {
			fmt.Printf(" %d. %s  -> %f \n", i+1, productos[i].nombre, productos[i].precio)
		}
		fmt.Printf(" %d. Confirmar Pedido... \n", i+1)
		opcc, e = capturarEntero("Por favor Selecciones una Opcion", 1, i+1)
		if e != nil {
			salir = 1
		}
		if opcc > 0 && opcc <= len(productos) {
			seleccionados = append(seleccionados, productos[opcc-1])
			fmt.Println("Producto Agregado Correctamente..")
		} else if opcc == (i + 1) {
			if len(seleccionados) == 0 {
				e = errors.New("No se han seleccionado Productos")
			} else {
				e = nil
			}
			salir = 1
		} else {
			fmt.Printf(" ** Opción no valida, Cantidad de Errores: %d ** \n Enter para continuar.. ")
			fmt.Scanf("%s \n")
			salir = 1
		}
	}
	return
}

func capturarEntero(texto string, min int, max int) (valorInt int, err error) {
	cnt_e := 0
	for cnt_e == 0 {
		dato := ""
		fmt.Println(texto)
		fmt.Scanf("%s \n", &dato)
		valorInt, err = strconv.Atoi(strings.TrimSpace(dato))
		if err == nil {
			if valorInt >= min && valorInt <= max {
				cnt_e = 1
				err = nil
			}
		}
	}
	return
}

func validarIngredientes(pedido []Producto) (producto []Producto, e error) {
	e = nil
	for i := 0; i < len(pedido); i++ {
		canal := make(chan Respuesta)
		go pedido[i].validarIngresientes(canal)
		respuesta := <-canal
		producto = append(producto, respuesta.producto)
		if respuesta.estado < 0 {
			e = errors.New("No hay Elementos Suficientes para Preparar el Pedido..")
			producto = nil
			producto = append(producto, respuesta.producto)
			break
		}
	}
	return
}
func prepararPedido(pedido []Producto) (producto []Producto, e error) {
	e = nil
	for i := 0; i < len(pedido); i++ {
		canal := make(chan Respuesta)
		go pedido[i].preparaProducto(canal)
		respuesta := <-canal
		if respuesta.estado < 0 {
			e = errors.New("No hay Elementos Suficientes para Preparar el Pedido..")
			producto = nil
			producto = append(producto, respuesta.producto)
			break
		} else {
			producto = append(producto, respuesta.producto)
		}
	}
	return
}

func valCntIngrediente(id int) (cnt int, e error) {
	cnt = -1
	ingredientes.mutex.Lock()
	for i := 0; i < len(ingredientes.elementos); i++ {
		if ingredientes.elementos[i].id == id {
			cnt = ingredientes.elementos[i].cantidad
			break
		}
	}
	if cnt == -1 {
		cnt = 0
		e = errors.New("El elemento no Existe..")
	}
	ingredientes.mutex.Unlock()
	return
}
func descontarIngrediente(id int, cnt int) (e error) {
	e = errors.New("El elemento no Existe..")
	ingredientes.mutex.Lock()
	for i := 0; i < len(ingredientes.elementos); i++ {
		if ingredientes.elementos[i].id == id {
			if ingredientes.elementos[i].cantidad < cnt {
				e = errors.New("Cantidad no es Suficiente..")
			} else {
				e = nil
				ingredientes.elementos[i].cantidad -= cnt
			}
			break
		}
	}
	ingredientes.mutex.Unlock()
	return
}
