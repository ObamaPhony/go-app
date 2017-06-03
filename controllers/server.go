package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome.\n")
}

func Server(bindInter string) {
	router := httprouter.New()
	router.GET("/", Index)

	http.ListenAndServe(bindInter, router)
}
