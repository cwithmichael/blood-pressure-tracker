package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/cwithmichael/blood-pressure-tracker/models"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Env struct {
	ds models.Datastore
}

func (env *Env) ReadingsHandler(w http.ResponseWriter, r *http.Request) {
	readings, err := env.ds.AllReadings()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to get readings"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readings)
}

func (env *Env) ReadingsPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	reading := &models.Reading{}
	if err := decoder.Decode(&reading); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := env.ds.AddReading(reading)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reading)
}

func (env *Env) ReadingsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	res := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(res[2])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = env.ds.DeleteReading(id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func main() {
	ds, err := models.NewDS("redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	env := &Env{ds}
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})
	r.HandleFunc("/readings", env.ReadingsHandler).Methods("GET")
	r.HandleFunc("/readings", env.ReadingsPostHandler).Methods("POST")
	r.HandleFunc("/readings/{id:[0-9]+}", env.ReadingsDeleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", c.Handler(r)))
}
