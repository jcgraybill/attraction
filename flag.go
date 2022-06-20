package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// TODO: finish me!

func generateFlagSplashScreen(segments int, generator func(int) *ebiten.Image) *ebiten.Image {
	splash := ebiten.NewImage(w, h)
	splash.Fill(bg)

	flag := ebiten.NewImage(w, w*3/5)
	flag.Fill(fg)

	stripe := ebiten.NewImage(w, flag.Bounds().Dy()/6)
	stripeOpts := &ebiten.DrawImageOptions{}
	for i := 0; i < segments; i++ {
		stripe.Fill(rainbowColors[i])
		flag.DrawImage(stripe, stripeOpts)
		stripeOpts.GeoM.Translate(0, float64(stripe.Bounds().Dy()))
	}

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, float64(h/2-flag.Bounds().Dy()/2))

	splash.DrawImage(flag, opts)

	message := "press esc to continue"
	text.Draw(splash, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, h-10, fg)

	return splash
}
