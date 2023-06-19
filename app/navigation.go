package app

import "github.com/kubex/definitions-go/translation"

type IntegrationLocation string

const (
	IntegrationLocationPageNav IntegrationLocation = "nav.page" // Load within the page
	IntegrationLocationAppNav  IntegrationLocation = "nav.app"  // Load as the primary app
	IntegrationLocationAction  IntegrationLocation = "action"   // Perform an action, expecting a toast response
	IntegrationLocationPanel   IntegrationLocation = "panel"    // Load within a panel
)

type LaunchMode string

const (
	LaunchModePage  LaunchMode = "page"  // (default) Render in the page
	LaunchModeModal LaunchMode = "modal" // Launch in a modal
	//LaunchModeLightbox LaunchMode = "lightbox" // Launch in a modal
	LaunchModeWindow  LaunchMode = "window"  // Launch in a new window
	LaunchModeOverlay LaunchMode = "overlay" // Right side overlay app
)

type IntegrationPoint struct {
	IntegrateApp GlobalAppID         `json:"integrateApp,omitempty"` // what app to load into
	Location     IntegrationLocation `json:"location,omitempty"`     // Where to load the app
	PathID       string              `json:"pathID,omitempty"`       //remote app path ID
	Navigation   Navigation          `json:"navigation,omitempty"`   // Navigation, when panel location, dst path is used
}

type Navigation struct {
	Icon            string           `json:"icon,omitempty"` // Material Design Icon Name
	Text            translation.Text `json:"text,omitempty"`
	Title           translation.Text `json:"title,omitempty"`
	DestinationPath string           `json:"destinationPath,omitempty"`
	LaunchMode      LaunchMode       `json:"launchMode,omitempty"`
}

type NavigationSection struct {
	ID       string           `json:"id"`
	Priority int32            `json:"priority,omitempty"`
	Text     translation.Text `json:"text,omitempty"`
}
