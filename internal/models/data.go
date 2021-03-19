package models

type Data struct {
	Id     int64   `json:"id"`
	Date   string  `json:"date"`
	Views  int64   `json:"views,omitempty"`
	Clicks int64   `json:"clicks,omitempty"`
	Cost   float64 `json:"cost,omitempty"`
	Cpc    float64 `json:"cpc,omitempty"`
	Cpm    float64 `json:"cpm,omitempty"`
}
