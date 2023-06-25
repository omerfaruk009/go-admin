package models

type Permission struct {
	Id   uint   `json:"id" gorm:"primaryKey"  gorm:"unique"`
	Name string `json:"name"`
}
