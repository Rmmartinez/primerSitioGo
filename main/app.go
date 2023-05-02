package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Users struct {
	Name   string
	Skills string
	Age    int
}

func Index(rw http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles("templates/index.html")

	user := Users{"Magui", "Desarrolladora Web", 24}

	if err != nil {
		panic(err)
	} else {
		//template.Execute(rw, nil)
		//Envía la estructura al template
		template.Execute(rw, user)
	}

}

func main() {

	http.HandleFunc("/", Index)

	//Crea el servidor
	fmt.Println("El servidor está corriendo en el puerto 3000")
	fmt.Println("Run server: http://localhost:3000")
	http.ListenAndServe("localhost:3000", nil)
}
