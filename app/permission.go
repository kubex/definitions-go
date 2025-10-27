package app

import "github.com/kubex/definitions-go/translation"

type Permission struct {
	Key         string                           `json:"key"`
	Name        translation.Text                 `json:"name"`
	Description translation.Text                 `json:"description,omitempty"`
	Meta        []PermissionMeta                 `json:"meta,omitempty"`
	Constraints []PermissionConstraintDefinition `json:"constraints,omitempty"`
}

type PermissionMeta struct {
	Key         string           `json:"key"`
	Name        translation.Text `json:"name"`
	Description translation.Text `json:"description,omitempty"`
}

type PermissionConstraintDefinition struct {
	Field string                   `json:"field"`
	Type  PermissionConstraintType `json:"type"`
}

type PermissionConstraintType string

const (
	PermissionConstraintTypeValue PermissionConstraintType = "value"
)

type PermissionConstraintOperator string

const (
	PermissionConstraintOperatorLessThan           PermissionConstraintOperator = "lessThan"
	PermissionConstraintOperatorGreaterThan        PermissionConstraintOperator = "greaterThan"
	PermissionConstraintOperatorEqual              PermissionConstraintOperator = "equal"
	PermissionConstraintOperatorNotEqual           PermissionConstraintOperator = "notEqual"
	PermissionConstraintOperatorLessThanOrEqual    PermissionConstraintOperator = "lessThanOrEqual"
	PermissionConstraintOperatorGreaterThanOrEqual PermissionConstraintOperator = "greaterThanOrEqual"
)

var PermissionConstraintOperatorDisplayValues = map[PermissionConstraintOperator]string{
	PermissionConstraintOperatorEqual:              "=",
	PermissionConstraintOperatorNotEqual:           "!=",
	PermissionConstraintOperatorLessThan:           "<",
	PermissionConstraintOperatorLessThanOrEqual:    "<=",
	PermissionConstraintOperatorGreaterThan:        ">",
	PermissionConstraintOperatorGreaterThanOrEqual: ">=",
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

const PermissionResourceAll = "*"

type PermissionStatement struct {
	Effect      PermissionEffect       `json:"e"`
	Permission  ScopedKey              `json:"p"`
	Resource    string                 `json:"r"` // path or resource indicator defined by the app
	Meta        map[string]string      `json:"m,omitempty"`
	Constraints []PermissionConstraint `json:"c,omitempty"`
}

type PermissionConstraint struct {
	Field    string                       `json:"f"`
	Type     PermissionConstraintType     `json:"g"`
	Operator PermissionConstraintOperator `json:"o"`
	Value    interface{}                  `json:"v"`
}
