package main

import (
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func generateStartingSplashScreen() *ebiten.Image {
	var splash = ebiten.NewImage(w, h)
	splash.Fill(bg)

	imgBytes, err := gd.ReadFile("assets/magnet-ebiten.png")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		panic(err)
	}

	opts := &ebiten.DrawImageOptions{}
	// TODO resize by hand to anti-alias
	opts.GeoM.Scale(float64(w)/float64(img.Bounds().Dx()), float64(w)/float64(img.Bounds().Dx()))
	opts.GeoM.Translate(0, float64(h/2)-(float64(img.Bounds().Dy())*float64(w)/float64(img.Bounds().Dx()))/2)
	splash.DrawImage(ebiten.NewImageFromImage(img), opts)

	message := "ebitengine game jam"
	text.Draw(splash, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, h/3, fg)

	return splash
}
