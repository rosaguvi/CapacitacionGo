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

type Mascota struct {
	id, nombre, categoria string
	edad, peso            float32
}
type Vacuna struct {
	id, nombre, categoria string
	edad_min, cc_kg       float32
}
type Jarabe struct {
	edad_min, edad_max, cc_kg float32
}

var LetrasId = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

var lis_mascotas []Mascota
var lis_Vacunas []Vacuna

var jarabe_r1 = Jarabe{2.9999, 12, 1.5}
var jarabe_r2 = Jarabe{12, 30, 2.5}
var jarabe_r3 = Jarabe{30, 9999999, 3}

var jarabe = [3]Jarabe{jarabe_r1, jarabe_r2, jarabe_r3}

// var masct = Mascota{id: "12NG", nombre: "Pelusa", categoria: "Gato", edad: 10, peso: 2}
// lis_mascotas = append (lis_mascotas, masct)
// mascota = Mascota{"WG15", "Mariposa", "Perro", 2, 5}
// lis_mascotas = append (lis_mascotas, mascota)
// mascota = Mascota{"89PE", "Paco", "Loro", 3, 1}
// lis_mascotas = append (lis_mascotas, mascota)

func main() {
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
	cargarDatosMascotas()
	cargarDatosVacunas()
	e = generarMenu()

}

func generarMenu() (e error) {
	salir := 0
	opcc := 0
	for salir == 0 {
		fmt.Println("1. Ver todas las mascotas")
		fmt.Println("2. Buscar Mascotas ")
		fmt.Println("3. Registrar mascota ")
		fmt.Println("4. Dar salida ")
		fmt.Println("5. Ver Vacunas ")
		fmt.Println("6. Aplicar vacuna ")
		fmt.Println("7. Suministrar Jarabe ")
		fmt.Println("8. Salir ")
		opcc, e = capturarEntero("Por favor Selecciones una Opcion", 1, 8)
		if opcc == 8 || e != nil {
			salir = 1
		}
		switch opcc {
		case 1:
			listarMascotas("")
		case 2:
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Ingresa el Criterio de Busqueda")
			dato, _ := reader.ReadString('\n')
			dato = strings.TrimSpace(dato)
			listarMascotas(dato)
		case 3:
			nuevaMascota()
		case 4:
			darSalidaMascota()
		case 5:
			listarVacunas()
		case 6:
			aplicarVacunaMascota()
		case 7:
			suministrarJarabeMascota()
		}
	}
	return
}

func buscarMascota(cadena string) (rMascota []Mascota, err error) {
	if cadena == "" {
		rMascota = lis_mascotas
	} else {
		for _, dato := range lis_mascotas {
			if strings.Contains(strings.ToUpper(dato.categoria), strings.ToUpper(cadena)) ||
				strings.Contains(strings.ToUpper(dato.nombre), strings.ToUpper(cadena)) ||
				strings.Contains(strings.ToUpper(dato.id), strings.ToUpper(cadena)) {
				rMascota = append(rMascota, dato)
			}
		}
	}
	if len(rMascota) == 0 {
		err = errors.New("No se encontro ningun Registro")
	}
	return

}
func crearMascota(nMascota Mascota) (err error) {
	err = validarDatos(nMascota)
	if err != nil {
		return
	}
	for _, dato := range lis_mascotas {
		if strings.ToUpper(dato.categoria) == strings.ToUpper(nMascota.categoria) &&
			strings.ToUpper(dato.nombre) == strings.ToUpper(nMascota.nombre) &&
			dato.edad == nMascota.edad &&
			dato.peso == nMascota.peso {
			err = errors.New("Mascota ya Existe en el listado")
			return
		}
	}
	nMascota.id = generarIdentificador()
	lis_mascotas = append(lis_mascotas, nMascota)
	if nMascota.edad == 0 {
		err = errors.New("Posible Error: La edad se crea en cero")
	}
	return
}

func eliminarMascota(idMascota string) (err error) {
	for key, dato := range lis_mascotas {
		if dato.id == idMascota {
			ElimiarDatoMascota(lis_mascotas, key)
			return
		}
	}
	err = errors.New("Error la mascota no existe")
	return
}

func validarDatos(vMascota Mascota) (err error) {
	txtError := ""
	if vMascota.nombre == "" {
		txtError = "Datos incompletos: Falta el dato para Nombre"
	}
	if vMascota.categoria == "" {
		if txtError == "" {
			txtError = "Datos incompletos: Falta el dato para Categoria"
		} else {
			txtError = txtError + ", Categoria"
		}
	}
	if vMascota.peso == 0 {
		if txtError == "" {
			txtError = "Datos incompletos: Falta el dato para Peso"
		} else {
			txtError = txtError + ", Peso"
		}
	}
	if txtError == "" {
		return
	}
	err = errors.New(txtError)
	return
}

func generarIdentificador() (codigo string) {
	codigo = ""
	for len(codigo) < 4 {
		id := r.Intn(72)
		for key, letra := range LetrasId {
			if key == id {
				codigo = codigo + string(letra)
			}
		}
	}
	return
}

