package model

import (
	"Service/service-server/db"
)

type Duty struct {
	Name     string `json:"name"`
	Position string `json:"position"`
	Phone    string `json:"phone"`
	Date     string `json:"date"`
}

//Create
func (d *Duty) NewCreate() bool {
	var result Duty
	db.GetDB().Where(&d).Find(&result)
	if result.Date != "" {
		return false
	}
	db.GetDB().Create(d)
	return true
}

//Retrieve
func (d *Duty) NewRetrieve() []Duty {
	var dutyList []Duty
	db.GetDB().Where(d).Find(&dutyList)
	return dutyList
}

// func RetrieveDutyByDate(dutyDate string) []Duty {
// 	var dutyList []Duty
// 	db.GetDB().Where(&Duty{
// 		Date: dutyDate,
// 	}).Find(&dutyList)
// 	return dutyList
// }

//Update

//Delete
