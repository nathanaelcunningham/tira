
package main

import (
	"github.com/rivo/tview"
)

func main() {
    app := tview.NewApplication()

    flex := tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(tview.NewBox().SetBorder(true).SetTitle("Server"), 3, 1, false).
        AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
            AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
                AddItem(tview.NewBox().SetBorder(true).SetTitle("Projects"), 0, 1, false).
                AddItem(tview.NewBox().SetBorder(true).SetTitle("Users"), 0, 1, false).
                AddItem(tview.NewBox().SetBorder(true).SetTitle("Status"), 0, 1, false), 0, 1, false).
            AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
                AddItem(tview.NewBox().SetBorder(true).SetTitle("Issues"), 0, 1, false),0,4,false), 0,1,false)


    if err := app.SetRoot(flex, true).Run(); err != nil {
        panic(err)
    }
}

