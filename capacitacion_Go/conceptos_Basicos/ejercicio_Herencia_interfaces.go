package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

type ifVehiculo interface {
	Encender()
	Imprimir()
	//retornarDatos()
}

type Vehiculo struct {
	marca, modelo, aho, color, estado, tipo string
	precio                                  int
	caract                                  map[string]string
	impuesto                                float32
}

type Moto struct {
	Vehiculo
	cc string
}

type Carro struct {
	Vehiculo
	puertas     int
	transmision string
}

func (m Moto) Encender() {
	if m.estado == "OK" {
		fmt.Println("Vehiculo Tipo Moto Enecedido Correctamente con CC de ", m.caract["cc"])
	} else {
		fmt.Println("Vehiculo Tipo Moto Averiado con CC de ", m.cc)
	}
}
func (c Carro) Encender() {
	if c.estado == "OK" {
		fmt.Println("Vehiculo Tipo Carro Enecedido Correctamente con Transmision de ", c.transmision)
	} else {
		fmt.Println("Vehiculo Tipo Moto Averiado con Transmision de ", c.transmision)
	}
}
func (m Moto) retornarDatos(v []Vehiculo) (motos []Moto) {
	for _, vehiculo := range v {
		if vehiculo.tipo == "moto" {
			moto, e := retornarMoto(vehiculo)
			if e != nil {
				continue
			}
			motos = append(motos, moto)
		}
	}
	return motos
}
func (c Carro) retornarDatos(v []Vehiculo) (carros []Carro) {
	for _, vehiculo := range v {
		if vehiculo.tipo == "carro" {
			carro, e := RetornarCarro(vehiculo)
			if e != nil {
				continue
			}
			carros = append(carros, carro)
		}
	}
	return carros
}
func (v *Vehiculo) calcularImpuesto() {
	v.impuesto = float32(v.precio) * float32(0.465)
}

func (v *Vehiculo) generarEstado() {
	if r.Intn(2) == 0 {
		v.estado = "Averiado"
	} else {
		v.estado = "OK"
	}
	fmt.Println(v.estado)
}

func (v Vehiculo) imprimirV() {
	fmt.Println("Tipo de Vehiculo: " + v.tipo)
	fmt.Println("Marca: " + v.marca)
	fmt.Println("Modelo: " + v.modelo)
	fmt.Println("Año: " + v.aho)
	fmt.Printf("Estado: %s \n", v.estado)
	fmt.Println("Color: " + v.color)
	fmt.Printf("Precio: %d \n ", v.precio)
	fmt.Printf("Vlr Impuesto: %f \n ", v.impuesto)
}
func (c Carro) Imprimir() {
	c.imprimirV()
	fmt.Printf("Cantidad de Puertas: %d \n", c.puertas)
	fmt.Println("Tipo de Transmisión: " + c.transmision)
}
func (m Moto) Imprimir() {
	m.imprimirV()
	fmt.Println("CC: ", m.cc)
}

// ***********************    main

func main() {
	fmt.Println("Iniciando Programa... ")
	vehiculos := cargarDatosArchivo()
	fmt.Println("en le Main ")
	fmt.Println("Imprimiendo Vehiculos sin estado")
	//ImprimirVehiculos(vehiculos)
	vehiculos = establecerEstado(vehiculos)
	vehiculos = cargarImpuesto(vehiculos)
	fmt.Println("Imprimiendo Vehiculos con estado")
	fmt.Println(vehiculos)
	ImprimirVehiculos(vehiculos)
	//ImprimirVehiculos(vehiculos)
	encenderVehiculos(vehiculos)
	// fmt.Println("vehiculos[0]")
	// fmt.Println(vehiculos[0])
	// fmt.Println("moto")
	// fmt.Println(moto)
	// fmt.Println("carro")
	// fmt.Println(carro)
	// carro.Imprimir()
}

