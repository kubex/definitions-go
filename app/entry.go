package app

import "github.com/kubex/definitions-go/translation"

type IntegrationLocation string

const (
	IntegrationLocationPageNav IntegrationLocation = "nav.page" // Load within the page
	IntegrationLocationAction  IntegrationLocation = "action"   // Perform an action, expecting a toast response
	IntegrationLocationPanel   IntegrationLocation = "panel"    // Load within a panel
)

type LaunchMode string

const (
	LaunchModePage   LaunchMode = "page"   // (default) Render in the page
	LaunchModeModal  LaunchMode = "modal"  // Launch in a modal
	LaunchModeWindow LaunchMode = "window" // Launch in a new window
	LaunchModeSlide  LaunchMode = "slide"  // Right side overlay app
)

type IntegrationPoint struct {
	IntegrateApp        GlobalAppID         `json:"integrateApp,omitempty"`        // Which app to integrate into
	Location            IntegrationLocation `json:"location,omitempty"`            // Where to place the integration
	LocationID          string              `json:"locationID,omitempty"`          // Location ID if multiple locations available
	PathID              string              `json:"pathID,omitempty"`              // Remote app path ID
	PreferredWidth      int                 `json:"preferredWidth,omitempty"`      // Preferred width of the integration
	EntryPoint          EntryPoint          `json:"entryPoint,omitempty"`          // How the integration is presented
	PanelActions        []EntryPoint        `json:"panelActions,omitempty"`        // Actions to add to the panel
	PanelTabs           []EntryPoint        `json:"panelTabs,omitempty"`           // Tabs to show - panels only
	RequiredPermissions []ScopedKey         `json:"requiredPermissions,omitempty"` // Permissions that must be set for the user to see this item
}

type EntryPoint struct {
	Icon                string           `json:"icon,omitempty"` // Material Design Icon Name
	Text                translation.Text `json:"text,omitempty"`
	Title               translation.Text `json:"title,omitempty"`
	DestinationPath     string           `json:"destinationPath,omitempty"`
	LaunchMode          LaunchMode       `json:"launchMode,omitempty"`
	RequiredPermissions []ScopedKey      `json:"requiredPermissions,omitempty"` // Permissions that must be set for the user to see this item
}

func NewEntryPoint(destination string, text translation.Text) *EntryPoint {
	return &EntryPoint{
		Text:            text,
		DestinationPath: destination,
	}
}

func (n *EntryPoint) WithLaunchMode(launchMode LaunchMode) *EntryPoint {
	n.LaunchMode = launchMode
	return n
}

func (n *EntryPoint) WithIcon(icon string) *EntryPoint {
	n.Icon = icon
	return n
}

func (n *EntryPoint) WithTitle(title translation.Text) *EntryPoint {
	n.Title = title
	return n
}
