package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Reading struct {
	ID          int   `json:"id"`
	Systolic    int   `json:"systolic"`
	Diastolic   int   `json:"diastolic"`
	Pulse       int   `json:"pulse"`
	ReadingDate int64 `json:"readingDate"`
}

var readings map[int]Reading
var count int

func ReadingsHandler(w http.ResponseWriter, r *http.Request) {
	var values []Reading
	for _, v := range readings {
		values = append(values, v)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
}

func ReadingsPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var reading Reading
	if err := decoder.Decode(&reading); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}
	reading.ReadingDate = time.Now().Unix()
	reading.ID = count
	readings[count] = reading
	count++
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reading)
}

func ReadingsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	res := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(res[2])
	_, ok := readings[id]
	if err != nil || !ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"Message": "Invalid Request"})
	} else {
		delete(readings, id)
		json.NewEncoder(w).Encode(id)
	}
}

func main() {
	readings = make(map[int]Reading)
	r := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})
	r.HandleFunc("/readings", ReadingsHandler).Methods("GET")
	r.HandleFunc("/readings", ReadingsPostHandler).Methods("POST")
	r.HandleFunc("/readings/{id:[0-9]+}", ReadingsDeleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", c.Handler(r)))
}