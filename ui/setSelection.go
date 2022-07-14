package ui

import (
	"az-func-deploy/config"

	"github.com/rivo/tview"
)

func NewSetSelectionLayout(config *config.DeployConfig) *tview.Form {
	form := tview.NewForm()
	options := []string{}
	for _, set := range config.Sets {
		options = append(options, set.Name)
	}
	form.AddDropDown("Deployment Set", options, 0, func(option string, optionIndex int) {})
	return form
}
