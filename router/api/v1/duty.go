package v1

import (
	"Service/service-server/db/model"
	"log"

	"github.com/gin-gonic/gin"
)

func Duty(c *gin.Context) {
	dutySwitch := c.PostForm("duty")
	if dutySwitch == "0" {
	} else if dutySwitch == "1" {
	}
	c.JSON(200, gin.H{"msg": "success"})
}

func DutyList(c *gin.Context) {
	res := make(map[string]interface{})
	date := c.Query("date")
	if date == "" {
		res["msg"] = "请输入想要查询的日期"
		c.JSON(200, res)
		return
	}

	p := &model.Duty{
		Date: date,
	}
	dutyList := p.NewRetrieve()
	res["msg"] = "成功"
	res["dutyList"] = dutyList
	c.JSON(200, res)
}

func AddDuty(c *gin.Context) {
	var body map[string]string
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Fatalln(err)
	}

	name := body["name"]
	phone := body["phone"]
	position := body["position"]
	date := body["date"]

	res := make(map[string]interface{})
	d := model.Duty{
		Name:     name,
		Phone:    phone,
		Position: position,
		Date:     date,
	}
	success := d.NewCreate()
	if !success {
		res["msg"] = "添加失败"
		c.JSON(200, res)
		return
	}

	res["msg"] = "添加成功"
	c.JSON(200, res)
	return
}
