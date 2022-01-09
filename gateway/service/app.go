/**
 * @Author xuzhipeng
 * @Description
 * @Date 2022/1/9 7:12 下午
 **/
package service

import (
	"gaea/gateway/tools/app"
	"github.com/gin-gonic/gin"
)

func AppAdd(c *gin.Context) {

}

func AppDel(c *gin.Context) {

}

func AppList(c *gin.Context) {
	app.OK(c, "data", "获取app列表成功")
}

func AppUpdate(c *gin.Context) {

}
