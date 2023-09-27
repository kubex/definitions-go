package app

import "github.com/kubex/definitions-go/translation"

type SettingsPage struct {
	ID           string
	Name         translation.Text
	Panels       []SettingsPanel
	AdvancedPath string
}

type SettingsPanel struct {
	Name        translation.Text
	Description translation.Text
	Settings    []Property
	Order       int
}
