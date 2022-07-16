package main

import (
	"az-func-deploy/config"
	"az-func-deploy/deployment"
	"az-func-deploy/ui"
	"os"

	"github.com/rivo/tview"
)

const (
	__pageNameSetSelection string = "Select Deployment Set"
	__pageNameDeployment   string = "Deployment"
)

func main() {
	conf := config.ReadConfigOrPanic()

	if len(os.Args) > 1 && os.Args[1] == "--cli" {
		deployment.DeployFunctions(conf, os.Stdout, true)
	} else {
		app := tview.NewApplication()
		pages := tview.NewPages()
		pageDeployment := tview.NewTextView()

		pageDeployment.
			SetDynamicColors(true).
			SetChangedFunc(func() {
				app.Draw()
				pageDeployment.ScrollToEnd()
			})
		outputWriter := tview.ANSIWriter(pageDeployment)
		pageSelection := ui.NewSetSelectionLayout(conf, func() {
			pages.SwitchToPage(__pageNameDeployment)
			pages.SetTitle(__pageNameDeployment)
			go deployment.DeployFunctions(conf, outputWriter, false)
		})

		pages.SetBorder(true).
			SetTitle(__pageNameSetSelection)

		pages.AddPage(__pageNameSetSelection, pageSelection, true, true)
		pages.AddPage(__pageNameDeployment, pageDeployment, true, false)

		if err := app.SetRoot(pages, true).Run(); err != nil {
			panic(err)
		}
	}
}
