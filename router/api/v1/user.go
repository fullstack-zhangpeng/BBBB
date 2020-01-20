//用户
package v1

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "success"})
}

//新增文章标签
func AddUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "success"})
}

func EditUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "success"})
}

//新增文章标签
func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "success"})
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "success"})
}