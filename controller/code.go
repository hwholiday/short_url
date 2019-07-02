/*
 * @Author: holiday
 * @Date: 2019-07-02 16:32:59
 * @Last Modified by: holiday
 * @Last Modified time: 2019-07-02 16:34:47
 */
package controller

import "github.com/gin-gonic/gin"

func (c *BaseController) ResponseData(g *gin.Context, data interface{}) {
	g.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "success",
	})
}

func (c *BaseController) ResponseDataFailure(g *gin.Context, data interface{}) {
	g.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "failure",
	})
}
