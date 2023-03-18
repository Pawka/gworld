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
		S       = 1000
		edge    = 0.7
		zoom    = 180
		style   = StyleMap
		debug   = false
		stretch = 1
	)

	fmt.Printf("Seed: %d\n", seed)

	rand.Seed(uint64(seed))
	noise := opensimplex.New(seed)
	var i, j float64
	dc := gg.NewContext(S, S)
	for i = 0; i < S; i++ {
		for j = 0; j < S; j++ {
			var _n float64

			_n = normalize(noise.Eval2(i/zoom, j/zoom/stretch))
			if style == StyleSharp {
				if _n > edge { //} && _n >= rand.Float64() {
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
		"deep_water": {92, 160, 180},
		"water":      {131, 194, 213},
		"sand":       {233, 204, 119},
		"grass":      {141, 191, 113},
		"forest":     {116, 168, 119},
		"rocks":      {178, 172, 150},
		"snow":       {255, 255, 255},
	}

	c := colors["snow"]
	switch true {
	case nval < 0.1:
		c = colors["deep_water"]
	case nval < 0.3:
		c = colors["water"]
	case nval < 0.4:
		c = colors["sand"]
	case nval < 0.6:
		c = colors["grass"]
	case nval < 0.85:
		c = colors["forest"]
	case nval < 0.9:
		c = colors["rocks"]
	}

	ctx.SetRGB255(c.r, c.g, c.b)
}

func normalize(val float64) float64 {
	return (val + 1.0) / 2
}
