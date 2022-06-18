package main

import (
	"embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var regular *font.Face

//go:embed assets
var gd embed.FS

func init() {
	ttbytes, err := gd.ReadFile("assets/Righteous-Regular.ttf")
	if err == nil {
		tt, err := opentype.Parse(ttbytes)
		if err == nil {
			fontface, err := opentype.NewFace(tt, &opentype.FaceOptions{
				Size:    24,
				DPI:     72,
				Hinting: font.HintingFull,
			})
			if err == nil {
				regular = &fontface
			} else {
				panic(err)
			}

		} else {
			panic(err)
		}
	} else {
		panic(err)

	}

}
