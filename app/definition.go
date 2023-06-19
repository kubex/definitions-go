package app

import (
	"encoding/json"
	"errors"

	"github.com/kubex/definitions-go/translation"
)

const VendorAppID = "kx" // setting your application ID to this value should be done for vendor shared

type Definition struct {
	ID              GlobalAppID        `json:"id"`
	Name            translation.Text   `json:"name"`
	Description     translation.Text   `json:"description,omitempty"`
	Endpoint        string             `json:"endpoint,omitempty"`
	UIMode          UIMode             `json:"UIMode,omitempty"`
	Category        Category           `json:"category,omitempty"`
	Icon            string             `json:"icon,omitempty"`         // Default icon to use for this application
	Dependencies    []GlobalAppID      `json:"dependencies,omitempty"` // Other applications this app depends on
	Permissions     []Permission       `json:"permissions,omitempty"`  // Permissions made available by this application
	Navigation      []Navigation       `json:"navigation,omitempty"`   // Global app navigation
	Paths           []Path             `json:"paths,omitempty"`
	Unify           []IntegrationPoint `json:"unify,omitempty"` // How to link with other applications
	ActivationSteps []ActivationStep   `json:"activationSteps,omitempty"`

	Homepage       string `json:"homepage,omitempty"`       // https:// url
	TermsOfService string `json:"termsOfService,omitempty"` // https:// url
	PrivacyPolicy  string `json:"privacyPolicy,omitempty"`  // https:// url

	SupportEmail string `json:"supportEmail,omitempty"`
}

func (d *Definition) WithPath(path Path) *Definition {
	d.Paths = append(d.Paths, path)
	return d
}

func FromJson(jsonBytes []byte) (*Definition, error) {
	def := &Definition{}
	if err := json.Unmarshal(jsonBytes, def); err != nil {
		return nil, errors.New("unable to decode app definition json")
	}
	return def, nil
}
