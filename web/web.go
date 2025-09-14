package web

import (
	"GoNote/config"
	"GoNote/storage/fstorage"
	template "GoNote/web/static"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	store *fstorage.FileStore
	r     *gin.Engine
}

func NewServer() *WebServer {
	gin.SetMode(gin.ReleaseMode)

	store := fstorage.NewFileStore("db")
	return &WebServer{store: store}
}

func (ws *WebServer) Run() {
	ws.r = gin.Default()

	ws.r.Use(ws.SessionMiddleware())
	ws.r.LoadHTMLGlob("web/temp/*")
	template.RouteStaticFiles(ws.r)
	ws.SetupRoutesPages()
	ws.r.Run(config.Cfg.Server.Host + ":" + strconv.Itoa(config.Cfg.Server.Port))
}
