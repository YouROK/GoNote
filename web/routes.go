package web

import (
	"GoNote/web/pages"
)

func (ws *WebServer) SetupRoutesPages() {
	all := ws.r.Group("/")
	all.GET("/", pages.HandleIndex)
	all.GET("/note/:noteID", pages.HandleNote)
	all.POST("/note/:noteID/checkpass", pages.CheckNotePassword)
	all.GET("/note/:noteID/edit", pages.EditNote)
	all.POST("/note/:noteID/pub", pages.PublishNote)

	all.GET("/sitemap.xml", pages.Sitemap)
}
