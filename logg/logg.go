package logg

import (
	"io"
	"log"
	"os"
	"strings"

	"GoNote/config"
)

type levelWriter struct {
	writer io.Writer
	level  int
}

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var levelMap = map[string]int{
	"debug": DEBUG,
	"info":  INFO,
	"warn":  WARN,
	"error": ERROR,
}

func (lw levelWriter) Write(p []byte) (n int, err error) {
	// проверяем префикс уровня сообщения
	msg := string(p)
	if strings.HasPrefix(msg, "[DEBUG]") && lw.level > DEBUG {
		return len(p), nil
	}
	if strings.HasPrefix(msg, "[INFO]") && lw.level > INFO {
		return len(p), nil
	}
	if strings.HasPrefix(msg, "[WARN]") && lw.level > WARN {
		return len(p), nil
	}
	// ERROR всегда пишем
	return lw.writer.Write(p)
}

func Init() {
	var out io.Writer
	if config.Cfg.Logging.File != "" {
		f, err := os.OpenFile(config.Cfg.Logging.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("не удалось открыть файл логов: %v", err)
		}
		out = f
	} else {
		out = os.Stdout
	}

	// определяем уровень из конфига
	lvl, ok := levelMap[strings.ToLower(config.Cfg.Logging.Level)]
	if !ok {
		lvl = INFO
	}

	// переопределяем стандартный логгер
	log.SetOutput(levelWriter{writer: out, level: lvl})
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("[INFO] ")
}

// Функции для удобства (можно использовать вместо логического префикса)
func Debug(v ...interface{}) { log.Println(append([]interface{}{"[DEBUG]"}, v...)...) }
func Info(v ...interface{})  { log.Println(append([]interface{}{"[INFO]"}, v...)...) }
func Warn(v ...interface{})  { log.Println(append([]interface{}{"[WARN]"}, v...)...) }
func Error(v ...interface{}) { log.Println(append([]interface{}{"[ERROR]"}, v...)...) }
