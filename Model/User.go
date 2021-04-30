package PushModel

import "time"

type Push struct {
	ID  interface{}	`gorm:"type:varchar(100);primary_key"`
	ProgressName interface{}	`gorm:"type:varchar(40);"`
	PushTitle    interface{}	`gorm:"type:TEXT"`
	PushContents interface{}	`gorm:"type:TEXT"`
	createat time.Time
}

func (Push) TableName() string {
	return "PushResqeust"
}

