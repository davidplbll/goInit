package main

import (
	"fmt"
	// "io/ioutil"
	 "log"
	//"os"
)

func main()  {
	// //defer: se ejecuta al final de la funcion
	// archivo :="prueba.txt"
	// f, err:= os.Open(archivo)
	// defer f.Close()
	// if err != nil{
	// 	fmt.Println("error")
	// 	fmt.Println("error")
	// 	os.Exit(1)
	// }
	ejempanic()
	//panic
}

func ejempanic()  {
	defer func (){
		reco := recover()
		if reco !=nil{
			log.Fatalf("error %v", reco )
		}
	}()

	a := 1
	if a==1{
		panic("valor panic")
	}
	fmt.Println("error")
}