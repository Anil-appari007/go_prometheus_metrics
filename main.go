package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "running"})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run("localhost:8080")
}
