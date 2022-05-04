package main

import "github.com/gin-gonic/gin"

func main() {

	app := gin.Default()
	app.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Msg  string
			Code int
		}{"12121212211212", 200})
	})

	app.Run(":7666")
}
