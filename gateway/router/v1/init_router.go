/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 6:40 下午
 **/
package v1

import (
	. "gaea/gateway/middleware"
	"gaea/gateway/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router     *gin.Engine
	api_prefix = "/api/v1"
)

func Initrouter() *gin.Engine {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	//gin.New()返回一个*Engine 指针
	//而gin.Default()不但返回一个*Engine 指针，而且还进行了debugPrintWARNINGDefault()和engine.Use(Logger(), Recovery())其他的一些中间件操作
	router = gin.Default()
	//router = gin.New()

	//使用日志
	//router.Use(gin.Logger())
	//使用Panic处理方案
	//router.Use(gin.Recovery())

	router.Use(InitErrorHandler)
	router.Use(InitAccessLogMiddleware)

	// 未知调用方式
	router.NoMethod(InitNoMethodJson)
	// 未知路由处理
	router.NoRoute(InitNoRouteJson)

	app := router.Group(api_prefix + "/app")
	{
		app.GET("", service.AppList)
		app.POST("", service.AppAdd)
		app.PUT("", service.AppUpdate)
		app.DELETE("", service.AppDel)
	}
	router.POST("/pp", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "post",
		})
	})

	return router

}
