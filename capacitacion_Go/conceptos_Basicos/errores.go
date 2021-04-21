package main

import (
	"fmt"
	"errors"
)

func main(){
	resultado , err  := operar(-3,2, "suma")
	if (err != nil)	{
		fmt.Println (err) 
	}else {
		fmt.Println (resultado)
	}
}

func operar(x,y int, operacion string ) ( resultado int  , err error ){
	if x < 0 {
		err = errors.New("La variable X es menor a 0")
		return 0 , err
	}
	switch operacion {
	case "suma":
		resultado , err =  (x + y) , nil
	case "resta": 
		resultado , err =  (x-y) , nil
	case "mutiplicacion": 
		resultado , err =  (x*y) , nil
	case "division": 
		resultado , err =  (x/y) , nil
	default:
		resultado , err =  x  , nil	 
	}

	return
}