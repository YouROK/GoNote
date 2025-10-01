package web

import (
	"GoNote/config"
	"GoNote/localize"
	"GoNote/storage"
	template "GoNote/web/static"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	store storage.Store
	r     *gin.Engine
}

func NewServer() *WebServer {
	store, err := storage.NewStore(storage.FS_STORE, "db")
	if err != nil {
		log.Fatal("Error create db:", err)
	}
	return &WebServer{store: store}
}

func (ws *WebServer) Run() {
	ws.r = gin.Default()

	ws.r.Use(ws.SessionMiddleware(), localize.LocalizerMiddleware())
	ws.r.LoadHTMLGlob("web/temp/*")
	template.RouteStaticFiles(ws.r)
	ws.SetupRoutesPages()
	ws.r.Run(config.Cfg.Server.Host + ":" + strconv.Itoa(config.Cfg.Server.Port))
}
