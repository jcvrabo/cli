package i18n

import "code.cloudfoundry.org/cli/v9/util/ui"

var T ui.TranslateFunc

type LocaleReader interface {
	Locale() string
}

func Init(config LocaleReader) ui.TranslateFunc {
	t, _ := ui.GetTranslationFunc(config)
	return t
}
