package translation

import "testing"

func TestFromMap(t *testing.T) {
	m := map[string]string{
		"EN": "English",
		"FR": "French",
	}

	enDefault := FromMap(m, "EN")
	if enDefault.Fallback != "English" {
		t.Error("incorrect fallback for English")
	}

	if enDefault.Get("FR") != "French" {
		t.Error("incorrect French translation for English")
	}

	frDefault := FromMap(m, "FR")
	if frDefault.Fallback != "French" {
		t.Error("incorrect fallback for French")
	}

	esDefault := FromMap(m, "ES")
	if esDefault.Fallback != "" {
		t.Error("incorrect fallback for ES")
	}
	if esDefault.Get("EN") != "English" {
		t.Error("incorrect English translation for ES")
	}
}
