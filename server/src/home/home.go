package home

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("header.html", "footer.html")
	err := t.Execute(w, map[string]string{"Title": "My title", "Body": "Hi this is my body"})
	if err != nil {
		panic(err)
	}
}
