package controllers

import (
	"fmt"
	"net/http"
)

func notFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not Found: %#v", req.RequestURI)
}
