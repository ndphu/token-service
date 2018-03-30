package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/token-service/domain"
	"github.com/ndphu/token-service/service"
)

func Create(r *gin.RouterGroup) {
	r.POST("/create", func(c *gin.Context) {
		input := domain.Token{}
		err := c.ShouldBindJSON(&input)
		if err != nil {
			c.JSON(500, gin.H{"err": err.Error()})
			panic(err)
		} else {
			token, err := service.CreateToken(&input)
			if err != nil {
				c.JSON(500, gin.H{"err": err.Error()})
			} else {
				c.JSON(200, gin.H{"token": token})
			}
		}
	})
}
