package models

type Data struct {
	Id     int64   `json:"id"`
	Date   string  `json:"date"`
	Views  int64   `json:"views"`
	Clicks int64   `json:"clicks"`
	Cost   float64   `json:"cost"`
	Cpc    float64 `json:"cpc,omitempty"`
	Cpm    float64 `json:"cpm,omitempty"`
}
