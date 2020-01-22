//计划

package v1

import (
	"MyService/model"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//新增计划
func AddPlan(c *gin.Context) {
	planContent := c.PostForm("planContent")
	createBy := c.Param("createBy")
	if len(planContent) == 0 {
		c.JSON(200, gin.H{"msg": "计划内容不能为空"})
		return
	}
	model.AddPlan(planContent, createBy)
	c.JSON(200, gin.H{"msg": "success"})
}

//删除计划
func DeletePlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 0 {
		c.JSON(200, gin.H{"msg": "删除失败，未找到此条计划"})
		return
	}
	if !model.ExistPlanById(id) {
		fmt.Println("000")
		c.JSON(200, gin.H{"msg": "删除失败，未找到此条计划"})
		return
	}
	model.DeletePlan(id)
	c.JSON(200, gin.H{"msg": "success"})
}

//修改计划
func EditPlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 0 {
		c.JSON(200, gin.H{"msg": "编辑失败，未找到此条计划"})
		return
	}
	if !model.ExistPlanById(id) {
		fmt.Println("000")
		c.JSON(200, gin.H{"msg": "编辑失败，未找到此条计划"})
		return
	}
	content := c.PostForm("planContent")
	if len(content) == 0 {
		c.JSON(200, gin.H{"msg": "编辑失败，计划内容不能为空"})
	}
	model.EditPlan(id, content)
	c.JSON(200, gin.H{"msg": "success"})
}

//查询计划
func GetPlan(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("id"))
	count, _ := strconv.Atoi(c.Query("id"))
	fmt.Printf("page: %d count: %d", page, count)
	plans := model.GetPlan()
	c.JSON(200, gin.H{"msg": "success", "planList": plans})
}

func GetOnePlan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id < 0 {
		c.JSON(200, gin.H{"msg": "未找到此条计划"})
		return
	}
	plan := model.GetOnePlan(id)
	c.JSON(200, gin.H{"msg": "success", "plan": plan})
}
