package template

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteStaticFiles(route *gin.Engine) {

	route.GET("/android-icon-144x144.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesandroidicon144x144png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesandroidicon144x144png)
	})

	route.GET("/android-icon-192x192.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesandroidicon192x192png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesandroidicon192x192png)
	})

	route.GET("/android-icon-36x36.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesandroidicon36x36png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesandroidicon36x36png)
	})

	route.GET("/android-icon-48x48.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesandroidicon48x48png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesandroidicon48x48png)
	})

	route.GET("/android-icon-72x72.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesandroidicon72x72png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesandroidicon72x72png)
	})

	route.GET("/android-icon-96x96.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesandroidicon96x96png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesandroidicon96x96png)
	})

	route.GET("/apple-icon-114x114.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon114x114png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon114x114png)
	})

	route.GET("/apple-icon-120x120.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon120x120png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon120x120png)
	})

	route.GET("/apple-icon-144x144.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon144x144png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon144x144png)
	})

	route.GET("/apple-icon-152x152.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon152x152png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon152x152png)
	})

	route.GET("/apple-icon-180x180.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon180x180png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon180x180png)
	})

	route.GET("/apple-icon-57x57.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon57x57png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon57x57png)
	})

	route.GET("/apple-icon-60x60.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon60x60png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon60x60png)
	})

	route.GET("/apple-icon-72x72.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon72x72png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon72x72png)
	})

	route.GET("/apple-icon-76x76.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleicon76x76png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleicon76x76png)
	})

	route.GET("/apple-icon-precomposed.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleiconprecomposedpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleiconprecomposedpng)
	})

	route.GET("/apple-icon.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesappleiconpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesappleiconpng)
	})

	route.GET("/browserconfig.xml", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesbrowserconfigxml))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/xml; charset=utf-8", Filesbrowserconfigxml)
	})

	route.GET("/favicon-16x16.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesfavicon16x16png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesfavicon16x16png)
	})

	route.GET("/favicon-32x32.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesfavicon32x32png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesfavicon32x32png)
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

	route.GET("/manifest.json", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesmanifestjson))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/json", Filesmanifestjson)
	})

	route.GET("/ms-icon-144x144.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesmsicon144x144png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesmsicon144x144png)
	})

	route.GET("/ms-icon-150x150.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesmsicon150x150png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesmsicon150x150png)
	})

	route.GET("/ms-icon-310x310.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesmsicon310x310png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesmsicon310x310png)
	})

	route.GET("/ms-icon-70x70.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesmsicon70x70png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesmsicon70x70png)
	})

	route.GET("/robots.txt", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesrobotstxt))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/plain; charset=utf-8", Filesrobotstxt)
	})
}
