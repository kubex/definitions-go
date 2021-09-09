package app

import "github.com/kubex/definitions-go/translation"

type Permission struct {
	Key         string           `json:"key"`
	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description,omitempty"`
}

type PermissionEffect string

const (
	PermissionEffectAllow PermissionEffect = "Allow"
	PermissionEffectDeny  PermissionEffect = "Deny"
)

type PermissionPolicy struct {
	Uuid        string
	Name        translation.Text
	Description translation.Text
	Statements  []PermissionStatement
}

type PermissionStatement struct {
	Effect     PermissionEffect `json:"effect"`
	Permission ScopedKey        `json:"permission"`
	Resource   string           `json:"resource"` // path or resource indicator defined by the app
}
