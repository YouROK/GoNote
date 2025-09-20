package web

import (
	"GoNote/config"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateEntry struct {
	timestamps []time.Time
}

var (
	rateRecords = make(map[string]*rateEntry)
	mu          sync.Mutex
)

func init() {
	go cleaner()
}

func cleaner() {
	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		now := time.Now()
		mu.Lock()
		for k, e := range rateRecords {
			if len(e.timestamps) == 0 || now.Sub(e.timestamps[len(e.timestamps)-1]) > time.Duration(config.Cfg.Antispam.WindowSec)*time.Second {
				delete(rateRecords, k)
			}
		}
		mu.Unlock()
	}
}

func SpamProtectionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP() + "|" + c.FullPath()
		now := time.Now()

		mu.Lock()
		e, exists := rateRecords[key]
		if !exists {
			e = &rateEntry{
				timestamps: []time.Time{now},
			}
			rateRecords[key] = e
			mu.Unlock()
			c.Next()
			return
		}

		// Убираем старые записи за окно
		var newTimestamps []time.Time
		for _, t := range e.timestamps {
			if now.Sub(t) <= time.Duration(config.Cfg.Antispam.WindowSec)*time.Second {
				newTimestamps = append(newTimestamps, t)
			}
		}
		e.timestamps = newTimestamps

		if len(e.timestamps) >= config.Cfg.Antispam.MaxRequests {
			mu.Unlock()
			// превышен лимит
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, please try later",
			})
			c.Abort()
			return
		}

		// добавляем текущий запрос
		e.timestamps = append(e.timestamps, now)
		mu.Unlock()

		c.Next()
	}
}
