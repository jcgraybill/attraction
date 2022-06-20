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

var flags [4]menuItem
var menuSelected int

func init() {
	flags = [4]menuItem{
		{name: "rainbow",
			levelGenerator: getRainbowLevel,
		},
		{name: "trans",
			levelGenerator: getRainbowLevel,
		},
		{name: "nonbinary",
			levelGenerator: getRainbowLevel,
		},
		{name: "genderqueer",
			levelGenerator: getRainbowLevel,
		},
	}

	menuSelected = 0
}
func generateMenuImage() *ebiten.Image {

	var menuWidth int
	for _, flag := range flags {
		if w := text.BoundString(*regular, flag.name).Dx(); w > menuWidth {
			menuWidth = w
		}
	}

	var menu = ebiten.NewImage(w, h)
	menu.Fill(bg)

	for i, flag := range flags {
		text.Draw(menu, flag.name, *regular, w/2-text.BoundString(*regular, flag.name).Dx()/2, h*i/(len(flags)+1)+h/(len(flags)+1), fg)

		if i == menuSelected {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(10+w/2+menuWidth/2), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()))
			menu.DrawImage(generateRightArrow(), opts)
		}
		if i == menuSelected+1 {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(-10+w/2-menuWidth/2-text.BoundString(*regular, "X").Dy()), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()))
			menu.DrawImage(generateDownArrow(), opts)
		}
		if i == menuSelected-1 {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(-10+w/2-menuWidth/2-text.BoundString(*regular, "X").Dy()), float64(h*i/(len(flags)+1)+h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()))
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
