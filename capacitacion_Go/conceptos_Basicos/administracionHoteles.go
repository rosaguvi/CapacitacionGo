package main

import (
	"bufio"
	"errors"
	"fmt" //"io/ioutil"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	//"log"
)

type Hotel struct {
	id_hotel     int
	nombre       string
	habitaciones []Habitacion
	ingresos     int
}

type Habitacion struct {
	id_habitacion, estado  string
	precio, piso, id_hotel int
}

func (htl *Hotel) administrar() (err error) {
	//var hoteles []Hotel
	salir := 0
	cnt_error := 0
	opcc := ""
	//var e error
	for salir == 0 {
		fmt.Println("***** " + htl.nombre + " **** ")
		fmt.Println("   1. Ver Habitaciones")
		fmt.Println("   2. Crear habitaciones")
		fmt.Println("   3. Reservar/Liberar Habitación")
		fmt.Println("   4. Salir")
		fmt.Println("Por favor Selecciones una Opcion")
		fmt.Scanf("%s \n", &opcc)

		switch opcc {
		case "1":
			verHabitaciones(htl.habitaciones)
			fmt.Printf(" Continuar.. ")
			fmt.Scanf("%s \n")
		case "2":
			err = htl.crearHabitacion()
			if err != nil {
				salir = 1
				fmt.Println("Error al Crear la Habitacion..")
			} else {
				fmt.Println("*** Habitacion Creada Correctamente ** ")
			}
			verHabitaciones(htl.habitaciones)
			fmt.Printf(" Continuar.. ")
			fmt.Scanf("%s \n")
		case "3":
			err = htl.reservaLiberarHb()
			if err != nil {
				salir = 1
				fmt.Println("Error al Reservar o Liberar la Habitacion..")
			} else {
				fmt.Println("*** Habitacion Actualizada Correctamente ** ")
			}
			verHabitaciones(htl.habitaciones)
			fmt.Printf(" Continuar.. ")
			fmt.Scanf("%s \n")
		case "4":
			salir = 1
		default:
			cnt_error++
			fmt.Printf(" ** Opción no valida, Cantidad de Errores: %d ** \n Enter para continuar.. ", cnt_error)
			fmt.Scanf("%s \n", &opcc)
		}

		if cnt_error >= 5 {
			salir = 1
		}
	}
	return
}

func (htl *Hotel) crearHabitacion() (err error) {
	pisos := Consultar_Pisos(htl.habitaciones)
	piso := 0
	precio := 0
	dato := "Ingrese el piso para la Habitación : "
	piso, err = CapturarEntero(dato)
	if err != nil {
		return
	}
	dato = "Ingrese el Precio para la Habitación : "
	precio, err = CapturarEntero(dato)
	if err != nil {
		return
	}
	var hb Habitacion
	_, existe := pisos[piso]
	if existe {
		hb.id_habitacion = strconv.Itoa(piso) + "-" + strconv.Itoa(pisos[piso]+1)
	} else {
		hb.id_habitacion = strconv.Itoa(piso) + "-1"
	}
	hb.estado = "Vacia"
	hb.id_hotel = htl.id_hotel
	hb.precio = precio
	hb.piso = piso
	htl.habitaciones = append(htl.habitaciones, hb)
	err = nil
	return
}

func (htl *Hotel) reservaLiberarHb() (err error) {
	salir := 0
	cnt_error := 0
	//var e error
	for salir == 0 {
		verHabitaciones(htl.habitaciones)
		id_habita := ""
		fmt.Printf(" Ingrese el Id de la Habitación a Reservar o Liberar ")
		fmt.Scanf("%s \n", &id_habita)
		for i := 0; i < len(htl.habitaciones); i++ {
			if htl.habitaciones[i].id_habitacion == id_habita {
				if htl.habitaciones[i].estado == "Vacia" {
					htl.habitaciones[i].estado = "Reservada"
					fmt.Println("Habitacion " + id_habita + "Reservada Correctamente")
				} else {
					htl.habitaciones[i].estado = "Vacia"
					fmt.Println("Habitacion " + id_habita + "Liberada Correctamente")
				}
				err = nil
				salir = 1
				break
			}
		}
		if salir == 0 {
			cnt_error++
		}
		if cnt_error == 5 {
			salir = 1
			err = errors.New("La habitacion Soliciatda no existe")
		}
	}
	return
}
func (htl *Hotel) calcularIngreso() {
	ingresos := 0
	for i := 0; i < len(htl.habitaciones); i++ {
		if htl.habitaciones[i].estado != "Vacia" {
			ingresos += htl.habitaciones[i].precio
		}
	}
	htl.ingresos = ingresos
}

