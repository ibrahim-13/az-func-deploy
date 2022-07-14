package main

import (
	"az-func-deploy/config"
	"az-func-deploy/ui"

	"github.com/rivo/tview"
)

const (
	__pageNameSetSelection string = "Select Deployment Set"
	__pageNameDeployment          = "Deployment"
)

func main() {
	config := config.ReadConfigOrPanic()
	app := tview.NewApplication()
	setSelection := ui.NewSetSelectionLayout(config)
	pages := tview.NewPages()

	pages.SetBorder(true).
		SetTitle(__pageNameSetSelection)

	pages.AddPage(__pageNameSetSelection, setSelection, true, true)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}
