package home

import (
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/root/gopath/src/home/header.html", "/root/gopath/src/home/footer.html")
	err := t.Execute(w, map[string]string{"Title": "My title", "Body": "Hi this is my body"})
	if err != nil {
		panic(err)
	}
}
