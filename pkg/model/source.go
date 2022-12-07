package model

type Source struct {
	SourceID int64  `json:"id"`
	Name     string `json:"name"`
	Roles    string `json:"roles"`
}
