package models

type Publisheruserdomain struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Subaccountid uint   `json:"SubAccountID"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
}
