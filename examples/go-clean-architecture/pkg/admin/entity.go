package admin

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Password    string `json:"password,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Email       string `json:"email,omitempty"`
	Address     string `json:"address,omitempty"`
	DisplayPic  string `json:"display_pic,omitempty"`
}
