/*
 * @Author: holiday
 * @Date: 2019-07-02 14:50:49
 * @Last Modified by: holiday
 * @Last Modified time: 2019-07-02 16:53:08
 */
package server

import (
	"seqsvr/controller"
	"seqsvr/model/short_url"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	controller, err := controller.GetBaseController()
	if err != nil {
		panic(err)
	}
	r.Use(Middleware)
	v1 := r.Group("/v1")
	v1.POST("/create_url", controller.CreateLink)
}

func Middleware(c *gin.Context) {
	urlPath := c.Request.URL.Path
	if strings.Contains(urlPath, "create_url") {
		c.Next()
		return
	} else {
		if tool, err := short_url.GetServer(); err == nil {
			if links, err := tool.GetLinksByKeyword(strings.ReplaceAll(urlPath, "/", "")); err == nil {
				if links != nil {
					c.Redirect(301, links.Url)
					return
				}
			}
		}
		c.Status(404)
		c.Abort()
	}
}
