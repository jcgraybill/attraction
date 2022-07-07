package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var helpSplash *ebiten.Image
var helpImageW = 100
var helpImageH = 60

func generateHelpSplashScreen() *ebiten.Image {
	var topics [6]string
	var images [6]*ebiten.Image

	topics = [6]string{
		"use arrow keys to pulse\nan electromagnet.",
		"magnetic pieces will slide\none cell in that direction.",
		"moving pieces push other\nmovable pieces.",
		"push colored circles into\nmatching slots to win.",
		"you can push hollow\nblocks.",
		"black squares don't move.",
	}

	images = [6]*ebiten.Image{
		generateHelpImageOne(),
		generateHelpImageTwo(),
		generateHelpImageThree(),
		generateHelpImageFour(),
		generateHelpImageFive(),
		generateHelpImageSix(),
	}

	splash := ebiten.NewImage(w, h)
	splash.Fill(bg)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, float64(text.BoundString(*regular, "X").Dy()/2+10))
	for i, topic := range topics {
		splash.DrawImage(images[i], opts)
		opts.GeoM.Translate(0, float64(h/(len(topics))))
		text.Draw(splash, topic, *regular, 100, h*i/(len(topics))+text.BoundString(*regular, "X").Dy()+10, fg)
	}

	message := "'esc': continue"
	text.Draw(splash, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, h-10, fg)

	return splash
}

func generateHelpImageOne() *ebiten.Image {
	arrowDimensions := text.BoundString(*regular, "X").Dy()
	img := ebiten.NewImage(helpImageW, helpImageH)
	img.Fill(bg)
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(float64(helpImageW/2-arrowDimensions/2), 0)
	img.DrawImage(generateUpArrow(), opts)

	opts.GeoM.Reset()
	opts.GeoM.Translate(float64(helpImageW/2-arrowDimensions*2), float64(helpImageH/2-arrowDimensions/2))
	img.DrawImage(generateLeftArrow(), opts)

	opts.GeoM.Reset()
	opts.GeoM.Translate(float64(helpImageW/2+arrowDimensions), float64(helpImageH/2-arrowDimensions/2))
	img.DrawImage(generateRightArrow(), opts)

	opts.GeoM.Reset()
	opts.GeoM.Translate(float64(helpImageW/2-arrowDimensions/2), float64(helpImageH-arrowDimensions))
	img.DrawImage(generateDownArrow(), opts)

	return img
}

func generateHelpImageTwo() *ebiten.Image {
	img := ebiten.NewImage(helpImageW, helpImageH)
	img.Fill(bg)
	opts := &ebiten.DrawImageOptions{}
	img.DrawImage(generateMagnetImage(), nil)
	opts.GeoM.Translate(30, 7)
	img.DrawImage(generateRightArrow(), opts)
	return img
}

func generateHelpImageThree() *ebiten.Image {
	img := ebiten.NewImage(helpImageW, helpImageH)
	img.Fill(bg)
	img.DrawImage(generateMagnetImage(), nil)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(30, 0)
	img.DrawImage(generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 255}), opts)
	opts.GeoM.Translate(30, 7)
	img.DrawImage(generateRightArrow(), opts)
	return img
}

func generateHelpImageFour() *ebiten.Image {
	img := ebiten.NewImage(helpImageW, helpImageH)
	img.Fill(bg)
	img.DrawImage(generateGemImage(color.RGBA{R: 255, G: 0, B: 0, A: 255}), nil)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(50, 0)
	img.DrawImage(generateTargetImage(color.RGBA{R: 255, G: 0, B: 0, A: 255}), opts)
	opts.GeoM.Translate(-20, 7)
	img.DrawImage(generateRightArrow(), opts)
	return img
}

func generateHelpImageFive() *ebiten.Image {
	img := ebiten.NewImage(helpImageW, helpImageH)
	img.Fill(bg)
	img.DrawImage(generateMagnetImage(), nil)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(30, 0)
	img.DrawImage(generateBlockImage(), opts)
	opts.GeoM.Translate(30, 7)
	img.DrawImage(generateRightArrow(), opts)
	return img
}

func generateHelpImageSix() *ebiten.Image {
	img := ebiten.NewImage(helpImageW, helpImageH)
	img.Fill(bg)
	img.DrawImage(generateMagnetImage(), nil)
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(30, 0)
	img.DrawImage(generatePostImage(), opts)
	return img
}
