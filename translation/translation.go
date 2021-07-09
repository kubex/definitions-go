package translation

type Text struct {
	Fallback     string            `json:"fallback"`
	Translations map[string]string `json:"translations"`
}

func (t *Text) Get(language string) string {
	if txt, ok := t.Translations[language]; ok {
		return txt
	}
	return t.Fallback
}

func String(text string) Text {
	return Text{Fallback: text}
}
