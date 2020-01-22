package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Plan struct {
	Model

	PlanContent string `json:"planContent"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func GetPlans(maps interface {}) (count int){
	db.Model(&Plan{}).Where(maps).Count(&count)
	return
}

func ExistPlanByName(name string) bool {
	var plan Plan
	db.Select("id").Where("name = ?", name).First(&plan)
	if plan.ID > 0 {
		return true
	}
	return false
}

func AddPlan(planContent string, createdBy string) bool{
	fmt.Println(planContent)
	fmt.Println(createdBy)
	db.Create(&Plan {
		PlanContent : planContent,
		CreatedBy : createdBy,
	})
	return true
}

func (plan *Plan) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (plan *Plan) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}