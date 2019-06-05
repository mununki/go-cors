package main

import (
	"log"
	"net/http"

	goCors "github.com/mattdamon108/go-cors"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", goCors.Set(Index{}))

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Listening to... port 8080")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

type Index struct{}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := "This is CORS test on path: " + r.URL.Path
	w.Write([]byte(path))
}
