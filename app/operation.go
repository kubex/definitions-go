package app

// Operation describes a callable API endpoint for app-to-app communication.
type Operation struct {
	Key                 string      `json:"key"`
	Path                string      `json:"path,omitempty"`
	Method              string      `json:"method,omitempty"`
	Inputs              []Property  `json:"inputs,omitempty"`
	ResponseCode        int32       `json:"responseCode,omitempty"`
	Outputs             []Property  `json:"outputs,omitempty"`
	Description         string      `json:"description,omitempty"`
	RequiredPermissions []ScopedKey `json:"requiredPermissions,omitempty"`
	UserContextRequired bool        `json:"userContextRequired,omitempty"`
}
