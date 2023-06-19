package app

import "github.com/kubex/definitions-go/translation"

type Path struct {
	ID          string           `json:"id"` // Allow the path to be linked
	Name        translation.Text `json:"name,omitempty"`
	Description translation.Text `json:"description,omitempty"`

	Path   string `json:"path"` // with replacements, matches start, locating the most specific
	Method string `json:"method,omitempty"`

	RequestPermissions  []ScopedKey `json:"requestPermissions,omitempty"`  // Permissions that should be sent to this path
	RequiredPermissions []ScopedKey `json:"requiredPermissions,omitempty"` // Permissions that must be set for the user to call this page

	Navigation []Navigation `json:"navigation,omitempty"`
	Actions    []Navigation `json:"actions,omitempty"`
}
