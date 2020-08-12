package models

import (
	"github.com/gomodule/redigo/redis"
)

type Datastore interface {
	AllReadings() ([]*Reading, error)
	AddReading(*Reading) error
	DeleteReading(int) error
}

type DS struct {
	redis.Conn
}

func NewDS(dataSource string) (*DS, error) {
	redisConn, err := redis.Dial("tcp", dataSource)
	if err != nil {
		return nil, err
	}
	return &DS{redisConn}, nil

}
