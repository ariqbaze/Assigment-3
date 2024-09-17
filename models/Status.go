package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Status struct {
	Water int64 `json:"water"`
	Wind  int64 `json:"wind"`
}

func (p *Status) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}

func (p *Status) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	return
}
