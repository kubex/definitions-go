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

func NewPermissionKey(key string, gaid *GlobalAppID) PermissionKey {
	if gaid == nil {
		return PermissionKey{Key: key}
	}
	return PermissionKey{Key: key, GlobalAppID: *gaid}
}

type PermissionStatement struct {
	Effect     PermissionEffect `json:"effect"`
	Permission PermissionKey    `json:"permission"`
	Resource   string           `json:"resource"` // path or resource indicator defined by the app
}

type PermissionPolicy struct {
	Uuid        string
	Key         string
	Name        translation.Text
	Description translation.Text
	Statements  []PermissionStatement
}
