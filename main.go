package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Data struct {
	Resultat string
}

func cesar(message string) string {

	decalage := 3
	resultat := ""

	for i := 0; i < len(message); i++ {
		resultat += string(message[i] + byte(decalage))
	}

	return resultat
}

func atbash(message string) string {

	resultat := ""

	for i := 0; i < len(message); i++ {
		resultat += string('Z' - (message[i] - 'A'))
	}

	return resultat
}

func home(w http.ResponseWriter, r *http.Request) {

	resultat := ""

	fmt.Println("METHOD:", r.Method)

	if r.Method == "POST" {

		message := r.FormValue("message")
		algo := r.FormValue("algo")

		fmt.Println("MESSAGE:", message)
		fmt.Println("ALGO:", algo)

		if algo == "cesar" {
			resultat = cesar(message)
		} else if algo == "atbash" {
			resultat = atbash(message)
		}
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("TEMPLATE ERROR:", err)
		return
	}

	data := Data{
		Resultat: resultat,
	}

	tmpl.Execute(w, data)
}

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", home)

	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
