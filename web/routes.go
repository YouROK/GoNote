package web

import (
	"GoNote/web/pages"
)

func (ws *WebServer) SetupRoutesPages() {
	all := ws.r.Group("/")
	all.GET("/", pages.IndexPage)
	all.GET("/note/:noteID", pages.NotePage)
	all.GET("/note/:noteID/edit", pages.EditNotePage)

	all.POST("/new", SpamProtectionMiddleware(), pages.NewNote)
	all.POST("/edit/:noteID", pages.EditNote)
	all.POST("/note/:noteID/checkpass", SpamProtectionMiddleware(), pages.CheckNotePassword)

	all.POST("/report", SpamProtectionMiddleware(), pages.TGBotReport)

	all.GET("/sitemap.xml", pages.Sitemap)
	all.GET("/allnotes", pages.AllNotes)

	ws.r.NoRoute(pages.NotFound)
}
