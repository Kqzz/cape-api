package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"

	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func getCapeImg(username string) (image.Image, error) {
	resp, err := http.Get(fmt.Sprintf("http://s.optifine.net/capes/%v.png", username))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		img, _ := png.Decode(resp.Body)
		return img, nil
	} else {
		return nil, fmt.Errorf("cape not found on user %v", username)
	}

}

func cropCape(cape image.Image) (image.Image, bool, error) {
	var config cutter.Config
	var customCape bool
	bounds := cape.Bounds()
	if bounds.Max.X == 46 {
		config = cutter.Config{
			Width:  10,
			Height: 16,
			Anchor: image.Point{1, 1},
			Mode:   cutter.TopLeft,
		}
		customCape = false
	} else if bounds.Max.X == 92 {
		config = cutter.Config{
			Width:  20,
			Height: 32,
			Anchor: image.Point{2, 2},
			Mode:   cutter.TopLeft,
		}
		customCape = true
	} else {
		log.Fatalf("unknown image width%v", bounds.Max.X)
	}
	croppedImg, err := cutter.Crop(cape, config)
	if err != nil {
		log.Fatal(err)
	}
	return croppedImg, customCape, nil
}

func scaleCape(cape image.Image, scale int, customCape bool) (image.Image, error) {
	var xSize uint = uint(cape.Bounds().Max.X)
	scaleTo := xSize * uint(scale)
	if !customCape {
		scaleTo *= 2
	}
	m := resize.Resize(scaleTo, 0, cape, resize.NearestNeighbor)
	return m, nil
}

func getCapeBytes(username string, scale int) ([]byte, error) {

	originalImage, err := getCapeImg(username)
	if err != nil {
		return []byte{}, err
	}

	croppedCape, isCustomCape, err := cropCape(originalImage)
	if err != nil {
		return []byte{}, err
	}

	scaledCape, err := scaleCape(croppedCape, scale, isCustomCape)

	if err != nil {
		return []byte{}, err
	}

	buf := new(bytes.Buffer)
	err = png.Encode(buf, scaledCape)
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil

}

// func main() {
// 	fmt.Println(grabCapeBytes("Elektra", 10))
// }
