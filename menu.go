package main

import (
	"github.com/fogleman/gg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type menuItem struct {
	name           string
	completed      bool
	flagGenerator  func(int, int, int) *ebiten.Image
	levelGenerator func(int) Level
}

var flags [3]*menuItem
var menuSelected int

func init() {
	flags = [3]*menuItem{
		{name: "rainbow",
			levelGenerator: getRainbowLevel,
			flagGenerator:  generateRainbowFlag,
		},
		{name: "rainbow",
			levelGenerator: getRainbowLevel,
			flagGenerator:  generateRainbowFlag,
		},
		{name: "rainbow",
			levelGenerator: getRainbowLevel,
			flagGenerator:  generateRainbowFlag,
		},
	}

	menuSelected = 0
}
func generateMenuImage() *ebiten.Image {

	var menuWidth int
	endgame := true

	for _, flag := range flags {
		if !flag.completed {
			endgame = false
		}
		if mw := text.BoundString(*regular, flag.name).Dx(); mw > menuWidth {
			menuWidth = mw
		}
	}

	dc := gg.NewContext(w, h)
	dc.SetColor(fg)
	dc.DrawLine(float64(10+w/2+menuWidth/2),
		float64(h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()),
		float64(10+w/2+menuWidth/2),
		float64(float64(h*(len(flags)-1)/(len(flags)+1)+h/(len(flags)+1))),
	)
	dc.Stroke()

	var menu = ebiten.NewImage(w, h)
	menu.Fill(bg)
	if endgame {
		message := "thank you for playing"
		text.Draw(menu, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, 10+text.BoundString(*regular, message).Dy(), fg)
	} else {
		menu.DrawImage(ebiten.NewImageFromImage(dc.Image()), nil)
	}
	for i, flag := range flags {
		if flag.completed {
			flagWidth := 100
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(w/2-100/2), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()/2-flagWidth*3/10))
			menu.DrawImage(flag.flagGenerator(flagWidth, flagWidth*3/5, 100), opts)
		} else {
			text.Draw(menu, flag.name, *regular, w/2-text.BoundString(*regular, flag.name).Dx()/2, h*i/(len(flags)+1)+h/(len(flags)+1), fg)
		}

		if i == menuSelected && !flags[menuSelected].completed && !endgame {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(20+w/2+menuWidth/2), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()))
			menu.DrawImage(generateRightArrow(), opts)
		}
		if i == menuSelected+1 && !endgame {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(20+w/2+menuWidth/2), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()))
			menu.DrawImage(generateDownArrow(), opts)
		}
		if i == menuSelected-1 && !endgame {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(20+w/2+menuWidth/2), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()))
			menu.DrawImage(generateUpArrow(), opts)
		}
	}

	return menu
}

func generateRightArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(0, 0)
	dc.LineTo(float64(textHeight), float64(textHeight/2))
	dc.LineTo(0, float64(textHeight))
	dc.LineTo(0, 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateDownArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(0, 0)
	dc.LineTo(float64(textHeight), 0)
	dc.LineTo(float64(textHeight/2), float64(textHeight))
	dc.LineTo(0, 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}

func generateUpArrow() *ebiten.Image {
	var textHeight = text.BoundString(*regular, "X").Dy()

	dc := gg.NewContext(textHeight, textHeight)

	dc.MoveTo(float64(textHeight/2), 0)
	dc.LineTo(float64(textHeight), float64(textHeight))
	dc.LineTo(0, float64(textHeight))
	dc.LineTo(float64(textHeight/2), 0)
	dc.SetColor(fg)
	dc.Fill()
	return ebiten.NewImageFromImage(dc.Image())
}
