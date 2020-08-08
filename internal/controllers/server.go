package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Application", "ObamaPhony")
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "{\"result\": \"PONG\"}")
}

func Server(bindInter string) {
	r := mux.NewRouter()
	r.HandleFunc("/", handleRoot).Methods("GET")

	http.ListenAndServe(bindInter, r)
}
