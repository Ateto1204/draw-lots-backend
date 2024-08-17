package model

type Admin struct {
	StudentNumber string `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	Password      string `json:"pwd"`
}
