//计划

//CREATE TABLE `plan` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`planContent` varchar(100) DEFAULT '' COMMENT '标签名称',
//`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
//`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
//`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
//`deleted_on` int(10) unsigned DEFAULT '0',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='计划管理';

package v1

import (
	"MyService/model"
	"github.com/gin-gonic/gin"
)

//新增文章标签
func AddPlan(c *gin.Context) {
	planContent := c.Param("planContent")
	createBy := c.Param("createBy")
	if len(planContent) == 0 {
		c.AbortWithStatusJSON(200, gin.H{"msg": "计划内容不能为空"})
		return
	}
	model.AddPlan(planContent, createBy)
	c.JSON(200, gin.H{"msg": "success"})
}


func GetPlan(c *gin.Context) {
	planId := c.Param("planId")
	if len(planId) == 0 {
		//c.AbortWithStatusJSON(200, gin.H{"msg": "计划内容不能为空"})
	}
	c.JSON(200, gin.H{"msg": "success"})
}

func EditPlan(c *gin.Context) {
	planId := c.Param("planId")
	if len(planId) == 0 {
		c.AbortWithStatusJSON(200, gin.H{"msg": "编辑失败，未找到此条计划"})
	}
	planContent := c.Param("planContent")
	if len(planContent) == 0 {
		c.AbortWithStatusJSON(200, gin.H{"msg": "计划内容不能为空"})
	}
	c.JSON(200, gin.H{"msg": "success"})
}

//新增文章标签
func DeletePlan(c *gin.Context) {
	planId := c.Param("planId")
	if len(planId) == 0 {
		c.AbortWithStatusJSON(200, gin.H{"msg": "编辑失败，未找到此条计划"})
	}
	c.JSON(200, gin.H{"msg": "success"})
}