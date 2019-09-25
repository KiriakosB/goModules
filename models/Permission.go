package models

type Permission struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
