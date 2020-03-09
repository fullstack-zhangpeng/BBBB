package model

import (
	"Service/service-server/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Duty struct {
	gorm.Model
	Name     string `gorm:"UNIQUE_INDEX: name_date" json:"name"`
	Position string `json:"position"`
	Phone    string `json:"phone"`
	Date     string `gorm:"UNIQUE_INDEX: name_date" json:"date"`
}

//Create
func (d *Duty) Create() error {
	dbResult := db.GetDB().Create(d)
	if dbResult.RowsAffected == 0 {
		return dbResult.Error
	}
	return nil
}

func (d *Duty) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("ID", uuid.New())
	return nil
}

//Retrieve
func (d *Duty) Retrieve() []Duty {
	var dutyList []Duty
	db.GetDB().Where(d).Find(&dutyList)
	return dutyList
}

func (d *Duty) RetrieveOne() Duty {
	var duty Duty
	db.GetDB().Where(d).First(&duty)
	fmt.Printf("%v", duty)
	return duty
}

//Update

//Delete
func (d *Duty) Delete() error {
	dbResult := db.GetDB().Delete(d)
	if dbResult.RowsAffected == 0 {
		return dbResult.Error
	}
	return nil
}
