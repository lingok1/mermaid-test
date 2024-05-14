package main

import (
	"bytes"
	"image"
	"image/png"
	"io"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func SVG2PNG(svg io.Reader) (io.Reader, error) {
	if icon, err := oksvg.ReadIconStream(svg); err != nil {
		return nil, err
	} else {
		width, height := int(icon.ViewBox.W), int(icon.ViewBox.H)
		icon.SetTarget(0, 0, icon.ViewBox.W, icon.ViewBox.H)
		rgba := image.NewRGBA(image.Rect(0, 0, width, height))
		icon.Draw(rasterx.NewDasher(width, height, rasterx.NewScannerGV(width, height, rgba, rgba.Bounds())), 1)
		buf := &bytes.Buffer{}
		if err = png.Encode(buf, rgba); err != nil {
			return nil, err
		} else {
			return buf, nil
		}
	}
}
