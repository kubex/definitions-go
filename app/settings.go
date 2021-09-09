package app

import "github.com/kubex/definitions-go/translation"

type Settings struct {
	Panels []SettingsPanel
}

type SettingsPanel struct {
	Name        translation.Text
	Description translation.Text
	Settings    []Attribute
}
