package home

import (
	"fmt"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to yanzi travel")
}
