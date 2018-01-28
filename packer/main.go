package main

import "github.com/gobuffalo/packr"
import "net/http"

func main() {

	box := packr.NewBox("./assets")

	http.Handle("/", http.FileServer(box))
	http.ListenAndServe(":8080", nil)
}
