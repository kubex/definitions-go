package app

import (
	"github.com/kubex/definitions-go/translation"
)

type Definition struct {
	ID           GlobalAppID
	Name         translation.Text
	Description  translation.Text
	UIMode       UIMode
	Dependencies []GlobalAppID // Other applications this app depends on
	Permissions  []Permission  // Permissions made available by this application
	Paths        []Path
}

type Path struct {
	ID          string // Allow the path to be linked
	Name        translation.Text
	Description translation.Text

	Path                string
	Method              string
	RequestPermissions  []Permission // Permissions that should be sent to this path
	RequiredPermissions []Permission // Permissions that must be set for the user to call this page
}
