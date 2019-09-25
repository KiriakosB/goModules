package models

type Tasksservice struct {
	Taskid    uint `gorm:"primary_key;AUTO_INCREMENT" json:"TaskID"`
	Serviceid uint `gorm:"primary_key;AUTO_INCREMENT" json:"ServiceID"`
}
