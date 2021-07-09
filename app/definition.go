package app

import (
	"encoding/json"
	"errors"
	"github.com/kubex/definitions-go/translation"
)

type Definition struct {
	ID           GlobalAppID      `json:"id"`
	Name         translation.Text `json:"name"`
	Description  translation.Text `json:"description"`
	Endpoint     string           `json:"endpoint"`
	UIMode       UIMode           `json:"UIMode"`
	Dependencies []GlobalAppID    `json:"dependencies"` // Other applications this app depends on
	Permissions  []Permission     `json:"permissions"`  // Permissions made available by this application
	Paths        []Path           `json:"paths"`
	Icon         string           `json:"icon"` // Default icon to use for this application
}

func FromJson(jsonBytes []byte) (*Definition, error) {
	def := &Definition{}
	if err := json.Unmarshal(jsonBytes, def); err != nil {
		return nil, errors.New("unable to decode app definition json")
	}
	return def, nil
}

type Path struct {
	ID          string           `json:"id"` // Allow the path to be linked
	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description"`

	Path                string       `json:"path"`
	Method              string       `json:"method"`
	RequestPermissions  []Permission `json:"requestPermissions"`  // Permissions that should be sent to this path
	RequiredPermissions []Permission `json:"requiredPermissions"` // Permissions that must be set for the user to call this page

	AppNavigationSections []NavigationSection `json:"appNavigationSections"`
	AppNavigation         []Navigation        `json:"appNavigation"`
	PageNavigation        []Navigation        `json:"pageNavigation"`
}

type IntegrationPoint struct {
	IntegrateApp GlobalAppID `json:"integrateApp"`
	PathID       string      `json:"pathID"`
}

type Navigation struct {
	Text            translation.Text `json:"text"`
	Title           translation.Text `json:"title"`
	DestinationPath string           `json:"destinationPath"`
	SectionID       string           `json:"sectionID"`
	Point           IntegrationPoint `json:"point"`
}

type NavigationSection struct {
	ID   string           `json:"id"`
	Text translation.Text `json:"text"`
}
