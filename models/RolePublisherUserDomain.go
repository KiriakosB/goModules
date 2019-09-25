package models

type Rolepublisheruserdomain struct {
	Roleid                uint `gorm:"primary_key;AUTO_INCREMENT" json:"RoleID"`
	Publisheruserdomainid uint `gorm:"primary_key;AUTO_INCREMENT" json:"PublisherUserDomainID"`
}
