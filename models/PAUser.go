package models

import "time"

type Pauser struct {
	ID            uint       `gorm:"primary_key;AUTO_INCREMENT" json:"ID"`
	Pauserlevelid uint       `json:"PAUserLevelID"`
	Firstname     string     `json:"FirstName"`
	Lastname      string     `json:"LastName"`
	Username      string     `json:"Username"`
	Emailaddress  string     `json:"EmailAddress"`
	Passwordhash  []byte     `json:"PasswordHash"`
	Created       *time.Time `json:"Created"`
	LastUpdated   *time.Time `json:"Last_updated"`
	LastLogin     *time.Time `json:"Last_login"`
	LoginTimes    uint       `json:"Login_times"`
	Status        string     `json:"Status"`
}
