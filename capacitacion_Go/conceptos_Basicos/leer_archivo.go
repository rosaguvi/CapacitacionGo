package main 

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	data, err := ioutil.ReadFile("./hola.txt")
	if err != nil{
		fmt.Println("Algo salio Mal...")
	}else{
		fmt.Println(string(data))
	}
}