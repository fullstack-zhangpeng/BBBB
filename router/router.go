package router

import (
	v1 "MyService/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	//r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//新增计划
		apiv1.POST("/plan", v1.AddPlan)
		//删除计划
		apiv1.DELETE("/plan/:id", v1.DeletePlan)
		//修改计划
		apiv1.PUT("/plan/:id", v1.EditPlan)
		//获取计划列表
		apiv1.GET("/plan", v1.GetPlan)
		//获取指定 id 的计划
		apiv1.GET("/plan/:id", v1.GetOnePlan)
	}
	return r
}

//func InitRouter() *gin.Engine {
//	r := gin.New()
//
//	r.Use(gin.Logger())
//
//	r.Use(gin.Recovery())
//
//	gin.SetMode(setting.RunMode)
//

//	apiv1 := r.Group("/api/v1")
//	{
//		//获取标签列表
//		apiv1.GET("/tags", v1.GetTags)
//		//新建标签
//		apiv1.POST("/tags", v1.AddTag)
//		//更新指定标签
//		apiv1.PUT("/tags/:id", v1.EditTag)
//		//删除指定标签
//		apiv1.DELETE("/tags/:id", v1.DeleteTag)
//	}
//	return r
//}
