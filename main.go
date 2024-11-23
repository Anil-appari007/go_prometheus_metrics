package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func GetUsers(c *gin.Context) {
	var users = []user{
		{Name: "u1", Role: "dev"},
		{Name: "u2", Role: "qa"},
	}
	c.JSON(200, gin.H{"users": users})
}

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_total_requests",
		Help: "total number of http requests received",
	},
	[]string{"path"},
)

func prometheusMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		fmt.Println(path)
		totalRequests.WithLabelValues(path).Inc()
		totalRequests.WithLabelValues("all").Inc()
		ctx.Next()
	}
}

func init() {
	prometheus.Register(totalRequests)
}

func main() {
	r := gin.Default()
	r.Use(prometheusMiddleWare())
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "running"})
	})
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//
	r.GET("/users", GetUsers)
	r.Run("localhost:8080")
}
