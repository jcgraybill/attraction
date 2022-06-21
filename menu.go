package main

// TODO: it might be more clear if the up/down arrows are
// on the ends of the line rather than next to the flag
// names. Up/down would need to skip any completed levels.

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
		{name: "trans",
			levelGenerator: getTransLevel,
			flagGenerator:  generateTransFlag,
		},
		{name: "nonbinary",
			levelGenerator: getNonbinaryLevel,
			flagGenerator:  generateNonbinaryFlag,
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
	var menu = ebiten.NewImage(w, h)
	menu.Fill(bg)

	if endgame {
		message := "thank you for playing"
		text.Draw(menu, message, *regular, w/2-text.BoundString(*regular, message).Dx()/2, 10+text.BoundString(*regular, message).Dy(), fg)
	} else {
		dc := gg.NewContext(w, h)
		dc.SetColor(fg)
		dc.DrawLine(float64(10+w/2+menuWidth/2),
			float64(h/(len(flags)+1)-text.BoundString(*regular, "X").Dy()),
			float64(10+w/2+menuWidth/2),
			float64(float64(h*(len(flags)-1)/(len(flags)+1)+h/(len(flags)+1))),
		)
		dc.Stroke()
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
