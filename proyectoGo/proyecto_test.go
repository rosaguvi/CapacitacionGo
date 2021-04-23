package main

import (
	. "github.com/onsi/ginkgo" // BDD
	. "github.com/onsi/gomega" // Asserts
)

var _ = Describe("Buscar una o varias mascotas de una lista", func() {

	BeforeSuite(func() {
		lis_mascotas = nil
		lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
		lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
	})

	Context("Buscar sin ingresar Criterio", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 3", func() {
			rMascota, _ := buscarMascota("")
			Ω(len(rMascota)).Should(Equal(3))
		})
	})
	Context("buscar con el tipo de mascota perro", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 1", func() {
			rMascota, _ := buscarMascota("perro")
			Ω(len(rMascota)).Should(Equal(1))
		})
	})
	Context("buscar con el nombre de mascota Pelusa", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 1", func() {
			rMascota, _ := buscarMascota("Pelusa")
			Ω(len(rMascota)).Should(Equal(1))
		})
	})
	Context("buscar con el Id de mascota 89PE ", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 1", func() {
			rMascota, _ := buscarMascota("89PE")
			Ω(len(rMascota)).Should(Equal(1))
		})
	})
	Context("buscar cuando la palabra ingresada coincide con el nombre, el tipo de mascota y el el ID ", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 1", func() {
			rMascota, _ := buscarMascota("PE")
			Ω(len(rMascota)).Should(Equal(3))
		})
	})
	Context("buscar cuando la palabra ingresada no coincida con ningun criterio ", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 0", func() {
			rMascota, err := buscarMascota("Mascotas")
			Ω(len(rMascota)).Should(Equal(0))
			Expect(err.Error()).To(Equal("No se encontro ningun Registro"))
		})
	})
})

var _ = Describe("Agregar una mascota al listado de mascotas", func() {

	Context("crear una mascota con todos los datos correctos", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		})
		It("La cantidad de registros debe ser 2", func() {
			nMascota := Mascota{nombre: "Mariposa", categoria: "Perro", edad: 2, peso: 5}
			crearMascota(nMascota)
			Ω(len(lis_mascotas)).Should(Equal(2))
		})
	})
	Context("agregar una mascota Repetida", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		})
		It("La cantidad de registros debe ser 1 y generar error", func() {
			nMascota := Mascota{nombre: "Pelusa", categoria: "Gato", edad: 10, peso: 2}
			err := crearMascota(nMascota)
			Ω(len(lis_mascotas)).Should(Equal(1))
			Expect(err.Error()).To(Equal("Mascota ya Existe en el listado"))
		})
	})
	Context("Agregar una mascota sin nombre", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		})
		It("La cantidad de registros debe ser 1 y generar error", func() {
			nMascota := Mascota{categoria: "Perro", edad: 2, peso: 5}
			err := crearMascota(nMascota)
			Ω(len(lis_mascotas)).Should(Equal(1))
			Expect(err.Error()).To(Equal("Datos incompletos: Falta el dato para Nombre"))
		})
	})
	Context("Agregar una mascota sin categoria", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		})
		It("La cantidad de registros debe ser 1 y generar error", func() {
			nMascota := Mascota{nombre: "Mariposa", edad: 2, peso: 5}
			err := crearMascota(nMascota)
			Ω(len(lis_mascotas)).Should(Equal(1))
			Expect(err.Error()).To(Equal("Datos incompletos: Falta el dato para Categoria"))
		})
	})
	Context("Agregar una mascota sin peso", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		})
		It("La cantidad de registros debe ser 1 y generar error", func() {
			nMascota := Mascota{nombre: "Mariposa", categoria: "Gato", edad: 2}
			err := crearMascota(nMascota)
			Ω(len(lis_mascotas)).Should(Equal(1))
			Expect(err.Error()).To(Equal("Datos incompletos: Falta el dato para Peso"))
		})
	})
	Context("Agregar una mascota sin edad", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
		})
		It("La cantidad de registros debe ser 1 y generar error", func() {
			nMascota := Mascota{nombre: "Mariposa", categoria: "Gato", peso: 2}
			err := crearMascota(nMascota)
			Ω(len(lis_mascotas)).Should(Equal(2))
			Expect(err.Error()).To(Equal("Posible Error: La edad se crea en cero"))
		})
	})
})

var _ = Describe("Eliminar una mascota al listado de mascotas", func() {

	Context("Eliminar correctamente una mascota enviando el Id de la Mascota", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 2", func() {
			eliminarMascota("12NG")
			Ω(len(lis_mascotas)).Should(Equal(2))
		})
	})
	Context("Error por que el Id de la mascota no existe", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de registros debe ser 3 y generar error", func() {
			err := eliminarMascota("WG22")
			println(lis_mascotas)
			Ω(len(lis_mascotas)).Should(Equal(3))
			Expect(err.Error()).To(Equal("Error la mascota no existe"))
		})
	})
})

