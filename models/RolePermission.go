package models

type Rolepermission struct {
	Roleid       uint `gorm:"primary_key;AUTO_INCREMENT" json:"RoleID"`
	Permissionid uint `gorm:"primary_key;AUTO_INCREMENT" json:"PermissionID"`
}
