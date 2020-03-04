package router

import (
	v1 "Service/service-server/router/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	apiv1 := r.Group("/api/v1")
	{
		//打开或关闭值班推送
		apiv1.POST("/duty", v1.Duty)
		apiv1.GET("/dutyList", v1.DutyList)
	}
	return r
}