func ElimiarDatoMascota(listaMascotas []Mascota, index int) {
	lis_mascotas = nil
	lis_mascotas = append(listaMascotas[:index], listaMascotas[index+1:]...)
}

func buscarVacuna(cadena string) (rVacuna []Vacuna, err error) {
	if cadena == "" {
		rVacuna = lis_Vacunas
	} else {
		for _, dato := range lis_Vacunas {
			if strings.Contains(strings.ToUpper(dato.categoria), strings.ToUpper(cadena)) ||
				strings.Contains(strings.ToUpper(dato.nombre), strings.ToUpper(cadena)) ||
				strings.Contains(strings.ToUpper(dato.id), strings.ToUpper(cadena)) {
				rVacuna = append(rVacuna, dato)
			}
		}
	}
	if len(rVacuna) == 0 {
		err = errors.New("No se encontro ningun Registro")
	}
	return
}
func buscarVacunaID(idVacuna string) (rVacuna []Vacuna, err error) {
	if idVacuna == "" {
		err = errors.New("No hay datos de entrada")
	} else {
		for _, dato := range lis_Vacunas {
			if dato.id == idVacuna {
				rVacuna = append(rVacuna, dato)
			}
		}
	}
	if len(rVacuna) == 0 {
		err = errors.New("la Vacuna no existe")
	}
	return
}
func buscarMascotaID(idMascota string) (rMascota []Mascota, err error) {
	if idMascota == "" {
		err = errors.New("No hay datos de entrada")
	} else {
		for _, dato := range lis_mascotas {
			if dato.id == idMascota {
				rMascota = append(rMascota, dato)
			}
		}
	}
	if len(rMascota) == 0 {
		err = errors.New("la Mascota no existe")
	}
	return
}
func aplicarVacuna(idVacuna, idMascota string) (cnt_cc float32, err error) {
	if idVacuna == "" || idMascota == "" {
		err = errors.New("No se Procesara, se Deben enviar los dos ID´s")
	} else {
		mascota, e := buscarMascotaID(idMascota)
		if e != nil {
			err = errors.New("No se Puede aplicar la Vacuna: " + fmt.Sprintf(e.Error()))
			return
		}
		vacuna, ev := buscarVacunaID(idVacuna)
		if ev != nil {
			err = errors.New("No se Puede aplicar la Vacuna: " + fmt.Sprintf(ev.Error()))
			return
		}
		if mascota[0].categoria != vacuna[0].categoria {
			err = errors.New("No se Puede aplicar la Vacuna: No es para este tipo de mascota")
			return
		}
		if mascota[0].edad < vacuna[0].edad_min {
			err = errors.New("No se Puede aplicar la Vacuna: La mascota no tiene la edad solicitada")
			return
		}
		cnt_cc = vacuna[0].cc_kg * (mascota[0].peso / 2)

	}
	return
}
func suministrarJarabe(idMascota string) (cnt_cc float32, err error) {
	if idMascota == "" {
		err = errors.New("No se Procesara, se Deben enviar el ID")
	} else {
		mascota, e := buscarMascotaID(idMascota)
		if e != nil {
			err = errors.New("No se Puede suministrar Jarabe: " + fmt.Sprintf(e.Error()))
			return
		}
		if mascota[0].edad <= jarabe[0].edad_min {
			err = errors.New("No se Puede suministrar Jarabe: La mascota no tiene la edad solicitada")
			return
		}
		for _, ran_jarabe := range jarabe {
			if mascota[0].edad > ran_jarabe.edad_min && mascota[0].edad <= ran_jarabe.edad_max {
				cnt_cc = ran_jarabe.cc_kg * (mascota[0].peso / 2)
				return
			}
		}
		cnt_cc = jarabe[2].cc_kg * (mascota[0].peso / 2)
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
func listarMascotas(datoBusquda string) {
	mascotas, e := buscarMascota(datoBusquda)
	if e != nil {
		fmt.Println("Se ha generado un error y no se listaran las mascotas...", e)
		return
	}
	fmt.Println("****** Listado de Mascotas *****")
	fmt.Println("ID     Nombre       Tipo       Edad         Peso Libras")
	for _, mascota := range mascotas {
		fmt.Printf("%s   %-12s %-10s %-11f  %f\n", mascota.id, mascota.nombre, mascota.categoria, mascota.edad, mascota.peso)
	}
	fmt.Println("Enter para continuar.. ")
	fmt.Scanf("%s \n")
	return
}
func cargarDatosMascotas() {
	archivo, err := os.Open("./Archivos/mascotas.txt")
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
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
		datos := strings.Split(linea, ", ")
		var mascota Mascota
		mascota.nombre = datos[0]
		mascota.categoria = datos[1]
		edad, e := strconv.ParseFloat(strings.TrimSpace(datos[2]), 6)
		if e != nil {
			fmt.Println(e)
			continue
		}
		mascota.edad = float32(edad)
		peso, e1 := strconv.ParseFloat(strings.TrimSpace(datos[3]), 6)
		if e1 != nil {
			fmt.Println(e1)
			continue
		}
		mascota.peso = float32(peso)
		e = crearMascota(mascota)
		if e != nil {
			fmt.Println(e)
			continue
		}
	}
	archivo.Close()
}

func cargarDatosVacunas() {
	archivo, err := os.Open("./Archivos/vacunas.txt")
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
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
		datos := strings.Split(linea, ", ")
		var vacuna Vacuna
		vacuna.id = datos[0]
		vacuna.nombre = datos[1]
		vacuna.categoria = datos[2]
		dato, e := strconv.ParseFloat(strings.TrimSpace(datos[3]), 6)
		if e != nil {
			fmt.Println(e)
			continue
		}
		vacuna.edad_min = float32(dato)
		dato, e = strconv.ParseFloat(strings.TrimSpace(datos[4]), 6)
		if e != nil {
			fmt.Println(e)
			continue
		}
		vacuna.cc_kg = float32(dato)
		lis_Vacunas = append(lis_Vacunas, vacuna)
	}
	archivo.Close()
}

