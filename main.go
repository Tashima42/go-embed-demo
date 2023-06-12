package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed app/dist
var frontendFiles embed.FS

func main() {
	strippedFrontendFiles, err := fs.Sub(frontendFiles, "app/dist")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(strippedFrontendFiles)))
	http.ListenAndServe(":8080", nil)
}
