package models

type Pauserlevel struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Name string `json:"Name"`
}
