package webguard

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Headers value struct
// We can add custom value in headers
type Headers struct {
	StrictTransportSecurity string
	XContentTypeOptions     string
	XFrameOptions           string
	XXSSProtection          string
	ReferrerPolicy          string
}

// CSP:content security policy struct
// We will use Csp struct to use user's csp header's value
type Csp struct {
	ContentSecurityPolicy string
}

//This middleware secures webapp by securing headers.
//I am using variadic parameters to  pass arguments in middleware or not
/*
It's completely user's choice to add their values to headers otherwise
 it's not necessary to pass struct as a value in middleware.
 Middleware will use default values.
*/

func WebGuard(H ...*Headers) gin.HandlerFunc {
	var Header *Headers
	//iterating over H to use first given struct's values
	for i := 0; i < len(H); i++ {
		Header = H[0]
	}
	if len(H) > 1 {
		log.Fatal("Can accept only one argument in middleware")
	}
	return func(c *gin.Context) {
		//check if argument is passed or not
		if H == nil {
			//Start securing headers
			//set Strict-Transport-Security header
			c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

			//set X-Content-Type-Options header
			c.Header("X-Content-Type-Options", "nosniff")

			//set X-Frame-Options header
			c.Header("X-Frame-Options", "DENY")

			//set X-XSS-Protection header
			c.Header("X-XSS-Protection", "1; mode=block")

			//set Referrer-Policy header
			c.Header("Referrer-Policy", "no-referrer-when-downgrade")
		} else {
			//we will use given values in header which are passed in parameters

			//set Strict-Transport-Security header
			c.Header("Strict-Transport-Security", Header.StrictTransportSecurity)

			//set X-Content-Type-Options header
			c.Header("X-Content-Type-Options", Header.XContentTypeOptions)

			//set X-Frame-Options header
			c.Header("X-Frame-Options", Header.XFrameOptions)

			//set X-XSS-Protection header
			c.Header("X-XSS-Protection", Header.XXSSProtection)

			//set Referrer-Policy header
			c.Header("Referrer-Policy", Header.ReferrerPolicy)
		}
		//if every thing is Ok call Next()
		c.Next()
	}
}

// CSP webguard middleware
func CSP(C ...*Csp) gin.HandlerFunc {
	var CSP *Csp
	//iterating over C to use first given struct's values
	for i := 0; i < len(C); i++ {
		CSP = C[0]
	}
	if len(C) > 1 {
		log.Fatal("Can accept only one argument in middleware")
	}
	return func(c *gin.Context) {
		//check if argument is passed or not
		if C == nil {
			//set Content-Security-Policy header
			c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline' 'unsafe-eval'; object-src 'none'; style-src 'self' 'unsafe-inline'; img-src 'self'; connect-src 'self'; font-src 'self'; frame-src 'none';")
		} else {
			//we will use given value in header which is passed in parameters
			//set Content-Security-Policy header
			c.Header("Content-Security-Policy", CSP.ContentSecurityPolicy)
		}
	}
}
