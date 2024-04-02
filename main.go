package main

import (
	"GroupieTracker/core"
)

func main() {
	// core.W.Resize(fyne.NewSize(1000, 600))
	 core.W.SetMainMenu(core.CreateMainMenu())

	core.ShowHomePage()
	// core.a.Run()

	core.W.ShowAndRun()
}