var _ = Describe("Buscar una o varias vacunas de una lista", func() {

	Context("Buscar sin ingresar Criterio", func() {
		BeforeEach(func() {
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de registros debe ser 2", func() {
			rVacuna, _ := buscarVacuna("")
			Ω(len(rVacuna)).Should(Equal(2))
		})
	})
	Context("buscar con el tipo de mascota perro", func() {
		BeforeEach(func() {
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de registros debe ser 1", func() {
			rVacuna, _ := buscarVacuna("perro")
			Ω(len(rVacuna)).Should(Equal(1))
		})
	})
	Context("buscar con el nombre de vacuna Vacuna 2", func() {
		BeforeEach(func() {
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de registros debe ser 1", func() {
			rVacuna, _ := buscarVacuna("Vacuna 2")
			Ω(len(rVacuna)).Should(Equal(1))
		})
	})
	Context("buscar con el Id de Vacuna 1 ", func() {
		BeforeEach(func() {
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de registros debe ser 1", func() {
			rVacuna, _ := buscarVacuna("1")
			Ω(len(rVacuna)).Should(Equal(1))
		})
	})
	Context("buscar cuando la palabra ingresada coincide con el nombre y el ID ", func() {
		BeforeEach(func() {
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de registros debe ser 2", func() {
			rVacuna, _ := buscarVacuna("2")
			Ω(len(rVacuna)).Should(Equal(2))
		})
	})
	Context("buscar cuando la palabra ingresada no coincida con ningun criterio ", func() {
		BeforeEach(func() {
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de registros debe ser 0", func() {
			rVacuna, err := buscarVacuna("vacunas")
			Ω(len(rVacuna)).Should(Equal(0))
			Expect(err.Error()).To(Equal("No se encontro ningun Registro"))
		})
	})
})

var _ = Describe("Aplicar una vacuna a una mascota", func() {

	Context("aplicar la vacuna de forma extosa", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de CC a aplicar debe ser 11.5", func() {
			resultado, _ := aplicarVacuna("1", "WG15")
			Ω(resultado).Should(Equal(float32(11.5)))
		})
	})
	Context("Error por que el Id de la mascota no existe", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de CC debe ser 0 y generar error", func() {
			cnt_cc, err := aplicarVacuna("1", "WG22")
			Ω(cnt_cc).Should(Equal(float32(0)))
			Expect(err.Error()).To(Equal("No se Puede aplicar la Vacuna: la Mascota no existe"))
		})
	})
	Context("Error por que el Id de la vacuna no existe", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de CC debe ser 0 y generar error", func() {
			cnt_cc, err := aplicarVacuna("3", "WG15")
			Ω(cnt_cc).Should(Equal(float32(0)))
			Expect(err.Error()).To(Equal("No se Puede aplicar la Vacuna: la Vacuna no existe"))
		})
	})
	Context("Error por que la vacuna no es para el tipo de mascota", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de CC debe ser 0 y generar error", func() {
			cnt_cc, err := aplicarVacuna("1", "12NG")
			Ω(cnt_cc).Should(Equal(float32(0)))
			Expect(err.Error()).To(Equal("No se Puede aplicar la Vacuna: No es para este tipo de mascota"))
		})
	})
	Context("Error por que la edad de la mascota es menor a la minima de la vacuna", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 1, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
			lis_Vacunas = nil
			lis_Vacunas = append(lis_Vacunas, Vacuna{"1", "Vacuna para perro 2", "Perro", 1.5, 2.3})
			lis_Vacunas = append(lis_Vacunas, Vacuna{"2", "vacuna 2", "gato", 2, 1.5})
		})
		It("La cantidad de CC debe ser 0 y generar error", func() {
			cnt_cc, err := aplicarVacuna("1", "WG15")
			Ω(cnt_cc).Should(Equal(float32(0)))
			Expect(err.Error()).To(Equal("No se Puede aplicar la Vacuna: La mascota no tiene la edad solicitada"))
		})
	})
})

var _ = Describe("Suministrar Jarabe a una mascota", func() {

	Context("suministrar Jarabe de forma extosa para el rango uno ", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 20, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 36, 1})
		})
		It("La cantidad de CC a aplicar debe ser 3", func() {
			resultado, _ := suministrarJarabe("12NG")
			Ω(resultado).Should(Equal(float32(3)))
		})
	})
	Context("suministrar Jarabe de forma extosa para el rango dos ", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 20, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 36, 1})
		})
		It("La cantidad de CC a aplicar debe ser 12.5", func() {
			resultado, _ := suministrarJarabe("WG15")
			Ω(resultado).Should(Equal(float32(12.5)))
		})
	})
	Context("suministrar Jarabe de forma extosa para el rango tres ", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 20, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 36, 1})
		})
		It("La cantidad de CC a aplicar debe ser 3", func() {
			resultado, _ := suministrarJarabe("89PE")
			Ω(resultado).Should(Equal(float32(3)))
		})
	})
	Context("Error por que el Id de la mascota no existe", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 2, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de CC debe ser 0 y generar error", func() {
			cnt_cc, err := suministrarJarabe("WG22")
			Ω(cnt_cc).Should(Equal(float32(0)))
			Expect(err.Error()).To(Equal("No se Puede suministrar Jarabe: la Mascota no existe"))
		})
	})
	Context("Error por que la edad de la mascota es menor a la minima para el Jarabe", func() {
		BeforeEach(func() {
			lis_mascotas = nil
			lis_mascotas = append(lis_mascotas, Mascota{"12NG", "Pelusa", "Gato", 10, 2})
			lis_mascotas = append(lis_mascotas, Mascota{"WG15", "Mariposa", "Perro", 1, 5})
			lis_mascotas = append(lis_mascotas, Mascota{"89PE", "Paco", "Loro", 3, 1})
		})
		It("La cantidad de CC debe ser 0 y generar error", func() {
			cnt_cc, err := suministrarJarabe("WG15")
			Ω(cnt_cc).Should(Equal(float32(0)))
			Expect(err.Error()).To(Equal("No se Puede suministrar Jarabe: La mascota no tiene la edad solicitada"))
		})
	})
})
