package main

import "fmt"
var matriz [5][7]int
func main()  {
	tabla:= [10] int {1,2,3,4,5,6,7,8,9,0}
	for i:=0; i< len(tabla); i++{
		fmt.Println(tabla[i])
	}
	matriz[4][4]=5
	fmt.Println(matriz)
	matriz2:=[] int {2,3,4}
	fmt.Println(matriz2)
	variante2()
	variante3()
}

func variante2()  {
	elementos:=[5]int{1,2,3,4,5}
	porcion:= elementos[3:]
	fmt.Println(porcion)
}

func variante3()  {
	elementos:= make([]int, 0,0)
	for i:= 0;i<100;i++{
		elementos=append(elementos,i)
	}
	fmt.Println(elementos)
}