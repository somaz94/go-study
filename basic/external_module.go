package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func external_modules() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandlers)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func HomeHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, My Name is Somaz!"))
}
