package router

import (
	"MyService/middleware/jwt"
	"MyService/router/api"
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

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/users", v1.Test)
		//新建标签
		apiv1.POST("/users", v1.AddUser)
		//更新指定标签
		apiv1.PUT("/users/:id", v1.EditUser)
		//删除指定标签
		apiv1.DELETE("/users/:id", v1.DeleteUser)
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
