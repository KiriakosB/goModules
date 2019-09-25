package models

type Task struct {
	ID           uint   `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Permissionid uint   `json:"PermissionID"`
	Name         string `json:"Name"`
	Description  string `json:"Description"`
}
