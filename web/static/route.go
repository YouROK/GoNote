package template

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteStaticFiles(route *gin.Engine) {

	route.GET("/apple-touch-icon.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappletouchiconpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappletouchiconpng)
	})

	route.GET("/banner.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesbannerpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesbannerpng)
	})

	route.GET("/css/gonote.css", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filescssgonotecss))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/css; charset=utf-8", Filescssgonotecss)
	})

	route.GET("/favicon-96x96.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesfavicon96x96png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesfavicon96x96png)
	})

	route.GET("/favicon.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesfaviconico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesfaviconico)
	})

	route.GET("/favicon.svg", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesfaviconsvg))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/svg+xml", Filesfaviconsvg)
	})

	route.GET("/js/editor.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjseditorjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjseditorjs)
	})

	route.GET("/js/shared.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjssharedjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjssharedjs)
	})

	route.GET("/js/viewer.js", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesjsviewerjs))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/javascript; charset=utf-8", Filesjsviewerjs)
	})

	route.GET("/robots.txt", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesrobotstxt))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/plain; charset=utf-8", Filesrobotstxt)
	})

	route.GET("/site.webmanifest", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filessitewebmanifest))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/manifest+json", Filessitewebmanifest)
	})

	route.GET("/web-app-manifest-192x192.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Fileswebappmanifest192x192png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Fileswebappmanifest192x192png)
	})

	route.GET("/web-app-manifest-512x512.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Fileswebappmanifest512x512png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Fileswebappmanifest512x512png)
	})
}
