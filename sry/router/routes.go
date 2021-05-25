package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.Default() //创建一个默认的路由引擎
	r.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "get Book",
		})
	})
	r.POST("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "post Book",
		})
	})
	r.PUT("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "put Book",
		})
	})
	r.DELETE("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "delete Book",
		})
	})
	return r
}