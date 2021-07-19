package router

import (
	"net/http"
	"sry/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup() *gin.Engine {
	r := gin.Default() //创建一个默认的路由引擎
	r.POST("/login", SignUpHandler)
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "get Book",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post Book",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "put Book",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "delete Book",
		})
	})
	return r
}

//登录或注册功能
func SignUpHandler(c *gin.Context) {
	//1，获取用户输入的数据，参数校验
	p := new(service.UserSignUpInfo)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param: ", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	zap.L().Info("Sign up success")

	//2，业务处理 将用户输入的数据保存进数据库
	//3，返回响应
}
