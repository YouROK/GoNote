package web

import (
	"GoNote/config"
	"GoNote/localize"
	"GoNote/static"
	"GoNote/storage"
	"GoNote/tgbot"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WebServer struct {
	store storage.Store
	r     *gin.Engine
}

func NewServer() *WebServer {
	var err error
	var store storage.Store
	switch config.Cfg.DB.Type {
	case "fs":
		store, err = storage.NewStore(storage.FsStore, "db")
	case "bbolt":
		store, err = storage.NewStore(storage.BoltdbStore, "db")
	case "sqlite":
		store, err = storage.NewStore(storage.SqliteStore, "db")
	default:
		store, err = storage.NewStore(storage.FsStore, "db")
	}

	if err != nil {
		log.Fatal("Error create db:", err)
	}

	go store.RemoveExpiredSessions()
	tgbot.SetStore(store)

	return &WebServer{store: store}
}

func (ws *WebServer) Run() {
	ws.r = gin.Default()

	ws.r.Use(ws.SessionMiddleware(), localize.LocalizerMiddleware())
	static.RouteEmbedFiles(ws.r)
	ws.SetupRoutesPages()
	err := ws.r.Run(config.Cfg.Server.Host + ":" + strconv.Itoa(config.Cfg.Server.Port))
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
