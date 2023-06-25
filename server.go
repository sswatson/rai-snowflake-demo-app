package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type CustomFileServer struct {
	root http.FileSystem
}

func (c *CustomFileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := os.Stat(filepath.Join(".", r.URL.Path))
	if os.IsNotExist(err) {
		// If the file doesn't exist, default to index.html
		http.ServeFile(w, r, "./index.html")
		return
	}
	http.FileServer(c.root).ServeHTTP(w, r)
}

func main() {
	fs := &CustomFileServer{http.Dir(".")}
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
