package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome.\n")
}

func Make(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "<h1>This is a WIP (Work-in-Progress) route.</h1>")
}

func Server() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/make", Make)

	http.ListenAndServe(":8080", router)
}
