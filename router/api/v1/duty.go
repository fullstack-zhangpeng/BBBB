package v1

import (
	"Service/service-server/db/model"
	"github.com/gin-gonic/gin"
	"log"
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
		res["code"] = -1
		res["msg"] = "请输入想要查询的日期"
		c.JSON(200, res)
		return
	}

	p := &model.Duty{
		Date: date,
	}
	dutyList := p.Retrieve()
	res["code"] = 0
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

	if (name == "") || (position == "") {
		res["code"] = -1
		res["msg"] = "添加失败，姓名或位置为空"
		c.JSON(200, res)
		return
	}

	d := model.Duty{
		Name:     name,
		Phone:    phone,
		Position: position,
		Date:     date,
	}

	exist := checkExist(d)

	if exist {
		res["code"] = -1
		res["msg"] = "该条记录已存在，请刷新查看"
		c.JSON(200, res)
		return
	}

	err := d.Create()
	if err != nil {
		res["code"] = -1
		res["msg"] = err.Error()
		c.JSON(200, res)
		return
	}
	res["code"] = 0
	res["msg"] = "添加成功"
	c.JSON(200, res)
}

func checkExist(duty model.Duty) bool {
	retrieveResult := duty.RetrieveOne()
	if retrieveResult.ID > 0 {
		return true
	} else {
		return false
	}
}

func DeleteDuty(c *gin.Context) {
	var duty model.Duty
	if err := c.ShouldBindJSON(&duty); err != nil {
		log.Fatalln(err)
	}

	res := make(map[string]interface{})

	deleteErr := duty.Delete()
	if deleteErr != nil {
		res["code"] = -1
		res["msg"] = deleteErr.Error()
		c.JSON(200, res)
		return
	}

	res["code"] = 0
	res["msg"] = "删除成功"
	c.JSON(200, res)
}
