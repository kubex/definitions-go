package app

import (
	"encoding/json"
	"errors"
	"github.com/kubex/definitions-go/translation"
)

type Definition struct {
	ID                   GlobalAppID
	Name                 translation.Text
	Description          translation.Text
	Endpoint             string
	UIMode               UIMode
	Dependencies         []GlobalAppID // Other applications this app depends on
	Permissions          []Permission  // Permissions made available by this application
	Paths                []Path
	AuthenticationHeader []string // headers to send through - possibly template/replacement based
}

func FromJson(jsonBytes []byte) (*Definition, error) {
	def := &Definition{}
	if err := json.Unmarshal(jsonBytes, def); err != nil {
		return nil, errors.New("unable to decode app definition json")
	}
	return def, nil
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
