package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Senior struct {
	StudentNumber string      `json:"id" gorm:"primaryKey"`
	Name          string      `json:"name"`
	Class         string      `json:"class"`
	LineId        string      `json:"line_id"`
	Quota         int         `json:"quota"`
	ChildrenId    StringArray `json:"children_id" gorm:"type:text" `
}

type StringArray []string

func (arr StringArray) Value() (driver.Value, error) {
	return json.Marshal(arr)
}

func (arr *StringArray) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), arr)
}
