package app

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kubex/definitions-go/translation"
)

const VendorAppID = "kx" // setting your application ID to this value should be done for vendor shared

type Definition struct {
	ID          GlobalAppID `json:"id"`
	Endpoint    string      `json:"endpoint,omitempty"`
	DefaultPath string      `json:"defaultPath,omitempty"` // Default path to use when opening the app

	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description,omitempty"`

	UIMode   UIMode   `json:"UIMode,omitempty"`
	Category Category `json:"category,omitempty"`
	Icon     string   `json:"icon,omitempty"` // Default icon to use for this application

	Dependencies    []GlobalAppID      `json:"dependencies,omitempty"` // Other applications this app depends on
	Permissions     []Permission       `json:"permissions,omitempty"`  // Permissions made available by this application
	Roles           []PermissionPolicy `json:"roles,omitempty"`        // Roles made available by this application
	Paths           []Path             `json:"paths,omitempty"`
	Unify           []IntegrationPoint `json:"unify,omitempty"` // How to link with other applications
	ActivationSteps []ActivationStep   `json:"activationSteps,omitempty"`
	ListenToEvents  []ScopedKey        `json:"listenToEvents,omitempty"`
	Navigation      []Navigation       `json:"navigation,omitempty"` // App navigation
	NavigationUri   string             `json:"navigationUri,omitempty"`

	Configuration     []SettingsPage
	ConfigurationPath string // the path to use when settings are managed by the app

	Homepage       string `json:"homepage,omitempty"`       // https:// url
	TermsOfService string `json:"termsOfService,omitempty"` // https:// url
	PrivacyPolicy  string `json:"privacyPolicy,omitempty"`  // https:// url

	SupportEmail string `json:"supportEmail,omitempty"`

	PrefixRedirect map[string]string `json:"prefixRedirect,omitempty"` // Matching prefixes to redirect  e.g. [CST:CST => 'view/$1'] $1 includes the prefix
	QuickCodes     map[string]string `json:"quickCodes,omitempty"`     // Matching Codes to redirect  e.g. [CST => 'view/$1'] $1 is replaced by everything after the code

	PermittedProxyPaths []string `json:"permittedProxyPaths,omitempty"` // Paths that can be proxied by the platform, without auth / modification

	Hash string `json:"hash,omitempty"` // Hash of the definition for change detection, latest hash can be returned in HealthResponse
}

func (d *Definition) GetHash(updateIfEmpty bool) string {
	if d == nil {
		return ""
	}
	currentHash := d.Hash
	d.Hash = ""
	jsonBytes, _ := json.Marshal(d)
	result := fmt.Sprintf("%x", md5.Sum(jsonBytes))
	if currentHash == "" {
		if updateIfEmpty {
			d.Hash = result
		}
		return result
	}
	d.Hash = currentHash
	return d.Hash
}

func (d *Definition) WithPath(path ...Path) *Definition {
	d.Paths = append(d.Paths, path...)
	return d
}

func FromJson(jsonBytes []byte) (*Definition, error) {
	def := &Definition{}
	if err := json.Unmarshal(jsonBytes, def); err != nil {
		return nil, errors.New("unable to decode app definition json")
	}
	return def, nil
}
