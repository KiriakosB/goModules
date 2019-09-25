package models

type Subaccountpublisher struct {
	Subaccountid uint `gorm:"primary_key;AUTO_INCREMENT" json:"SubAccountID"`
	Publisherid  uint `gorm:"primary_key;AUTO_INCREMENT" json:"PublisherID"`
}
