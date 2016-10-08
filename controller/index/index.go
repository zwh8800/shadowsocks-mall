package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zwh8800/shadowsocks-mall/conf"
	"github.com/zwh8800/shadowsocks-mall/controller"
	"github.com/zwh8800/shadowsocks-mall/render"
	"github.com/zwh8800/shadowsocks-mall/util"
)

func Index(c *gin.Context) {
	c.Render(http.StatusOK, render.NewRender("index.html", gin.H{
		"site":        conf.Conf.Site,
		"social":      conf.Conf.Social,
		"prod":        conf.Conf.Env.Prod,
		"haha":        util.HahaGenarate(),
		"currentPage": controller.CurrentPageInIndex,
	}))
}
