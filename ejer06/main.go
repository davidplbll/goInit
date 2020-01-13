package main

import "fmt"

func main()  {
	fmt.Println(tres(5,6))
	//numer,estado:=dos(9)
	//fmt.Println(numer)
	//fmt.Println(estado)
}

func uno(numero int) (z int) {
	z=5
	return
}

func dos(numero int) (int,bool) {
	if numero==1{
		return numero *2 , true
	}else{
		return numero *2 , false
	}
}

func tres(numero ...int) int  {
	total:=0
	for _, num := range numero {
		total+=num
	}
	return total
}