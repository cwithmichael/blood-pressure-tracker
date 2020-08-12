package handlers

import (
	"encoding/json"
	"github.com/cwithmichael/blood-pressure-tracker/ds"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Env struct {
	DS ds.Datastore
}

func (env *Env) ReadingsHandler(w http.ResponseWriter, r *http.Request) {
	readings, err := env.DS.AllReadings()
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
	reading := &ds.Reading{}
	if err := decoder.Decode(&reading); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := env.DS.AddReading(reading)
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
	err = env.DS.DeleteReading(id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}
