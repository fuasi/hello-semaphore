package main

func main() {

	app := gin.Default()
	app.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Msg  string
			Code int
		}{"LSCB", 200})
	})

	app.Run(":7666")
}
