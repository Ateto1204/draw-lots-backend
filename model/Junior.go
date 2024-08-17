package model

type Junior struct {
	StudentNumber string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	LineId        string `json:"lineid,omitempty"`
	Parent        Senior `json:"parent,omitempty"`
}
