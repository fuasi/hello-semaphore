package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"msg":  "hel11lo",
		})
	})

	err := engine.Run(":7666")
	if err != nil {
		log.Println(err.Error())
		return
	}
}
