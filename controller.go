package main

import (
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	requestParameter := c.Query("n")
	i, err := strconv.Atoi(requestParameter)
	if err != nil || i<=0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"message" : "Please set positive number in request parameter",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": calcFibonacciNum(i),
	})
}

func noroot (c *gin.Context){
	c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"message" : "no route",
	})
}

func calcFibonacciNum(n int) *big.Int {
	switch {
	case n == 1 || n == 2:
		log.Print("1")
		return big.NewInt(1)
	case n >= 3:
		var f0, f1, f2 *big.Int
		f0 = big.NewInt(1)
		f1 = big.NewInt(1)
		f2 = big.NewInt(0)
		for i := 3; i <= n; i++ {
			f2.Add(f0, f1)
			f0 = f1
			f1 = f2
			f2 = big.NewInt(0)
		}
		return f1
	}
	return big.NewInt(0)
}
