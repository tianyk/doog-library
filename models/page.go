package models

type Page struct {
	Start int         `json:"start"`
	Count int         `json:"count"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
