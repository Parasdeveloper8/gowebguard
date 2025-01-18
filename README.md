# gowebguard

## Introducing GoWebGuard

This package is very helpful for you if you want to secure your web application.
To use this package , follow given instructions.
Use this **command** to get started :-

```go
    go get github.com/Parasdeveloper8/gowebguard@v1.0.0
```
After installing ,use like this :-

```go
    r := gin.Default()
    r.Use(webguard.WebGuard())
```
Now everything will work fine but if you wanna customize configurations,create a instance of webguard.Headers like this
```go
   r := gin.Default()
   headers := &webguard.Headers{
		ContentSecurityPolicy:   "your value",
		StrictTransportSecurity: "your value",
		XContentTypeOptions:     "your value",
		XFrameOptions:           "your value",
		XXSSProtection:          "your value",
		ReferrerPolicy:          "your value",
	}
 r.Use(webguard.WebGuard(headers))
```
This package secures web application by setting headers like **ContentSecurityPolicy**,**StrictTransportSecurity**,**XContentTypeOptions**,**XFrameOptions**.
**This package will be improved continuously**.
Creator := *paras prajapat*
  
