package template

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouteStaticFiles(route *gin.Engine) {

	route.GET("/android-icon-144x144.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidicon144x144png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidicon144x144png)
	})

	route.GET("/android-icon-192x192.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidicon192x192png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidicon192x192png)
	})

	route.GET("/android-icon-36x36.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidicon36x36png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidicon36x36png)
	})

	route.GET("/android-icon-48x48.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidicon48x48png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidicon48x48png)
	})

	route.GET("/android-icon-72x72.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidicon72x72png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidicon72x72png)
	})

	route.GET("/android-icon-96x96.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgandroidicon96x96png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgandroidicon96x96png)
	})

	route.GET("/apple-icon-114x114.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon114x114png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon114x114png)
	})

	route.GET("/apple-icon-120x120.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon120x120png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon120x120png)
	})

	route.GET("/apple-icon-144x144.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon144x144png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon144x144png)
	})

	route.GET("/apple-icon-152x152.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon152x152png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon152x152png)
	})

	route.GET("/apple-icon-180x180.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon180x180png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon180x180png)
	})

	route.GET("/apple-icon-57x57.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon57x57png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon57x57png)
	})

	route.GET("/apple-icon-60x60.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon60x60png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon60x60png)
	})

	route.GET("/apple-icon-72x72.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon72x72png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon72x72png)
	})

	route.GET("/apple-icon-76x76.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleicon76x76png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleicon76x76png)
	})

	route.GET("/apple-icon-precomposed.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleiconprecomposedpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleiconprecomposedpng)
	})

	route.GET("/apple-icon.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgappleiconpng))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgappleiconpng)
	})

	route.GET("/browserconfig.xml", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgbrowserconfigxml))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "text/xml; charset=utf-8", Filesimgbrowserconfigxml)
	})

	route.GET("/favicon-16x16.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfavicon16x16png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgfavicon16x16png)
	})

	route.GET("/favicon-32x32.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfavicon32x32png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgfavicon32x32png)
	})

	route.GET("/favicon-96x96.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfavicon96x96png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgfavicon96x96png)
	})

	route.GET("/favicon.ico", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgfaviconico))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/vnd.microsoft.icon", Filesimgfaviconico)
	})

	route.GET("/manifest.json", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmanifestjson))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "application/json", Filesimgmanifestjson)
	})

	route.GET("/ms-icon-144x144.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmsicon144x144png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgmsicon144x144png)
	})

	route.GET("/ms-icon-150x150.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmsicon150x150png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgmsicon150x150png)
	})

	route.GET("/ms-icon-310x310.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmsicon310x310png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgmsicon310x310png)
	})

	route.GET("/ms-icon-70x70.png", func(c *gin.Context) {
		etag := fmt.Sprintf("%x", md5.Sum(Filesimgmsicon70x70png))
		c.Header("Cache-Control", "public, max-age=31536000")
		c.Header("ETag", etag)
		c.Data(200, "image/png", Filesimgmsicon70x70png)
	})
}
