package main

import "fmt"

func main()  {
	paises := make(map[string]string, 5)
	paises["mexico"]="DFD"
	fmt.Println(paises["mexico"])
	campeonato:=map[string]int{
		"colombia":10}
	fmt.Println(campeonato)
	campeonato["river"]=23
	fmt.Println(campeonato)
	campeonato["river"]=63
	fmt.Println(campeonato)
	delete(campeonato,"river")
	fmt.Println(campeonato)
	
	for equipo, puntage := range campeonato{
		fmt.Println(equipo)
		fmt.Println(puntage)
	}
	puntage, existe := campeonato["river"]
	fmt.Println(puntage)
	fmt.Println(existe)

}