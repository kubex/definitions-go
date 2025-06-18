package app

import "github.com/kubex/definitions-go/translation"

type SLI struct {
	Metric string
	Name   translation.Text
	Labels []string
}
