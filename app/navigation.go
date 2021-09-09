package app

import "github.com/kubex/definitions-go/translation"

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
