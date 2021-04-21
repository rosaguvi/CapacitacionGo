package main

import "fmt"

func main() {
	fmt.Println("Inicio del Main")

	for i := 0; i < 20; i++ {
		go func(num int) {
			fmt.Println("Imprimiendo desde la go Rutine ", num)
		}(i)
	}

	fmt.Println("fin del main")
	fmt.Scanf("%d \n")
}
