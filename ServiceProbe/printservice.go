package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func printService(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	w.Write([]byte("hello" + " " + name))
}
