package main

import "fmt"

func main(){
	i:=1
	for i<10{
		fmt.Println(i)
		i++
	}
	for l:=1;l<=10;l++{
		fmt.Println(l)
	}
	for {
		fmt.Println(i)
		break
	}

	var p=1
	for p<10{
		fmt.Printf("valor de p %d",p)
		if p==1 {
			p++
			continue
		}
		p++
	}
	var m=0
	RUTINA:
		fmt.Printf("valor de p ")
		for m=0 ; m < 10; m++ {
			if m==1{
				m++
				goto RUTINA
			}
		}
}