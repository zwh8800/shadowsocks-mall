package route

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/zwh8800/shadowsocks-mall/controller/console"
	"github.com/zwh8800/shadowsocks-mall/controller/index"
	"github.com/zwh8800/shadowsocks-mall/controller/issue"
	"github.com/zwh8800/shadowsocks-mall/controller/sitemap"
)

func Route(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		index.ErrorHandler(c, 404, errors.New("404 Not Found"))
	})

	indexGroup := r.Group("/")
	{
		indexGroup.GET("/", index.Index)
	}

	consoleGroup := r.Group("console")
	{
		consoleGroup.GET("/index", console.Index)
	}

	issueGroup := r.Group("issue")
	{
		issueGroup.GET("/new", issue.New)
	}

	r.GET("/sitemap.xml", sitemap.SiteMap)
	r.Static("/static", "./static")
}
