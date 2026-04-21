package main

import "github.com/gin-gonic/gin"

// import "fmt"

func main() {
	// fmt.Println("🚀 Todo API Server")
	// fmt.Println("✅ Project setup complete!")
	// fmt.Println("📁 Folder structure created")

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Todo Api is running",
			"status": "success",
		})
	})

	router.Run(":3000")

}