func main() {
	//var hoteles []Hotel
	// var e error
	// hoteles, e = crearHoteles(hoteles)
	// if e != nil {
	// 	fmt.Println("Se presento un Error.. ", e)
	// } else {
	// 	fmt.Println("Se crearon lo siguientes Hoteles.. ")
	// 	fmt.Println(hoteles)
	// }
	var e error
	defer func() {
		if e != nil {
			fmt.Println("....")
			fmt.Println("....")
			fmt.Println("....")
			fmt.Println("Programa Terminado Con Errores ")
			fmt.Println(e)
		} else {
			fmt.Println("....")
			fmt.Println("....")
			fmt.Println("....")
			fmt.Println("Programa Correctamente ")
		}
	}()
	e = cargarMenu()
}

func cargarMenu() (e error) {
	var hoteles []Hotel
	salir := 0
	cnt_error := 0
	opcc := ""
	for salir == 0 {
		fmt.Println("***** Línea de Hoteles **** ")
		fmt.Println("   1. Ver Hoteles")
		fmt.Println("   2. Registrar Hotel")
		fmt.Println("   3. Calcular Total Ingresos")
		fmt.Println("   4. Salir")
		fmt.Println("Por favor Selecciones una Opcion")
		fmt.Scanf("%s \n", &opcc)

		switch opcc {
		case "1":
			hoteles, e = verHoteles(hoteles)
			if e != nil {
				salir = 1
				fmt.Println("Error en el Menu de Ver Hoteles\n Enter para continuar.. ")
				fmt.Scanf("%s \n")
			}
		case "2":
			hoteles, e = crearHoteles(hoteles)
			if e != nil {
				salir = 1
				fmt.Println("Error al Crear el o lo Hoteles")
			} else {
				fmt.Println("*** Hoteles creados Correctamente ** ")
			}
			listarHoteles(hoteles)
		case "3":
			hoteles = calcularIngresos(hoteles)
			listarHoteles(hoteles)
		case "4":
			salir = 1
		default:
			cnt_error++
			fmt.Printf(" ** Opción no valida, Cantidad de Errores: %d ** \n Enter para continuar.. ", cnt_error)
			fmt.Scanf("%s \n", &opcc)
		}

		if cnt_error == 5 {
			e = errors.New("Error de Manejo del sistema")
			salir = 1
		}
	}
	exportarArchivo(hoteles)
	return
}

func crearHoteles(hotels1 []Hotel) (hotels []Hotel, err error) {
	crear := "S"
	idHotel := len(hotels1) + 1
	err = nil
	hotels = hotels1
	for crear == "S" {
		var dato = ""
		cnt_pisos := 0
		cnt_habitaciones := 0
		var habitaciones []Habitacion
		var htl Hotel
		htl.id_hotel = idHotel
		htl.ingresos = 0
		fmt.Println("*** Creando Nuevo Hotel *** ")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa un nombre")
		dato, err = reader.ReadString('\n')
		htl.nombre = strings.TrimSpace(dato)
		dato = "cantidad de pisos para Hotel " + htl.nombre
		cnt_pisos, err = CapturarEntero(dato)
		if err != nil {
			break
		} else {
			dato = "cantidad de habitaciones por Piso Para el Hotel" + htl.nombre
			cnt_habitaciones, err = CapturarEntero(dato)
			if err != nil {
				break
			} else {
				habitaciones, err = crearHabitaciones(cnt_pisos, cnt_habitaciones, idHotel)
				if err != nil {
					break
				} else {
					htl.habitaciones = habitaciones
				}
			}
		}
		hotels = append(hotels, htl)
		idHotel++
		fmt.Println("Crear otro Hotel (S/N):  ")
		fmt.Scanf("%s \n", &crear)
	}
	return
}

func CapturarEntero(texto string) (valorInt int, err error) {
	for cnt_e := 0; cnt_e < 5; {
		dato := ""
		fmt.Println(texto)
		fmt.Scanf("%s \n", &dato)
		valorInt, err = strconv.Atoi(strings.TrimSpace(dato))
		if err != nil {
			cnt_e++
		} else {
			err = nil
			break
		}
	}
	return
}

