package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndphu/token-service/handler"
)

func main() {
	r := gin.Default()
	g := r.Group("/token-service")
	handler.Create(g)
	handler.Validate(g)
	r.Run(":8081")
}