func nuevaMascota() {
	crear := "S"
	for crear == "S" {
		var dato = ""
		var mascota Mascota
		fmt.Println("*** Ingresando Nueva Mascota *** ")
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Ingresa un nombre para la Mascota")
		dato, _ = reader.ReadString('\n')
		mascota.nombre = strings.TrimSpace(dato)
		fmt.Println("Ingresa el Tipo de Mascota")
		dato, _ = reader.ReadString('\n')
		mascota.categoria = strings.TrimSpace(dato)
		dato = "Edad en meses para la mascota " + mascota.nombre
		mascota.edad, _ = Capturarfloat(dato)
		dato = "peso en Libras la mascota " + mascota.nombre
		mascota.peso, _ = Capturarfloat(dato)
		crearMascota(mascota)
		listarMascotas(mascota.nombre)
		fmt.Println("Ingresar otra Mascota (S/N):  ")
		fmt.Scanf("%s \n", &crear)
	}
	return
}
func darSalidaMascota() {
	dato := ""
	fmt.Println("*** Generando Salida Mascota *** ")
	fmt.Println("Ingresa el Id de la Mascota")
	fmt.Scanf("%s \n", &dato)
	e := eliminarMascota(dato)
	if e != nil {
		fmt.Println("Error al Eliminar el Registro: ", fmt.Sprintf(e.Error()))
	} else {
		fmt.Println("Mascota Eliminada Correctamente.. \n Enter para Continuar")
		fmt.Scanf("%s \n", &dato)
	}
}

func Capturarfloat(texto string) (valorFloat float32, err error) {
	for {
		dato := ""
		fmt.Println(texto)
		fmt.Scanf("%s \n", &dato)
		valorFlt, e := strconv.ParseFloat(strings.TrimSpace(dato), 32)
		if e == nil {
			valorFloat = float32(valorFlt)
			break
		}
	}
	return
}

func listarVacunas() {
	vacunas, e := buscarVacuna("")
	if e != nil {
		fmt.Println("Se ha generado un error y no se listaran las Vacunas...", e)
		return
	}
	fmt.Println("****** Listado de Vacunas *****")
	fmt.Println("ID     Nombre       Tipo Mascota    Edad Minima   CC_Kg")
	for _, vacuna := range vacunas {
		fmt.Printf("%-4s   %-12s %-16s %-11f  %f\n", vacuna.id, vacuna.nombre, vacuna.categoria, vacuna.edad_min, vacuna.cc_kg)
	}
	fmt.Println("Enter para continuar.. ")
	fmt.Scanf("%s \n")
	return
}

func aplicarVacunaMascota() {
	vacuna := ""
	mascota := ""
	fmt.Println("*** Aplicar Vacuna a una Mascota *** ")
	fmt.Println("Ingresa el Id de la Vacuna")
	fmt.Scanf("%s \n", &vacuna)
	fmt.Println("Ingresa el Id de la Mascota")
	fmt.Scanf("%s \n", &mascota)
	dosis, e := aplicarVacuna(vacuna, mascota)
	if e != nil {
		fmt.Println("Error al Aplicar la Vacuna: ", fmt.Sprintf(e.Error()))
	} else {
		fmt.Printf("La dosis a Aplicar es: %f \n Enter para Continuar", dosis)
		fmt.Scanf("%s \n")
	}
}
func suministrarJarabeMascota() {
	mascota := ""
	fmt.Println("*** Suministrar Jarabe a una Mascota *** ")
	fmt.Println("Ingresa el Id de la Mascota")
	fmt.Scanf("%s \n", &mascota)
	dosis, e := suministrarJarabe(mascota)
	if e != nil {
		fmt.Println("Error al Suministrar Jarabe: ", fmt.Sprintf(e.Error()))
	} else {
		fmt.Printf("La dosis a Aplicar es: %f Mililitros \n Enter para Continuar", dosis)
		fmt.Scanf("%s \n")
	}
}
