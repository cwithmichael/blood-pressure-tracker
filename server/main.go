package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
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

var redisConn redis.Conn

func ReadingsHandler(w http.ResponseWriter, r *http.Request) {
	var readings []Reading
	keys, err := redis.Strings(redisConn.Do("SMEMBERS", "readings"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "failed to get readings"})
		return
	}
	for _, key := range keys {
		readingHash, err := redis.StringMap(redisConn.Do("HGETALL", key))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "failed to get readings"})
			return
		}
		reading := Reading{}
		for k, v := range readingHash {
			switch k {
			case "systolic":
				reading.Systolic, err = strconv.Atoi(v)
			case "diastolic":
				reading.Diastolic, err = strconv.Atoi(v)
			case "pulse":
				reading.Pulse, err = strconv.Atoi(v)
			case "reading-date":
				date, _ := strconv.Atoi(v)
				reading.ReadingDate = int64(date)
			}
		}
		readings = append(readings, reading)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(readings)
}

func ReadingsPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	var reading Reading
	if err := decoder.Decode(&reading); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	reading.ReadingDate = time.Now().Unix()
	count, err := redis.Int(redisConn.Do("GET", "readings-count"))
	if err != nil {
		_, err = redisConn.Do("SET", "readings-count", 0)
	}

	reading.ID = count
	countStr := strconv.Itoa(count)
	redisKey := "reading:" + countStr
	_, err = redisConn.Do("HMSET", redisKey,
		"systolic", reading.Systolic,
		"diastolic", reading.Diastolic,
		"pulse", reading.Pulse,
		"reading-date", reading.ReadingDate)
	_, err = redisConn.Do("SADD", "readings", redisKey)
	_, err = redisConn.Do("INCR", "readings-count")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reading)
}

func ReadingsDeleteHandler(w http.ResponseWriter, r *http.Request) {
	res := strings.Split(r.URL.Path, "/")
	id := res[2]
	exists, err := redis.Bool(redisConn.Do("HEXISTS", "reading:"+id, "sytolic"))
	if err != nil || !exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		_, err := redisConn.Do("DEL", "reading:"+id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	var err error
	redisConn, err = redis.Dial("tcp", "redis:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer redisConn.Close()
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
