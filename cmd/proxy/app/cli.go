package app

import "github.com/desertbit/grumble"

// App is used to register the grumble
var App = grumble.New(&grumble.Config{
	Name:                  "Proxify-ng",
	Description:           "Proxify-ng",
	HelpHeadlineUnderline: true,
	HelpSubCommands:       true,
})

func init() {
	App.SetPrintASCIILogo(func(a *grumble.App) {
		a.Println(":)")
	})
}
