package tests

import (
	"github.com/cwithmichael/blood-pressure-tracker/ds"
	"github.com/cwithmichael/blood-pressure-tracker/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDS struct{}

func (mds *mockDS) AllReadings() ([]*ds.Reading, error) {
	readings := make([]*ds.Reading, 0)
	readings = append(readings, &ds.Reading{Systolic: 120, Diastolic: 70, Pulse: 60})
	readings = append(readings, &ds.Reading{Systolic: 117, Diastolic: 65, Pulse: 65})
	return readings, nil
}

func (mds *mockDS) AddReading(reading *ds.Reading) error {
	return nil
}

func (mds *mockDS) DeleteReading(id int) error {
	return nil
}

func TestReadingsIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/readings", nil)

	env := handlers.Env{DS: &mockDS{}}
	http.HandlerFunc(env.ReadingsHandler).ServeHTTP(rec, req)

	_ = []map[string]int{
		{
			"id":          0,
			"systolic":    120,
			"diastolic":   70,
			"pulse":       60,
			"readingDate": 0,
		},
		{
			"id":          0,
			"systolic":    117,
			"diastolic":   65,
			"pulse":       65,
			"readingDate": 0,
		},
	}

}
