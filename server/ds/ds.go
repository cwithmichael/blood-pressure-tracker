package ds

import (
	"github.com/gomodule/redigo/redis"
)

// Datastore repesents our datastore for Reading(s)
type Datastore interface {
	AllReadings() ([]*Reading, error)
	AddReading(*Reading) error
	DeleteReading(int) error
}

// DS holds a Redis connection
// and will be used to implement the Datastore interface
type DS struct {
	redis.Conn
}

// NewDS returns a new DS
func NewDS(dataSource string) (*DS, error) {
	redisConn, err := redis.Dial("tcp", dataSource)
	if err != nil {
		return nil, err
	}
	return &DS{redisConn}, nil

}
