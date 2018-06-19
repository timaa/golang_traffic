package DTO

import (
	"time"
	"encoding/json"
)

type MyTime struct {
	time.Time
}

type TrafficDto struct {
	VisitParams map[string]string     `json:"visit_params"`
	SourceId    int64      `json:"source_id"`
	Cookie      string     `json:"cookie"`
	Page        string     `json:"page"`
	CreatedAt    MyTime     `json:"created_at"`
}

func (mt *MyTime) UnmarshalJSON(b []byte) error {
	var timeStr string
	err := json.Unmarshal(b, &timeStr)
	if err != nil {
		return err
	}

	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return err
	}

	*mt = MyTime{parsedTime}

	return nil
}
