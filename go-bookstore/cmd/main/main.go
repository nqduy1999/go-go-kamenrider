package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nqduy1999/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting server at port http//:localhost:8080\n")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
