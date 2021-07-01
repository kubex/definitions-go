package app

import "github.com/kubex/definitions-go/translation"

type Vendor struct {
	ID                 string
	Name               translation.Text
	Description        translation.Text
	AuthenticationData map[string]string
}
