package app

import "github.com/kubex/definitions-go/translation"

type Permission struct {
	Key         string
	Name        translation.Text
	Description translation.Text
}

type PermissionEffect string

const (
	PermissionEffectAllow PermissionEffect = "Allow"
	PermissionEffectDeny  PermissionEffect = "Deny"
)

type PermissionStatement struct {
	Effect     PermissionEffect
	Permission Permission
	Resource   string // path or resource indicator defined by the app
}

type PermissionPolicy struct {
	Key         string
	Name        translation.Text
	Description translation.Text
	Statements  []PermissionStatement
}
