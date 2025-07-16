package app

import "github.com/kubex/definitions-go/translation"

type IntegrationLocation string

const (
	IntegrationLocationPage   IntegrationLocation = "page"   // Top Level Page Navigation
	IntegrationLocationTab    IntegrationLocation = "tab"    // Header Navigation on Page
	IntegrationLocationAction IntegrationLocation = "action" // <app-actions>
	IntegrationLocationPanel  IntegrationLocation = "panel"  // <app-panel>
	// Deprecated: IntegrationLocationPageNav
	IntegrationLocationPageNav IntegrationLocation = IntegrationLocationTab
)

type LaunchMode string

const (
	LaunchModePage       LaunchMode = "page"   // (default) Render in the page
	LaunchModeModal      LaunchMode = "modal"  // Launch in a modal
	LaunchModeWindow     LaunchMode = "window" // Launch in a new window
	LaunchModeSlide      LaunchMode = "slide"  // Right side overlay app
	LaunchModeActionDrop LaunchMode = "adrop"  // Drop down from the action bar
	LaunchModeActionFill LaunchMode = "afill"  // Fill the action bar with the app
)

type IntegrationPoint struct {
	IntegrateApp        ScopedKey                    `json:"integrateApp,omitempty"`        // Which app to integrate into
	Location            IntegrationLocation          `json:"location,omitempty"`            // Where to place the integration
	LocationID          string                       `json:"locationID,omitempty"`          // Location ID if multiple locations available
	PathID              string                       `json:"pathID,omitempty"`              // Remote app path ID
	PreferredWidth      int                          `json:"preferredWidth,omitempty"`      // Preferred width of the integration
	MultiPanel          bool                         `json:"multiPanel,omitempty"`          // Should the integration Panel be transparent
	EntryPoint          EntryPoint                   `json:"entryPoint,omitempty"`          // How the integration is presented
	PanelActions        []EntryPoint                 `json:"panelActions,omitempty"`        // Actions to add to the panel
	PanelTabs           []EntryPoint                 `json:"panelTabs,omitempty"`           // Tabs to show - panels only
	Preferences         []IntegrationPointPreference `json:"preferences,omitempty"`         // Preferences for the integration
	RequiredPermissions []ScopedKey                  `json:"requiredPermissions,omitempty"` // Permissions that must be set for the user to see this item
}

type IntegrationPointPreference int32

const (
	IntegrationPointPreferenceNone               IntegrationPointPreference = 0
	IntegrationPointPreferencePanelTabsAsIconBar IntegrationPointPreference = 1
)

type Navigation struct {
	Title translation.Text `json:"title,omitempty"`
	Items []EntryPoint     `json:"items,omitempty"`
}

type EntryPoint struct {
	Icon                string           `json:"icon,omitempty"` // Material Design Icon Name
	Text                translation.Text `json:"text,omitempty"`
	Title               translation.Text `json:"title,omitempty"`
	DestinationPath     string           `json:"destinationPath,omitempty"`
	CountPath           string           `json:"countPath,omitempty"` // Count Path - expanding actions only
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
