package models

type Publisheruserdomainuser struct {
	Publisheruserdomainid uint `gorm:"primary_key;AUTO_INCREMENT" json:"PublisherUserDomainID"`
	Publisheruserid       uint `gorm:"primary_key;AUTO_INCREMENT" json:"PublisherUserID"`
}
