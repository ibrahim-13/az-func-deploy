package ui

import (
	"az-func-deploy/config"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func removeFuncs(form *tview.Form, dcon *config.DeployConfig) {
	currentSet := dcon.Sets[dcon.CurrentSet]
	for _, funcInfo := range currentSet.FuncInfos {
		form.RemoveFormItem(form.GetFormItemIndex(funcInfo.FuncName))
	}
}

func addFuncs(form *tview.Form, dcon *config.DeployConfig) {
	currentSet := dcon.Sets[dcon.CurrentSet]
	for i, funcInfo := range currentSet.FuncInfos {
		funcIndex := i
		form.AddCheckbox(funcInfo.FuncName, funcInfo.ShouldRun, func(checked bool) {
			dcon.Sets[dcon.CurrentSet].FuncInfos[funcIndex].ShouldRun = checked
			dcon.WriteConfig()
		})
	}
}

func NewSetSelectionLayout(dcon *config.DeployConfig, onDeploy func()) *tview.Form {
	form := tview.NewForm()
	optionsSet := []string{}
	for _, set := range dcon.Sets {
		optionsSet = append(optionsSet, set.Name)
	}
	form.AddDropDown("Deployment Set", optionsSet, dcon.CurrentSet, func(_ string, optionIndex int) {
		if optionIndex != dcon.CurrentSet {
			removeFuncs(form, dcon)
			dcon.CurrentSet = optionIndex
			addFuncs(form, dcon)
			dcon.WriteConfig()
		}
	})
	optionsMethods := config.GetDeploymentMethods()
	form.AddDropDown("Deployment Method", optionsMethods, config.GetDeploymentMethodIndex(dcon.Method), func(option string, _ int) {
		if option != dcon.Method {
			dcon.Method = option
			dcon.WriteConfig()
		}
	})
	addFuncs(form, dcon)
	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'd' {
			onDeploy()
			return nil
		}
		return event
	})
	form.AddButton("(d) Deploy", onDeploy)

	return form
}
