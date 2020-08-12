package models

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
)

type Reading struct {
	ID          int   `json:"id"`
	Systolic    int   `json:"systolic"`
	Diastolic   int   `json:"diastolic"`
	Pulse       int   `json:"pulse"`
	ReadingDate int64 `json:"readingDate"`
}

// AllReadings returns all of the blood pressure readings currently in the store
func (r *DS) AllReadings() ([]*Reading, error) {
	var readings []*Reading
	keys, err := redis.Strings(r.Do("SMEMBERS", "readings"))
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		readingHash, err := redis.StringMap(r.Do("HGETALL", key))
		if err != nil {
			return nil, err
		}
		reading := &Reading{}
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

		if err != nil {
			return nil, err
		}
		readings = append(readings, reading)
	}
	return readings, nil
}

// AddReading adds a Reading to the store
func (r *DS) AddReading(reading *Reading) error {
	if reading == nil {
		return errors.New("reading is nil")
	}
	count, err := redis.Int(r.Do("GET", "readings-count"))
	if err != nil {
		// Probably means we haven't set the count yet
		_, err = r.Do("SET", "readings-count", 0)
		if err != nil {
			return err
		}
	}

	reading.ID = count
	reading.ReadingDate = time.Now().Unix()
	countStr := strconv.Itoa(count)
	redisKey := "reading:" + countStr
	// Add the reading to our store
	_, err = r.Do("HMSET", redisKey,
		"systolic", reading.Systolic,
		"diastolic", reading.Diastolic,
		"pulse", reading.Pulse,
		"reading-date", reading.ReadingDate)
	if err != nil {

		return err
	}
	// Create a set to hold our readings
	_, err = r.Do("SADD", "readings", redisKey)
	if err != nil {
		return err
	}
	// Increment the number of readings
	_, err = r.Do("INCR", "readings-count")
	if err != nil {
		return err
	}

	return nil
}

// DeleteReading removes a Reading from the store based on the supplied id
func (r *DS) DeleteReading(id int) error {
	idStr := strconv.Itoa(id)
	readingKey := "reading:" + idStr
	_, err := r.Do("DEL", readingKey)
	if err != nil {
		return err
	}
	return nil
}
