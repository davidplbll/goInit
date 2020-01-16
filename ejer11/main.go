package main

import "fmt"

type humano interface{
	respirar()
	pensar()
	comer()
	sexo() string

}

type animal interface{
	respirar()
	comer()
	esCarnivoro() bool
}

type vegetal interface{
	clasificationVegetal()
}

type hombre struct{
	edad int
	altura float32
	peso float32
	respirando bool
	pensando bool
	comiendo bool
	esHombre bool
}
type mujer struct{
	hombre
}

type perro struct{
	respirando bool
	comiendo bool
	carnivoro bool
}

func(h *hombre) respirar(){
	h.respirando=true
}
func(h *hombre) comer(){
	h.comiendo=true
}
func(h *hombre) pensar(){
	h.pensando=true
}
func(p *perro) respirar(){
	p.respirando=true
}
func(p *perro) comer(){
	p.comiendo=true
}
func(p *perro) esCarnivoro() bool{
	return p.carnivoro
}

func AnimalesCarnivoros(an animal) int  {
	if an.esCarnivoro(){
		return 1
	}
	return 0
}
func AnimalesRespirar(an animal)  {
	an.respirar()
}
func(h *hombre) sexo() string{
	
	if h.esHombre {
		return "hombre"
	} else {
		return "mujer"
	}
}


func humanosRespirando(hu humano)  {
	hu.respirar()
	fmt.Println(hu.sexo()	)
}

func main()  {
	// Pedro:=new(hombre)
	// Pedro.esHombre=true
	// humanosRespirando(Pedro)
	// Maria:=new(mujer)
	// humanosRespirando(Maria)
	
	totalCarnivoros:=0
	Dogo:= new(perro)
	Dogo.carnivoro=true
	AnimalesRespirar(Dogo)
	totalCarnivoros+=AnimalesCarnivoros(Dogo)
	fmt.Println(totalCarnivoros)
	fmt.Println(Dogo)

}