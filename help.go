package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var helpSplash *ebiten.Image

func generateHelpSplashScreen() *ebiten.Image {
	splash := ebiten.NewImage(w, h)
	splash.Fill(bg)

	message := "this will be the help screen"
	text.Draw(splash, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, h-10, fg)

	return splash
}
