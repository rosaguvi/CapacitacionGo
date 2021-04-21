package main

import (
	"fmt"
	"math"
)

func main() {

}
func fizzBuzz(numero int) string {
	if numero <= 0 {
		return "No es un numero valido"
	}
	if math.Mod(float64(numero), 3) == 0 && math.Mod(float64(numero), 5) == 0 {
		return "fizzbuzz"
	}
	if math.Mod(float64(numero), 3) == 0 {
		return "fizz"
	}
	if math.Mod(float64(numero), 5) == 0 {
		return "buzz"
	}
	return (fmt.Sprint(numero))
}
