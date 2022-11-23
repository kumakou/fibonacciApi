package main

import "github.com/gin-gonic/gin"

func router() *gin.Engine {
	router := gin.Default()
	router.GET("/fib", getting)
	router.NoRoute(noroot)
	
	return router
}