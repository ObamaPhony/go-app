package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome.\n")
}

func Server() {
	router := httprouter.New()
	router.GET("/", Index)

	http.ListenAndServe(":8080", router)
}
