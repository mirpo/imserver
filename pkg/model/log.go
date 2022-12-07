package model

type Log struct {
	SourceID  int64  `json:"source_id"`
	CreatedAT int64  `json:"created_at"`
	Metrics   string `json:"metrics"`
}
