package sitemap

import (
	"net/http"

	"github.com/chonglou/sitemap"
	"github.com/gin-gonic/gin"
)

func SiteMap(c *gin.Context) {
	s := sitemap.New()

	c.XML(http.StatusOK, s)
}
