package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var helpSplash *ebiten.Image

func generateHelpSplashScreen() *ebiten.Image {
	var topics [6]string
	//	var images [6]*ebiten.Image

	topics = [6]string{
		"move colored circles into\nmatching slots.",
		"use arrow keys to pulse\nan electromagnet.",
		"magnetic pieces will move\none cell in that direction.",
		"moving pieces push other\nmoveable pieces.",
		"you can push hollow\nblocks.",
		"black squares don't move.",
	}

	splash := ebiten.NewImage(w, h)
	splash.Fill(bg)

	for i, topic := range topics {
		text.Draw(splash, topic, *regular, 100, h*i/(len(topics)+1)+h/(len(topics)+1), fg)
	}

	message := "'esc': continue"
	text.Draw(splash, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, h-10, fg)

	return splash
}
