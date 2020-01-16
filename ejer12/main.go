package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

func main()  {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo()  {
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err!=nil{
		fmt.Println("error")
	}else{
		fmt.Println(string(archivo))
	}
}

func leoArchivo2()  {
	archivo, err := os.Open("./archivo.txt")
	if err!=nil{
		fmt.Println("error")
	}else{
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan(){
			registro :=scanner.Text()
			fmt.Println(string(registro))
		}
	}
	archivo.Close()
}
func graboArchivo(){
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err!=nil{
		fmt.Println("error")
		return
	}
	fmt.Fprintln(archivo, "Esta es la linea prro")
	archivo.Close()
}
func graboArchivo2(){
	fileName := "./nuevoArchivo.txt"
	if Append(fileName, "\n Segunda linea")==false{
		fmt.Println("Error")
	}

}

func Append(archivo string, texto string) bool  {
	arch, err := os.OpenFile(archivo,os.O_WRONLY | os.O_APPEND, 0644)
	if err!=nil{
		fmt.Println("error open")
		return false
	}
	_, err=arch.WriteString(texto)
	if err!=nil{
		fmt.Println("error")
		return false
	}
	return true
}