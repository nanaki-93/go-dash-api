package model

type Entity struct {
	Collection string                 `json:"collection"`
	ID         string                 `json:"id"`
	Props      map[string]interface{} `json:"props"`
}
