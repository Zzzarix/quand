package domain

type History []Log

type Log struct {
	Key      string   `json:"key"`
	Date     string   `json:"date"`
	Question Question `json:"question"`
}
