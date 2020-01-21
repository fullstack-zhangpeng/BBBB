package model

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
	db.Create(&Plan {
		PlanContent : planContent,
		CreatedBy : createdBy,
	})
	return true
}