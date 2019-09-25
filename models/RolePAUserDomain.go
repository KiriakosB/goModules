package models

type Rolepauserdomain struct {
	Roleid         uint `gorm:"primary_key;AUTO_INCREMENT" json:"RoleID"`
	Pauserdomainid uint `gorm:"primary_key;AUTO_INCREMENT" json:"PAUserDomainID"`
}
