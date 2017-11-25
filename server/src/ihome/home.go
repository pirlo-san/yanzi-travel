package ihome

import (
	"html/template"
	"idebug"
	"net/http"
	"os"
)

var packageName string = "ihome"

func HomePage(w http.ResponseWriter, r *http.Request) {

	pageFiles := getPageFiles()
	t, _ := template.ParseFiles(pageFiles...)
	err := t.Execute(w, map[string]string{"Title": "My title", "Body": "Hi this is my body"})
	if err != nil {
		panic(err)
	}
}

func getPageFiles() (pageFiles []string) {
	pagePath := os.Getenv("GOPATH") + "/src/" + packageName
	pageFiles = []string{
		pagePath + "/footer.html",
		pagePath + "/header.html",
	}
	idebug.InfoLog("%s\n", pageFiles)
	return pageFiles
}
