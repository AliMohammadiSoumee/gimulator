package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"os"
	"image/png"
	"github.com/alidadar7676/gimulator/GUI"
)


func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Paper Soccer",
		Bounds: pixel.R(100, 100, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Green)
	win.SetSmooth(true)

	imd := GUI.CreateGame(GUI.World{})

	for !win.Closed() {
		win.Clear(colornames.Green)

		imd.Line(5)
		imd.Draw(win)
		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
