package translation

import (
	"githab.com/oleglacto/translater/internal/pkg/yandex"
	"strings"
	"unicode"
)

type YandexClient interface {
	Do(data yandex.Translate) (*yandex.TranslationsResult, error)
}

type UseCase struct {
	translator YandexClient
}

func NewUseCase(translator YandexClient) *UseCase {
	return &UseCase{translator: translator}
}

func (u UseCase) Translate(text string) (string, error) {
	text = strings.TrimSpace(text)
	runes := []rune(text)

	translate := yandex.Translate{Texts: []string{text}, TargetLanguageCode: "ru"}

	if unicode.Is(unicode.Cyrillic, runes[0]) {
		translate.TargetLanguageCode = "en"
	}

	result, err := u.translator.Do(translate)
	if err != nil {
		return "", err
	}

	return result.Translations[0].Text, nil
}
