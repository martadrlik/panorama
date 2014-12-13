// command panorama creates picture from pictures given by argument
package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

var (
	output = flag.String("output", "output.png", "Place output into this")
)

func main() {
	flag.Parse()
	a := make([]image.Image, flag.NArg())
	var w, h int
	for i := 0; i < flag.NArg(); i++ {
		r, err := os.Open(flag.Arg(i))
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		m, _, err := image.Decode(r)
		if err != nil {
			log.Fatal(err)
		}
		w += m.Bounds().Max.X
		if h < m.Bounds().Max.Y {
			h = m.Bounds().Max.Y
		}
		a[i] = m
	}
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	var offset int
	for _, n := range a {
		b := n.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				m.Set(offset+x, y, n.At(x, y))
			}
		}
		offset += b.Max.X
	}
	o, err := os.Create(*output)
	if err != nil {
		log.Fatal(err)
	}
	defer o.Close()
	err = png.Encode(o, m)
	if err != nil {
		log.Fatal(err)
	}
}
