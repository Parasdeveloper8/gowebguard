package main

import (
	"net/http"

	"github.com/Parasdeveloper8/gowebguard/webguard"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/*headers := &webguard.Headers{
		ContentSecurityPolicy:   "jhhihi",
		StrictTransportSecurity: "gegbhg",
		XContentTypeOptions:     "hhjhu",
		XFrameOptions:           "kjiji",
		XXSSProtection:          "huhlhi",
		ReferrerPolicy:          "bubuihi",
	}*/
	r.Use(webguard.WebGuard())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "Check Headers"})
	})
	r.Run(":4000")
}
