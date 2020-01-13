package main

import "fmt"
var calculo func (int,int) int =func(num1 int, num2 int) int{
	return num1+num2
}
func main()  {
	fmt.Println(calculo(1,2))
	calculo=func(num1 int, num2 int) int{
		return num1-num2
	}
	fmt.Println(calculo(1,2))
	tabla:=Tabla(2)
	for i := 0; i < 10; i++ {
		fmt.Println(tabla())
	}
}

func Tabla(valor int) func() int  {
	numero:=valor
	secuencia:=0
	return func() int{
		secuencia++
		return numero*secuencia
	}
}

