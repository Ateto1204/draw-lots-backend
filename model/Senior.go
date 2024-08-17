package model

type Senior struct {
	StudentNumber string   `json:"id" gorm:"primaryKey"`
	Name          string   `json:"name"`
	LineId        string   `json:"lineid"`
	Quota         int      `json:"quota"`
	Child         []Junior `json:"child"`
}
