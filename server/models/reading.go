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

func (r *DS) AddReading(reading *Reading) error {
	if reading == nil {
		return errors.New("Reading is nil")
	}
	reading.ReadingDate = time.Now().Unix()
	count, err := redis.Int(r.Do("GET", "readings-count"))
	if err != nil {
		_, err = r.Do("SET", "readings-count", 0)
	}

	reading.ID = count
	countStr := strconv.Itoa(count)
	redisKey := "reading:" + countStr
	_, err = r.Do("HMSET", redisKey,
		"systolic", reading.Systolic,
		"diastolic", reading.Diastolic,
		"pulse", reading.Pulse,
		"reading-date", reading.ReadingDate)
	_, err = r.Do("SADD", "readings", redisKey)
	_, err = r.Do("INCR", "readings-count")

	if err != nil {
		return err
	}

	return nil
}

func (r *DS) DeleteReading(id int) error {
	idStr := strconv.Itoa(id)
	exists, err := redis.Bool(r.Do("HEXISTS", "reading:"+idStr, "sytolic"))
	if err != nil || !exists {
		return errors.New("Unable to delete the requested reading: " + idStr)
	}
	_, err = r.Do("DEL", "reading:"+idStr)
	if err != nil {
		return err
	}
	return nil
}
