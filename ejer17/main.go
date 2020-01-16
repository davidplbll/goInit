package main

import "fmt"

var result int
func main()  {
	fmt.Println("inicio")
	result = operacionMid(sumar)(2,3)
	fmt.Println(result)
}

func sumar(a,b int) int {
	return a+b
}

func restar(a,b int) int {
	return a-b
}

func operacionMid(fun func(int,int) int )func(int,int) int{
	return func(a,b int) int{
		fmt.Println("inicio de operacion")
		return fun(a,b)
	}
}