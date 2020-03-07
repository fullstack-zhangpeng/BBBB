package model

import (
	"Service/service-server/db"
	"fmt"
)

type Duty struct {
	Name      string `json:"name"`
	Position  string `json:"position"`
	Phone     string `json:"phone"`
	Date      string `json:"dutyDate"`
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

func CreateDuty(name string, position string, phone string, date string) {
	d := Duty{
		Phone:    phone,
		Date:     date,
	}
	var result Duty

	db.GetDB().Where(&d).Find(&result)

	if result.Date == "" {
		fmt.Println("不存在")
	} else {
		fmt.Println("已存在")
	}

	return
	fmt.Printf("%v", d)
	newRecord := db.GetDB().NewRecord(d)
	fmt.Println(newRecord)
	db.GetDB().Create(Duty{
		Name:     name,
		Position: position,
		Phone:    phone,
		Date:     date,
	})
}

//Retrieve
func RetrieveDutyByDate(dutyDate string) []Duty {
	var dutyList []Duty
	db.GetDB().Where(&Duty{
		Date: dutyDate,
	}).Find(&dutyList)
	return dutyList
}

func (d *Duty) NewRetrieve() []Duty {
	var dutyList []Duty
	db.GetDB().Where(d).Find(&dutyList)
	return dutyList
}

//Update

//Delete
