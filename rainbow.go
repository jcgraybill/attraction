package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var rainbowColors [6]color.RGBA

func init() {
	rainbowColors = [6]color.RGBA{
		{R: 252, G: 0, B: 6, A: 255},
		{R: 253, G: 165, B: 7, A: 255},
		{R: 255, G: 255, B: 11, A: 255},
		{R: 17, G: 131, B: 1, A: 255},
		{R: 0, G: 0, B: 255, A: 255},
		{R: 117, G: 7, B: 135, A: 255},
	}
}

// TODO: take w,h,segments parameters & just generate the flag itself

func generateRainbowFlag(segments int) *ebiten.Image {

	if segments > len(rainbowColors) {
		segments = len(rainbowColors)
	}

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

func getRainbowLevel(which int) Level {

	rainbowLevels := [4]Level{
		{
			label: "level 1: life",
			moves: 2,
			cells: 4,
			flag:  generateRainbowFlag(1),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      2,
					y:      1,
					tile:   true,
					color:  rainbowColors[0],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
					x:        1,
					y:        2,
					moveable: true,
					color:    rainbowColors[0],
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        0,
					y:        2,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        1,
					y:        3,
					moveable: true,
					magnetic: true,
				},
			},
		},
		{
			label: "level 2: healing",
			moves: 6,
			cells: 4,
			flag:  generateRainbowFlag(2),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      2,
					y:      1,
					tile:   true,
					color:  color.RGBA{R: 253, G: 165, B: 7, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[1].color) },
					x:        1,
					y:        2,
					moveable: true,
					color:    color.RGBA{R: 253, G: 165, B: 7, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        0,
					y:        2,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      2,
				},
			},
		},

		{
			label: "level 3: sunlight and nature",
			moves: 6,
			cells: 4,
			flag:  generateRainbowFlag(4),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      1,
					y:      0,
					tile:   true,
					color:  color.RGBA{R: 255, G: 255, B: 11, A: 255},
				},
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[1].color) },
					x:      0,
					y:      0,
					tile:   true,
					color:  color.RGBA{R: 17, G: 131, B: 1, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[2].color) },
					x:        1,
					y:        1,
					moveable: true,
					color:    color.RGBA{R: 255, G: 255, B: 11, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
					x:        0,
					y:        1,
					moveable: true,
					color:    color.RGBA{R: 17, G: 131, B: 1, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        3,
					y:        0,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        3,
					y:        1,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateBlockImage() },
					x:        2,
					y:        2,
					moveable: true,
				},
			},
		},
		{
			label: "level 4: harmony and spirit",
			moves: 11,
			cells: 4,
			flag:  generateRainbowFlag(6),
			pieces: []*Piece{
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[0].color) },
					x:      3,
					y:      2,
					tile:   true,
					color:  color.RGBA{R: 0, G: 0, B: 255, A: 255},
				},
				{
					imgsrc: func() *ebiten.Image { return generateTargetImage(level.pieces[1].color) },
					x:      2,
					y:      1,
					tile:   true,
					color:  color.RGBA{R: 117, G: 7, B: 135, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[2].color) },
					x:        1,
					y:        2,
					moveable: true,
					color:    color.RGBA{R: 0, G: 0, B: 255, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateGemImage(level.pieces[3].color) },
					x:        1,
					y:        1,
					moveable: true,
					color:    color.RGBA{R: 117, G: 7, B: 135, A: 255},
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        0,
					y:        2,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateMagnetImage() },
					x:        2,
					y:        3,
					moveable: true,
					magnetic: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateBlockImage() },
					x:        2,
					y:        2,
					moveable: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateBlockImage() },
					x:        0,
					y:        1,
					moveable: true,
				},
				{
					imgsrc:   func() *ebiten.Image { return generateBlockImage() },
					x:        3,
					y:        1,
					moveable: true,
				},
				{
					imgsrc: func() *ebiten.Image { return generatePostImage() },
					x:      2,
					y:      0,
				},
			},
		},
	}

	if which > len(rainbowLevels)-1 {
		return last
	}
	return rainbowLevels[which]
}
