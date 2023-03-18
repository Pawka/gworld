package main

import (
	"fmt"
	"time"

	"github.com/fogleman/gg"
	opensimplex "github.com/ojrac/opensimplex-go"
	"golang.org/x/exp/rand"
)

type style int

const (
	StyleSharp style = iota + 1
	StyleBlended
	StyleMap
)

func main() {

	var seed int64 = time.Now().Unix()

	const (
		S               = 1000
		edge    float64 = 0.7
		scale   float64 = 200
		style           = StyleMap
		debug           = false
		stretch         = 1.1
	)

	rand.Seed(uint64(seed))
	noise := opensimplex.New(seed)
	var i, j float64
	dc := gg.NewContext(S, S)
	for i = 0; i < S; i++ {
		for j = 0; j < S; j++ {
			var _n float64

			_n = normalize(noise.Eval2(i/scale, j/scale/stretch))
			if style == StyleSharp {
				if _n > edge && _n >= rand.Float64() {
					dc.SetRGB255(0, 0, 0)
				} else {
					dc.SetRGB255(255, 255, 255)
				}
			} else if style == StyleBlended {
				dc.SetRGB(_n, _n, _n)
			} else if style == StyleMap {
				setColor(dc, _n)
			}
			if debug && i < 3 && j < 3 {
				fmt.Println(_n)
			}

			dc.DrawPoint(i, j, 1)
			dc.Fill()
		}
	}

	dc.SavePNG("out.png")
}

// setColor sets a color based on normalized value nval. nval is between 0 to 1.
func setColor(ctx *gg.Context, nval float64) {
	type color struct {
		r, g, b int
	}

	colors := map[string]color{
		"deep_water": {6, 57, 112},
		"water":      {30, 129, 176},
		"sand":       {234, 182, 118},
		"grass":      {129, 185, 52},
		"forest":     {53, 111, 70},
	}

	c := colors["forest"]
	switch true {
	case nval < 0.1:
		c = colors["deep_water"]
	case nval < 0.3:
		c = colors["water"]
	case nval < 0.4:
		c = colors["sand"]
	case nval < 0.7:
		c = colors["grass"]
	}

	ctx.SetRGB255(c.r, c.g, c.b)
}

func normalize(val float64) float64 {
	return (val + 1.0) / 2
}
