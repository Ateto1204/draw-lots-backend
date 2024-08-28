package model

type Junior struct {
	StudentNumber string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Password      string `json:"pwd"`
	Class         string `json:"class"`
	Line          string `json:"line"`
	Instagram     string `json:"ig"`
	ParentId      string `json:"parent"`
}
