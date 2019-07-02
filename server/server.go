/*
 * @Author: holiday
 * @Date: 2019-07-02 14:53:42
 * @Last Modified by: holiday
 * @Last Modified time: 2019-07-02 15:02:32
 */

package server

import (
	"net/http"
	"short_url/base"
	"short_url/base/tool"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(addr string) {
	base.ServerUrl = addr
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	Route(g)
	s := &http.Server{
		Handler:        g,
		Addr:           addr,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	tool.GetLogger().Info("server start success : " + addr)
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
