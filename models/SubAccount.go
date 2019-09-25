package models

type Subaccount struct {
	ID          uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Accountid   uint   `json:"AccountID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
