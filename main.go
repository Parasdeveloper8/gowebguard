package main

import (
	"net/http"

	"github.com/Parasdeveloper8/gowebguard/webguard"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/*headers := &webguard.Headers{
		StrictTransportSecurity: "gegbhg",
		XContentTypeOptions:     "hhjhu",
		XFrameOptions:           "kjiji",
		XXSSProtection:          "huhlhi",
		ReferrerPolicy:          "bubuihi",
	}*/
	//headers := &webguard.Headers{ContentSecurityPolicy: "kop"}
	r.Use(webguard.WebGuard())
	header := &webguard.Csp{
		ContentSecurityPolicy: "style-src:'self'",
	}
	r.Use(webguard.CSP(header, header))
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "Check Headers"})
	})
	r.Run(":4000")
}
