package model

type Senior struct {
	StudentNumber string   `json:"id" gorm:"primaryKey"`
	Name          string   `json:"name"`
	Class         string   `json:"class"`
	LineId        string   `json:"lineid"`
	Quota         int      `json:"quota"`
	ChildrenId    []string `json:"childrenId"`
}
