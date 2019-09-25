package models

type Account struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Saleshouseid uint   `json:"SalesHouseID"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
	Status       string `json:"Status"`
}
