package app

import "github.com/kubex/definitions-go/translation"

type Permission struct {
	Key         PermissionKey    `json:"key"`
	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description,omitempty"`
}

type PermissionEffect string

const (
	PermissionEffectAllow PermissionEffect = "Allow"
	PermissionEffectDeny  PermissionEffect = "Deny"
)

type PermissionKey struct {
	GlobalAppID
	Key string
}

type PermissionStatement struct {
	Effect     PermissionEffect `json:"e"`
	Permission PermissionKey    `json:"k"`
	Resource   string           `json:"r"` // path or resource indicator defined by the app
}

type PermissionPolicy struct {
	Uuid        string
	Key         string
	Name        translation.Text
	Description translation.Text
	Statements  []PermissionStatement
}
