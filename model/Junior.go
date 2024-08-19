package model

type Junior struct {
	StudentNumber string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Class         string `json:"class"`
	LineId        string `json:"line_id,omitempty"`
	ParentId      string `json:"parent_id,omitempty"`
}
