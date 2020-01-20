package api

import (
	"MyService/model"
	"MyService/util"
	"github.com/astaxie/beego/validation"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := util.InvalidParams
	if ok {
		isExist := model.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = util.ErrorAuthToken
			} else {
				data["token"] = token

				code = util.Success
			}

		} else {
			code = util.ErrorAuth
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : util.GetResponseMsg(code),
		"data" : data,
	})
}