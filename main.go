package main

import (
    "GroupieTracker/core"
    "fyne.io/fyne/v2"
)

func main() {
    core.W.Resize(fyne.NewSize(1000, 600))
    core.W.SetMainMenu(core.CreateMainMenu())

    core.ShowHomePage()

    core.W.ShowAndRun()
}
