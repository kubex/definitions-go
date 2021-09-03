package app

import (
	"encoding/json"
	"errors"
	"github.com/kubex/definitions-go/translation"
)

type Definition struct {
	ID           GlobalAppID      `json:"id"`
	Name         translation.Text `json:"name"`
	Description  translation.Text `json:"description,omitempty"`
	Endpoint     string           `json:"endpoint,omitempty"`
	UIMode       UIMode           `json:"UIMode,omitempty"`
	Category     Category         `json:"category,omitempty"`
	Icon         string           `json:"icon,omitempty"`         // Default icon to use for this application
	Dependencies []GlobalAppID    `json:"dependencies,omitempty"` // Other applications this app depends on
	Permissions  []Permission     `json:"permissions,omitempty"`  // Permissions made available by this application
	Paths        []Path           `json:"paths,omitempty"`
	Unify        []Navigation     `json:"unify,omitempty"` // How to link with other applications
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
	Name        translation.Text `json:"name,omitempty"`
	Description translation.Text `json:"description,omitempty"`

	Path                string       `json:"path"`
	Method              string       `json:"method,omitempty"`
	RequestPermissions  []Permission `json:"requestPermissions,omitempty"`  // Permissions that should be sent to this path
	RequiredPermissions []Permission `json:"requiredPermissions,omitempty"` // Permissions that must be set for the user to call this page

	AppNavigationSections []NavigationSection `json:"appNavigationSections,omitempty"`
	AppNavigation         []Navigation        `json:"appNavigation,omitempty"`
	PageNavigation        []Navigation        `json:"pageNavigation,omitempty"`
	PageActions           []Navigation        `json:"pageActions,omitempty"`
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
	LaunchModePage     LaunchMode = "page"     // (default) Render in the page
	LaunchModeModal    LaunchMode = "modal"    // Launch in a modal
	LaunchModeLightbox LaunchMode = "lightbox" // Launch in a modal
	LaunchModeWindow   LaunchMode = "window"   // Launch in a new window
	LaunchModeOverlay  LaunchMode = "overlay"  // Right side overlay app
)

type IntegrationPoint struct {
	IntegrateApp GlobalAppID         `json:"integrateApp,omitempty"`
	Location     IntegrationLocation `json:"location,omitempty"`
	PathID       string              `json:"pathID,omitempty"`
}

type Navigation struct {
	Icon            string           `json:"icon,omitempty"` // Material Design Icon Name
	Text            translation.Text `json:"text,omitempty"`
	Title           translation.Text `json:"title,omitempty"`
	DestinationPath string           `json:"destinationPath,omitempty"`
	SectionID       string           `json:"sectionID,omitempty"`
	LaunchMode      LaunchMode       `json:"launchMode,omitempty"`
	Point           IntegrationPoint `json:"point,omitempty"`
}

type NavigationSection struct {
	ID       string           `json:"id"`
	Priority int32            `json:"priority,omitempty"`
	Text     translation.Text `json:"text,omitempty"`
}
