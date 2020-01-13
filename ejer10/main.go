package main

import (
	"fmt"
	us "./user"
)


type pepe struct{
	us.Usuario
}
func main()  {
	User:= new(pepe)
	User.Id=12
	fmt.Println(User)
}