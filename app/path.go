package app

import "github.com/kubex/definitions-go/translation"

type Path struct {
	ID     string `json:"id"`   // Allow the path to be linked
	Path   string `json:"path"` // with replacements, matches start, locating the most specific
	Method string `json:"method,omitempty"`

	Name           translation.Text `json:"name,omitempty"`
	Description    translation.Text `json:"description,omitempty"`
	HideHeader     bool             `json:"hideHeader,omitempty"`
	HideBreadcrumb bool             `json:"hideBreadcrumb,omitempty"`
	PromptOnExit   bool             `json:"promptOnExit,omitempty"`

	RequestPermissions  []ScopedKey `json:"requestPermissions,omitempty"`  // Permissions that should be sent to this path
	RequiredPermissions []ScopedKey `json:"requiredPermissions,omitempty"` // Permissions that must be set for the user to call this page

	Navigation []EntryPoint `json:"navigation,omitempty"`
	Actions    []EntryPoint `json:"actions,omitempty"`
}

func NewPath(id, path string) *Path {
	return &Path{ID: id, Path: path}
}

func (p *Path) WithNavigation(navigation ...EntryPoint) *Path {
	p.Navigation = append(p.Navigation, navigation...)
	return p
}

func (p *Path) WithRequestPermissions(permissions ...ScopedKey) *Path {
	p.RequestPermissions = append(p.RequestPermissions, permissions...)
	return p
}

func (p *Path) WithRequiredPermissions(permissions ...ScopedKey) *Path {
	p.RequiredPermissions = append(p.RequiredPermissions, permissions...)
	return p
}

func (p *Path) WithActions(actions ...EntryPoint) *Path {
	p.Actions = append(p.Actions, actions...)
	return p
}

func (p *Path) WithMethod(method string) *Path {
	p.Method = method
	return p
}

func (p *Path) WithName(name translation.Text) *Path {
	p.Name = name
	return p
}

func (p *Path) WithDescription(description translation.Text) *Path {
	p.Description = description
	return p
}

func (p *Path) WithoutBreadCrumbs() *Path {
	p.HideBreadcrumb = true
	return p
}

func (p *Path) WithoutHeader() *Path {
	p.HideHeader = true
	return p
}