func cargarDatosArchivo() []Vehiculo {
	var vehiculos []Vehiculo
	archivo, err := os.Open("./archivos/vehiculos.txt")
	defer func() {
		fmt.Println("\n Cerrando el Archivo..")
		archivo.Close()
		if r := recover(); r != nil {
			fmt.Println("Recuperando el Flujo ", r)
		}
	}()
	if err != nil {
		fmt.Println("Algo salio Mal...")
		panic(err)
	}
	scanner := bufio.NewScanner(archivo)
	tipoVehiculo := ""
	for scanner.Scan() {
		linea := scanner.Text()
		if linea == "Carros" || linea == "Motos" {
			tipoVehiculo = linea
			continue
		}
		datos := strings.Split(linea, ",")
		var v Vehiculo
		caract := make(map[string]string)
		if tipoVehiculo == "Motos" {
			v.marca = datos[0]
			v.modelo = datos[1]
			v.aho = datos[2]
			caract["cc"] = datos[3]
			v.caract = caract
			v.precio = traerEntero(datos[4])
			v.color = datos[5]
			v.tipo = "moto"
		}
		if tipoVehiculo == "Carros" {
			v.marca = datos[0]
			v.modelo = datos[1]
			v.aho = datos[2]
			v.precio = traerEntero(datos[3])
			caract["puertas"] = datos[4]
			caract["transmision"] = datos[6]
			v.color = datos[5]
			v.caract = caract
			v.tipo = "carro"
		}
		// v.calcularImpuesto()
		//v.generarEstado()
		fmt.Println("en el llamado ", v.estado)
		vehiculos = append(vehiculos, v)
	}
	return vehiculos
}

func traerEntero(valor string) int {
	valorInt, err := strconv.Atoi(strings.TrimSpace(valor))
	if err != nil {
		return 0
	} else {
		return valorInt
	}
}
func RetornarCarro(v Vehiculo) (c Carro, err error) {
	if v.tipo == "carro" {
		c.tipo = v.tipo
		c.marca = v.marca
		c.modelo = v.modelo
		c.estado = v.estado
		c.aho = v.aho
		c.precio = v.precio
		c.puertas = traerEntero(v.caract["puertas"])
		c.transmision = v.caract["transmision"]
		c.color = v.color
		c.impuesto = v.impuesto
	} else {
		err = errors.New("El Vehiculo no es un carro")
		return c, err
	}
	return c, nil
}
func retornarMoto(v Vehiculo) (m Moto, err error) {
	if v.tipo == "moto" {
		m.tipo = v.tipo
		m.marca = v.marca
		m.modelo = v.modelo
		m.estado = v.estado
		m.aho = v.aho
		m.precio = v.precio
		m.cc = v.caract["cc"]
		m.color = v.color
		m.impuesto = v.impuesto
	} else {
		err = errors.New("El Vehiculo no es una moto")
		return m, err
	}
	return m, nil
}

func establecerEstado(v []Vehiculo) []Vehiculo {

	for i := 0; i < len(v); i++ {
		v[i].generarEstado()
		fmt.Println("actualizando estado", v[i])
	}
	return v
}
func cargarImpuesto(v []Vehiculo) []Vehiculo {
	for i := 0; i < len(v); i++ {
		v[i].calcularImpuesto()
		//fmt.Println("actualizando impuesto", v[i])
	}
	return v
}
func ImprimirVehiculos(v []Vehiculo) {
	for i := 0; i < len(v); i++ {
		fmt.Println(" ")
		fmt.Println("vehiculo ", i)
		if v[i].tipo == "moto" {
			m, _ := retornarMoto(v[i])
			m.Imprimir()
		} else if v[i].tipo == "carro" {
			c, _ := RetornarCarro(v[i])
			c.Imprimir()
		} else {
			continue
		}
	}
}
func encenderVehiculos(v []Vehiculo) {
	for i := 0; i < len(v); i++ {
		if v[i].tipo == "moto" {
			m, _ := retornarMoto(v[i])
			encenderVehiculo(m)
		} else if v[i].tipo == "carro" {
			c, _ := RetornarCarro(v[i])
			encenderVehiculo(c)
		} else {
			continue
		}
	}
}

func encenderVehiculo(iv ifVehiculo) {
	iv.Encender()
}
