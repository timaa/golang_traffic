package DTO

import "time"

type TrafficDto struct {
	VisitParams string     `json:"visit_params"`
	SourceId    int64      `json:"source_id"`
	Cookie      string     `json:"cookie"`
	Page        string     `json:"page"`
	CratedAt    time.Time  `json:"created_at"`
}
