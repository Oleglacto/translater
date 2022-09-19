package yandex

type Translate struct {
	Texts              []string `json:"texts"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
}

type TranslationResult struct {
	Text string `json:"text"`
}

type TranslationsResult struct {
	Translations []TranslationResult `json:"translations"`
}
