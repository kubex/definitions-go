package app

import (
	"encoding/json"
	"errors"
	"github.com/kubex/definitions-go/translation"
)

type Vendor struct {
	ID                 string
	Name               translation.Text
	Description        translation.Text
	AuthenticationData map[string]string
}

func VendorFromJson(jsonBytes []byte) (*Vendor, error) {
	def := &Vendor{}
	if err := json.Unmarshal(jsonBytes, def); err != nil {
		return nil, errors.New("unable to decode vendor definition json")
	}
	return def, nil
}
