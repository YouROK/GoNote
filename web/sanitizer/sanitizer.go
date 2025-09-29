package sanitizer

import (
	"regexp"
	"strings"
)

func SanitizeStyleAttrs(html string, elem string) string {
	re := regexp.MustCompile(`(?i)<` + elem + `[^>]*style="([^"]*)"`)
	return re.ReplaceAllStringFunc(html, func(m string) string {
		sub := re.FindStringSubmatch(m)
		if len(sub) < 2 {
			return m
		}
		safe := SanitizeStyle(sub[1])
		if safe == "" {
			// удаляем атрибут style целиком
			return strings.Replace(m, ` style="`+sub[1]+`"`, "", 1)
		}
		return strings.Replace(m, sub[1], safe, 1)
	})
}

// allowedCSS — whitelist свойств, которые мы разрешаем в inline style KaTeX
var allowedCSS = map[string]bool{
	"height":         true,
	"width":          true,
	"min-width":      true,
	"top":            true,
	"left":           true,
	"margin-right":   true,
	"margin-left":    true,
	"vertical-align": true,
	"display":        true,
	"position":       true, // но значения только relative|absolute
	"transform":      true, // только простые translate(...)
}

// допустимые значения для position
var allowedPositionValues = map[string]bool{
	"relative": true,
	"absolute": true,
}

// регексы для проверки значений
var (
	reNumberUnit         = regexp.MustCompile(`^-?\d+(\.\d+)?(em|px|rem|%)$`)                                // 0.2778em, -3.063em, 12px, 50%
	reNumber             = regexp.MustCompile(`^-?\d+(\.\d+)?$`)                                             // 0 or 0.5 (без единиц)
	reTransformTranslate = regexp.MustCompile(`^\s*translate(X|Y)?\(\s*-?\d+(\.\d+)?(em|px|rem|%)\s*\)\s*$`) // translate(0.1em) or translateX(-2px)
	reWordValue          = regexp.MustCompile(`^[a-zA-Z\-]+$`)                                               // inline, inline-block, middle, baseline и т.п.
)

// SanitizeStyle — фильтрует строку inline style по whitelist.
// Возвращает строку с безопасными парами key: value разделёнными точкой с запятой.
// Если ничего не осталось — возвращает пустую строку.
func SanitizeStyle(style string) string {
	parts := strings.Split(style, ";")
	var safeParts []string

	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		kv := strings.SplitN(p, ":", 2)
		if len(kv) != 2 {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(kv[0]))
		val := strings.TrimSpace(kv[1])
		lval := strings.ToLower(val)

		// запрещаем явные опасные конструкции
		if strings.Contains(lval, "url(") || strings.Contains(lval, "expression(") || strings.Contains(lval, "javascript:") || strings.Contains(lval, "behavior:") || strings.Contains(lval, "data:") && !strings.HasPrefix(lval, "data:") {
			continue
		}

		// проверяем, разрешено ли свойство
		if !allowedCSS[key] {
			continue
		}

		// валидация значений в зависимости от ключа
		switch key {
		case "position":
			// допускаем только relative или absolute
			if allowedPositionValues[strings.TrimSpace(lval)] {
				safeParts = append(safeParts, key+": "+val)
			}
		case "display", "vertical-align":
			// чаще всего словесные значения (inline, inline-block, middle, baseline),
			// допускаем либо слово, либо число/единицу (маловероятно, но безопасно)
			if reWordValue.MatchString(lval) || reNumberUnit.MatchString(lval) || reNumber.MatchString(lval) {
				safeParts = append(safeParts, key+": "+val)
			}
		case "transform":
			// очень ограниченно — только translate/translateX/translateY с единицами
			if reTransformTranslate.MatchString(val) {
				safeParts = append(safeParts, key+": "+val)
			}
		default:
			// для остальных — только числа с единицами или простое число
			if reNumberUnit.MatchString(lval) || reNumber.MatchString(lval) {
				safeParts = append(safeParts, key+": "+val)
			}
		}
	}

	// соединяем обратно
	if len(safeParts) == 0 {
		return ""
	}
	return strings.Join(safeParts, "; ")
}
