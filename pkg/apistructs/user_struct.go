package apistructs

import "gorm.io/gorm"

type ContactInfo struct {
	ContactID int    `json:"contactid" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	MobileNo  uint	`json:"mobileno"`
}

// TODO hash passord
type User struct {
	gorm.Model
	ContactID int `json:"contactid" gorm:"primary_key, unique;not null"`
	ContactDetails []ContactInfo `json:"contactdetails" gorm:"ForeignKey:ContactID"`
	UserName       string `json:"username"`
	Email          string `json:"email"`
	Age            int `json:"age"`
	Password       string `json:"password"`
}
