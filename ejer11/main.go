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

func(h *hombre) respirar(){
	h.respirando=true
}
func(h *hombre) comer(){
	h.comiendo=true
}
func(h *hombre) pensar(){
	h.pensando=true
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
	Pedro:=new(hombre)
	Pedro.esHombre=true
	humanosRespirando(Pedro)
	Maria:=new(mujer)
	humanosRespirando(Maria)
}