package main

import (
	"fmt"
	"time"

	"github.com/fogleman/gg"
	opensimplex "github.com/ojrac/opensimplex-go"
	"golang.org/x/exp/rand"
)

func main() {

	var seed int64 = time.Now().Unix()

	const (
		S               = 1000
		edge    float64 = 0.7
		scale   float64 = 160
		sharp           = false
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
			if sharp {
				if _n > edge && _n >= rand.Float64() {
					dc.SetRGB255(0, 0, 0)
				} else {
					dc.SetRGB255(255, 255, 255)
				}
			} else {
				dc.SetRGB(_n, _n, _n)
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

func normalize(val float64) float64 {
	return (val + 1.0) / 2
}
