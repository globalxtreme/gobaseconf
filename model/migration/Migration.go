package model

import "github.com/globalxtreme/gobaseconf/model"

type Migration struct {
	model.BaseModelWithoutID
	Reference string `gorm:"column:reference;type:varchar(250)"`
}

func (Migration) TableName() string {
	return "migrations"
}
