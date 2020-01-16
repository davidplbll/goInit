package main

import (
	"net/http"
	"fmt"
)

func main()  {
	http.HandleFunc("/",home)
	http.ListenAndServe(":3000",nil)
	fmt.Println("arriba")
}

func home(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w,r,"./index.html")
}