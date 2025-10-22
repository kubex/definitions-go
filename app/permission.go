package app

import "github.com/kubex/definitions-go/translation"

type Permission struct {
	Key         string           `json:"key"`
	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description,omitempty"`
	Meta        []PermissionMeta `json:"meta,omitempty"`
}

type PermissionMeta struct {
	Key         string           `json:"key"`
	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description,omitempty"`
}

type PermissionConstraint struct {
	Type     PermissionConstraintType     `json:"type"`
	Operator PermissionConstraintOperator `json:"operator"`
	Value    interface{}                  `json:"value"`
}

type PermissionConstraintType string

const (
	TypeValue    PermissionConstraintType = "value"
	TypeLocation PermissionConstraintType = "location"
)

type PermissionConstraintOperator string

const (
	OperatorLessThan           PermissionConstraintOperator = "lessThan"
	OperatorGreaterThan        PermissionConstraintOperator = "greaterThan"
	OperatorEqual              PermissionConstraintOperator = "equal"
	OperatorNotEqual           PermissionConstraintOperator = "notEqual"
	OperatorLessThanOrEqual    PermissionConstraintOperator = "lessThanOrEqual"
	OperatorGreaterThanOrEqual PermissionConstraintOperator = "greaterThanOrEqual"
)

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

const PermissionResourceAll = "*"

type PermissionStatement struct {
	Effect     PermissionEffect  `json:"e"`
	Permission ScopedKey         `json:"p"`
	Resource   string            `json:"r"` // path or resource indicator defined by the app
	Meta       map[string]string `json:"m,omitempty"`
}
