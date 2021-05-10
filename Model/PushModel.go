package Model

import "time"

type Push struct {
	ID  			string	`gorm:"type:varchar(100);primary_key"`
	ProgressName 	string	`gorm:"type:varchar(40);"`
	PushTitle    	string	`gorm:"type:TEXT"`
	PushContents 	string	`gorm:"type:TEXT"`
	Createat 		*time.Time	`gorm:"DEFAULT: NOW()"`
}

func (Push) TableName() string {
	return "PushRequest"
}