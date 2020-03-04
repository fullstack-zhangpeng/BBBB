package v1

import (
	"Service/service-server/cornManager"
	"Service/service-server/db/model"

	"github.com/gin-gonic/gin"
)

func DutyList(c *gin.Context) {
	res := make(map[string]interface{})
	dutyDate := c.Query("dutyDate")
	if dutyDate == "" {
		res["msg"] = "请输入想要查询的日期"
		c.JSON(200, res)
		return
	}

	dutyList := model.RetrieveByDate(dutyDate)
	res["msg"] = "成功"
	res["dutyList"] = dutyList
	c.JSON(200, res)
}

func Duty(c *gin.Context) {
	dutySwitch := c.PostForm("duty")
	if dutySwitch == "0" {
		cornManager.StopPushDuty()
	} else if dutySwitch == "1" {
		cornManager.StartPushDuty()
	}
	c.JSON(200, gin.H{"msg": "success"})
}
