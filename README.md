# gowebguard

## Introducing GoWebGuard

This package is very helpful for you if you want to secure your web application.
To use this package , follow given instructions.
Use this **command** to get started :-

```go
    go get github.com/Parasdeveloper8/gowebguard/v2@v2.0.2
```
After installing ,use like this :-

```go
    r := gin.Default()
    r.Use(webguard.WebGuard())
```

It's CSP middleware
```go
   r.Use(webguard.CSP())
  //This will add value to CSP header.
```

Now everything will work fine but if you wanna customize configurations,create a instance of webguard.Headers like this
```go
   r := gin.Default()
   headers := &webguard.Headers{
		StrictTransportSecurity: "your value",
		XContentTypeOptions:     "your value",
		XFrameOptions:           "your value",
		XXSSProtection:          "your value",
		ReferrerPolicy:          "your value",
	}
 cheader := &webguard.Csp{
		ContentSecurityPolicy: "your value",
	}//you can add your value to CSP header like this

 r.Use(webguard.WebGuard(headers))
 r.Use(webguard.CSP(cheader))
```

This package secures web application by setting headers like **ContentSecurityPolicy**,**StrictTransportSecurity**,**XContentTypeOptions**,**XFrameOptions**.


**This package will be improved continuously**.

Creator := *paras prajapat*
  
