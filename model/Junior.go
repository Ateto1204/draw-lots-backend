package model

type Junior struct {
	StudentNumber string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Password      string `json:"pwd"`
	Class         string `json:"class"`
	LineId        string `json:"line_id"`
	ParentId      string `json:"parent_id"`
}
