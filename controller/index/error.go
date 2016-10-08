package index

import (
	"github.com/gin-gonic/gin"

	"github.com/zwh8800/shadowsocks-mall/conf"
	"github.com/zwh8800/shadowsocks-mall/render"
)

func ErrorHandler(c *gin.Context, status int, err error) {
	c.Render(status, render.NewRender("error.html", gin.H{
		"err":  err,
		"site": conf.Conf.Site,
		"prod": conf.Conf.Env.Prod,
	}))
}
