package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func generateFlagSplashScreen(segments int, generator func(int, int, int) *ebiten.Image) *ebiten.Image {
	splash := ebiten.NewImage(w, h)
	splash.Fill(bg)

	flag := generator(w, w*3/5, segments)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, float64(h/2-flag.Bounds().Dy()/2))

	splash.DrawImage(flag, opts)

	message := "'esc': continue"
	text.Draw(splash, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, h-10, fg)

	return splash
}
