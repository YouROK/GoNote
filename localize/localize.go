package localize

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/locale.*.json
var LocaleFS embed.FS

var bundle *i18n.Bundle

func Init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	ff, err := LocaleFS.ReadDir("locales")
	if err != nil {
		log.Fatalln("Error reading locales:", err)
	}
	for _, entry := range ff {
		if !entry.IsDir() {
			_, err := bundle.LoadMessageFileFS(LocaleFS, "locales/"+entry.Name())
			if err != nil {
				log.Println("Error loading locales:", err)
			}
		}
	}
}

func LocalizerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accept := c.GetHeader("Accept-Language")
		matcher := language.NewMatcher(bundle.LanguageTags())

		tags, _, _ := language.ParseAcceptLanguage(accept)
		tag, _, _ := matcher.Match(tags...)

		localizer := i18n.NewLocalizer(bundle, tag.String())

		c.Set("localizer", localizer)
		c.Set("lang", tag.String()[:2])
		c.Next()
	}
}

func T(c *gin.Context, id string) string {
	l, ok := c.MustGet("localizer").(*i18n.Localizer)
	if !ok {
		return id
	}
	msg, err := l.Localize(&i18n.LocalizeConfig{MessageID: id})
	if err != nil {
		return id
	}
	return msg
}
