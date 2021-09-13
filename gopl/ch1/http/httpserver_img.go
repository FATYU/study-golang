package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var count int

func main() {

	http.HandleFunc("/img", showImgHandler)
	http.ListenAndServe("localhost:8888", nil)
}

func showImgHandler(w http.ResponseWriter, req *http.Request) {

	var cycles int64
	values := req.URL.Query()
	var cycle = values.Get("cycle")
	if cycle != "" {
		cycles, err := strconv.ParseInt(cycle, 0, 0)
		if err != nil {
			log.Fatal("ERROR:", err)
		}
		lissajous(w, cycles)
	}

	lissajous(w, cycles)
}

var palette = []color.Color{color.White, color.RGBA{0x7a, 0x79, 0xaa, 0xff}} //定义颜色变量 color.RGBA{0x00, 00, 0x00, 0xff}
const (
	whiteIndex = 0
	blackIndex = 1
	otherIndex = 1
)

func lissajous(out io.Writer, cycles int64) {

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	if cycles <= 0 {
		cycles = 5
	}
	freq := rand.Float64() * 3.0
	animation := gif.GIF{LoopCount: nframes}
	phase := 0.0
	//遍历处理每帧的数据
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(2*cycles)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), otherIndex)
		}
		phase += 0.1
		animation.Delay = append(animation.Delay, delay)
		animation.Image = append(animation.Image, img)

	}
	gif.EncodeAll(out, &animation)

}
