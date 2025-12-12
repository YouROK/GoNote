package static

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"mime"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed files/*
var staticFS embed.FS

//go:embed templates/*
var templateFS embed.FS

var htmlTmpl *template.Template

func init() {
	err := mime.AddExtensionType(".webmanifest", "application/manifest+json")
	if err != nil {
		log.Println("Error set mime type:", err)
	}

	subTmplFS, err := fs.Sub(templateFS, "templates")
	if err != nil {
		panic(err)
	}
	htmlTmpl = template.Must(template.ParseFS(subTmplFS, "*.go.html"))
}

func RouteEmbedFiles(route *gin.Engine) {
	route.SetHTMLTemplate(htmlTmpl)

	subFS, err := fs.Sub(staticFS, "files")
	if err != nil {
		panic(err)
	}
	route.StaticFS("/st", http.FS(subFS))
}
