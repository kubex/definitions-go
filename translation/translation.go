package translation

type Text struct {
	Fallback     string
	Translations map[string]string
}

func (t *Text) Get(language string) string {
	if txt, ok := t.Translations[language]; ok {
		return txt
	}
	return t.Fallback
}
