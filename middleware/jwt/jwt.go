package jwt

import (
	"MyService/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = util.Success
		token := c.Query("token")
		if token == "" {
			code = util.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = util.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = util.ErrorAuthCheckTokenTimeout
			}
		}

		if code != util.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : util.GetResponseMsg(code),
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}