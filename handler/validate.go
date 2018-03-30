package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndphu/token-service/service"
)

func Validate(r *gin.RouterGroup) {
	r.POST("/validate", func(c *gin.Context) {
		body := ValidateBody{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
		} else {
			fmt.Println("validating token: " + body.Token)
			claim, err := service.ValidateToken(body.Token)
			if err != nil {
				c.JSON(500, gin.H{"err": err.Error()})
			} else {
				c.JSON(200, gin.H{"claim": claim})
			}
		}
	})
}

type ValidateBody struct {
	Token string `json:"token"`
}
