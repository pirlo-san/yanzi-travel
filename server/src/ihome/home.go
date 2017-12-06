package ihome

import (
	"html/template"
	"idebug"
	"net/http"
	"os"
)

const packageName string = "ihome"

func HomePage(w http.ResponseWriter, r *http.Request) {

	pageFiles := getPageFiles()

	t, err := template.ParseFiles(pageFiles...)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, map[string]string{"Title": "Yanzi Travel", "Body": "Choose your best travel destination"})
	if err != nil {
		panic(err)
	}
}

func getPageFiles() (pageFiles []string) {
	pagePath := os.Getenv("GOPATH") + "/src/" + packageName
	pageFiles = []string{
		pagePath + "/header.html",
	}
	idebug.InfoLog("%s\n", pageFiles)
	return pageFiles
}
