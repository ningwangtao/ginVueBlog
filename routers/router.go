package routers

import (
	"ginVueBlog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	// 1，创建路由
	r := gin.Default()

	// 2，绑定路由规格，执行的函数
	// gin.Context，封装了request和response函数
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	// 3，监听端口，默认8080
	r.Run(utils.HttpPort)
}
