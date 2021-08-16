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
	Icon         string           `json:"icon"`  // Default icon to use for this application
	Unify        []Navigation     `json:"unify"` // How to link with other applications
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
	PageActions           []Navigation        `json:"pageActions"`
}

type IntegrationLocation string

const (
	IntegrationLocationPageNav IntegrationLocation = "nav.page"
	IntegrationLocationAppNav  IntegrationLocation = "nav.app"
	IntegrationLocationAction  IntegrationLocation = "action"
	IntegrationLocationPanel   IntegrationLocation = "panel"
)

type LaunchMode string

const (
	LaunchModePage    LaunchMode = "page"    // (default) Render in the page
	LaunchModeModal   LaunchMode = "modal"   // Launch in a modal
	LaunchModeWindow  LaunchMode = "window"  // Launch in a new window
	LaunchModeOverlay LaunchMode = "overlay" // Right side overlay app
)

type IntegrationPoint struct {
	IntegrateApp GlobalAppID         `json:"integrateApp"`
	Location     IntegrationLocation `json:"location"`
	PathID       string              `json:"pathID"`
}

type Navigation struct {
	Icon            string           `json:"icon"` // Material Design Icon Name
	Text            translation.Text `json:"text"`
	Title           translation.Text `json:"title"`
	DestinationPath string           `json:"destinationPath"`
	SectionID       string           `json:"sectionID"`
	LaunchMode      LaunchMode       `json:"launchMode"`
	Point           IntegrationPoint `json:"point"`
}

type NavigationSection struct {
	ID       string           `json:"id"`
	Priority int32            `json:"priority"`
	Text     translation.Text `json:"text"`
}