func crearHabitaciones(pisos, habitaciones, id_hotel int) (habitac []Habitacion, err error) {
	for i := 1; i <= pisos; i++ {
		for j := 1; j <= habitaciones; j++ {
			var hb Habitacion
			precio := 0
			dato := "ingrese Precio para la habitacion " + strconv.Itoa(j) + " del piso  " + strconv.Itoa(i) + ": "
			precio, err = CapturarEntero(dato)
			if err != nil {
				break
			} else {
				hb.id_habitacion = strconv.Itoa(i) + "-" + strconv.Itoa(j)
				hb.estado = "Vacia"
				hb.id_hotel = id_hotel
				hb.precio = precio
				hb.piso = i
				habitac = append(habitac, hb)
				err = nil
			}
		}
	}
	return
}

func verHoteles(hoteles []Hotel) ([]Hotel, error) {
	salir := 0
	cnt_error := 0
	opcc := 0
	var e error
	for salir == 0 {
		fmt.Println("***** Hoteles Disponibles **** ")
		i := 0
		for ; i < len(hoteles); i++ {
			fmt.Println("   " + strconv.Itoa(i+1) + " " + hoteles[i].nombre)
		}
		fmt.Println("   " + strconv.Itoa(i+1) + " Regresar ")
		opcc, e = CapturarEntero("Por favor Selecciones una Opcion")
		if e != nil {
			break
		}
		if opcc > 0 && opcc <= len(hoteles) {
			e = hoteles[opcc-1].administrar()
			if e != nil {
				salir = 1
				fmt.Printf(" Error al Administrar el Hotel Enter para Continuar")
				fmt.Scanf("%s \n")
			}
		} else if opcc == (i + 1) {
			salir = 1
		} else {
			cnt_error++
			fmt.Printf(" ** Opción no valida, Cantidad de Errores: %d ** \n Enter para continuar.. ", cnt_error)
			fmt.Scanf("%s \n", &opcc)
		}

		if cnt_error >= 5 {
			salir = 1
		}
	}
	return hoteles, e
}

func verHabitaciones(hbts []Habitacion) {
	pisos := make(map[int]string)
	for _, hbt := range hbts {
		_, existe := pisos[hbt.piso]
		if existe {
			pisos[hbt.piso] = pisos[hbt.piso] + "    Habitación " + hbt.id_habitacion
		} else {
			pisos[hbt.piso] = " Habitación " + hbt.id_habitacion
		}
		if hbt.estado != "Vacia" {
			pisos[hbt.piso] = pisos[hbt.piso] + "*"
		}
	}
	for piso, contenido := range pisos {
		fmt.Println(" Piso " + strconv.Itoa(piso) + " : [" + contenido + "]")
	}
}

func Consultar_Pisos(hbts []Habitacion) map[int]int {
	pisos := make(map[int]int)
	for _, hbt := range hbts {
		_, existe := pisos[hbt.piso]
		if existe {
			pisos[hbt.piso] = pisos[hbt.piso] + 1
		} else {
			pisos[hbt.piso] = 1
		}
	}
	return pisos
}

func listarHoteles(hoteles []Hotel) {
	for i := 0; i < len(hoteles); i++ {
		fmt.Println("   " + strconv.Itoa(i+1) + " " + hoteles[i].nombre)
		fmt.Println("  Ingresos ", hoteles[i].ingresos)
		verHabitaciones(hoteles[i].habitaciones)
		fmt.Println(" ")
	}
}
func calcularIngresos(hoteles []Hotel) []Hotel {
	totalIngreso := 0
	for i := 0; i < len(hoteles); i++ {
		hoteles[i].calcularIngreso()
		fmt.Println(hoteles[i].nombre, hoteles[i].ingresos)
		totalIngreso += hoteles[i].ingresos
	}
	fmt.Println("Total Ingresos", totalIngreso)
	fmt.Println("Continuar... ")
	fmt.Scanf("%s\n")
	return hoteles
}

func exportarArchivo(hoteles []Hotel) {
	texto := ""
	for i := 0; i < len(hoteles); i++ {
		texto = texto + "Hotel;" + strconv.Itoa(hoteles[i].id_hotel) + ";" + hoteles[i].nombre + ";" + strconv.Itoa(hoteles[i].ingresos) + " \n"
		for j := 0; j < len(hoteles[i].habitaciones); j++ {
			texto = texto + "Habitacion;" + strconv.Itoa(hoteles[i].habitaciones[j].id_hotel) + ";" + hoteles[i].habitaciones[j].id_habitacion + ";" + hoteles[i].habitaciones[j].estado + ";" + strconv.Itoa(hoteles[i].habitaciones[j].precio) + ";" + strconv.Itoa(hoteles[i].habitaciones[j].piso) + " \n"
		}
	}
	b := []byte(texto)
	err := ioutil.WriteFile("./archivos/hoteles.txt", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
