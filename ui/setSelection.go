package ui

import (
	"az-func-deploy/config"

	"github.com/rivo/tview"
)

func NewSetSelectionLayout(dcon *config.DeployConfig) *tview.Form {
	form := tview.NewForm()
	optionsSet := []string{}
	for _, set := range dcon.Sets {
		optionsSet = append(optionsSet, set.Name)
	}
	form.AddDropDown("Deployment Set", optionsSet, dcon.CurrentSet, func(option string, optionIndex int) {})
	optionsMethods := config.GetDeploymentMethods()
	form.AddDropDown("Deployment Set", optionsMethods, config.GetDeploymentMethodIndex(dcon.Method), func(option string, optionIndex int) {})
	return form
}
