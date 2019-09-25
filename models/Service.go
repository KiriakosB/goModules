package models

type Service struct {
	ID     uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Path   string `json:"Path"`
	Method string `json:"Method"`
}
