package main

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"os"

	bimg "gopkg.in/h2non/bimg.v1"
)

func main() {
	join4Img()
}

func join4Img() {
	buffer1, err := bimg.Read("img.jpeg")
	buffer2, err := bimg.Read("img.jpeg")
	buffer3, err := bimg.Read("img.jpeg")
	buffer4, err := bimg.Read("img.jpeg")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	resized1, err := bimg.NewImage(buffer1).SmartCrop(96, 96)
	resized2, err := bimg.NewImage(buffer2).SmartCrop(96, 96)
	resized3, err := bimg.NewImage(buffer3).SmartCrop(96, 96)
	resized4, err := bimg.NewImage(buffer4).SmartCrop(96, 96)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	img1, _, err := image.Decode(bytes.NewReader(resized1))
	img2, _, err := image.Decode(bytes.NewReader(resized2))
	img3, _, err := image.Decode(bytes.NewReader(resized3))
	img4, _, err := image.Decode(bytes.NewReader(resized4))

	//starting position of the second image (bottom left)
	sp2 := image.Point{img1.Bounds().Dx(), 0}
	sp3 := image.Point{0, img3.Bounds().Dy()}
	sp4 := image.Point{img3.Bounds().Dx(), img3.Bounds().Dy()}

	//new rectangle for the second image
	r2 := image.Rectangle{sp2, sp2.Add(img2.Bounds().Size())}
	r3 := image.Rectangle{sp3, sp3.Add(img3.Bounds().Size())}
	r4 := image.Rectangle{sp4, sp4.Add(img4.Bounds().Size())}

	//rectangle for the big image
	r := image.Rectangle{image.Point{0, 0}, r4.Max}

	rgba := image.NewRGBA(r)

	draw.Draw(rgba, img1.Bounds(), img1, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r2, img2, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r3, img3, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, r4, img4, image.Point{0, 0}, draw.Src)

	out, err := os.Create("./bisa.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	var opt jpeg.Options
	opt.Quality = 95

	jpeg.Encode(out, rgba, &opt)
}
