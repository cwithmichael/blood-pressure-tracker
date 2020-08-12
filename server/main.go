package main

import (
	"github.com/cwithmichael/blood-pressure-tracker/ds"
	"github.com/cwithmichael/blood-pressure-tracker/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	ds, err := ds.NewDS("redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	Env := &handlers.Env{ds}
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})
	r.HandleFunc("/readings", Env.ReadingsHandler).Methods("GET")
	r.HandleFunc("/readings", Env.ReadingsPostHandler).Methods("POST")
	r.HandleFunc("/readings/{id:[0-9]+}", Env.ReadingsDeleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", c.Handler(r)))
}
