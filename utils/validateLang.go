package utils

import (
	"github.com/jaayroots/go_base/config"
	"github.com/labstack/echo/v4"
)

func ValidateLangOrDefault(pctx echo.Context) string {

	langConfig := config.ConfigGetting().Langs
	lang := pctx.Param("lang")

	isValid := false
	for _, l := range langConfig {

		if lang == l {
			isValid = true
			break
		}
	}
	if !isValid {
		lang = langConfig[0]
	}

	return lang

}

func IsValidLang(lang string) bool {
	langConfig := config.ConfigGetting().Langs

	for _, l := range langConfig {
		if lang == l {
			return true
		}
	}
	return false
}
