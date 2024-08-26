package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Senior struct {
	StudentNumber string      `json:"id" gorm:"primaryKey"`
	Name          string      `json:"name"`
	Password      string      `json:"pwd"`
	Class         string      `json:"class"`
	LineId        string      `json:"line_id"`
	Quota         int         `json:"quota"`
	ChildrenId    StringArray `json:"children_id" gorm:"serializer:json" `
}

type StringArray []string

func (arr StringArray) Value() (driver.Value, error) {
	return json.Marshal(arr)
}

func (arr *StringArray) Scan(value interface{}) error {
	var byteData []byte

	switch v := value.(type) {
	case string:
		byteData = []byte(v)
	case []byte:
		byteData = v
	default:
		return errors.New("unsupported data type for StringArray")
	}

	return json.Unmarshal(byteData, arr)
}

func (arr StringArray) Append(value string) *StringArray {
	arr = append(arr, value)
	return &arr
}
