package models

type Pauserdomainuser struct {
	Pauserdomainid uint `gorm:"primary_key;AUTO_INCREMENT" json:"PAUserDomainID"`
	Pauserid       uint `gorm:"primary_key;AUTO_INCREMENT" json:"PAUserID"`
}
