package model

import (
	"Service/service-server/db"
)

type Duty struct {
	id        int    `gorm:"primary_key"`
	Name      string `json:"name"`
	Position  string `json:"position"`
	Phone     string `json:"phone"`
	DutyDate  string `json:"dutyDate"`
	createdAt int
	updatedAt int
	deletedAt int
}

//Create
func CreateDuty(name string, position string, phone string, dutyDate string) {
	db.GetDB().Create(Duty{
		Name:     name,
		Position: position,
		Phone:    phone,
		DutyDate: dutyDate,
	})
}

//Retrieve
func RetrieveByDate(dutyDate string) []Duty {
	var dutyList []Duty
	db.GetDB().Where(&Duty{
		DutyDate: dutyDate,
	}).Find(&dutyList)
	return dutyList
}

//Update

//Delete
