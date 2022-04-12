package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Mi Aplicacion GO")
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("Aplicacion GO Web inicializada en puerto 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}