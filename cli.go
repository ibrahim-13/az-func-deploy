package main

import (
	"az-func-deploy/config"
	"az-func-deploy/ui"

	"github.com/rivo/tview"
)

const (
	__pageNameSetSelection string = "Select Deployment Set"
	__pageNameDeployment   string = "Deployment"
)

func main() {
	config := config.ReadConfigOrPanic()
	app := tview.NewApplication()
	pages := tview.NewPages()
	pageSelection := ui.NewSetSelectionLayout(config, func() {
		pages.SwitchToPage(__pageNameDeployment)
		pages.SetTitle(__pageNameDeployment)
	})
	pageDeployment := tview.NewBox().SetTitle("Deployment")

	pages.SetBorder(true).
		SetTitle(__pageNameSetSelection)

	pages.AddPage(__pageNameSetSelection, pageSelection, true, true)
	pages.AddPage(__pageNameDeployment, pageDeployment, true, false)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
