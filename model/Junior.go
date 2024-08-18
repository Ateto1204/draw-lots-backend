package model

type Junior struct {
	StudentNumber string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Class         string `json:"class"`
	LineId        string `json:"lineid,omitempty"`
	ParentId      string `json:"parentId,omitempty"`
}
