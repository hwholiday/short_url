/*
 * @Author: holiday
 * @Date: 2019-07-02 15:03:57
 * @Last Modified by: holiday
 * @Last Modified time: 2019-07-02 16:46:41
 */
package controller

import (
	"fmt"
	"short_url/base"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (b *BaseController) CreateLink(g *gin.Context) {
	var (
		url     string
		err     error
		order   string
		status  string
		keyword string
	)
	url = g.PostForm("url")
	status = g.PostForm("status")
	keyword = g.PostForm("keyword")
	if !strings.Contains(url, "http") {
		url = "http://" + url
	}
	statusInt, err := strconv.Atoi(status)
	if err != nil {
		b.ResponseDataFailure(g, err.Error())
		return
	}
	if order, err = b.shortUrl.CreateLinks(url, keyword, statusInt); err != nil {
		b.ResponseDataFailure(g, err.Error())
		return
	}
	b.ResponseData(g, fmt.Sprintf("http://%s/%s", base.ServerUrl, order))
}
