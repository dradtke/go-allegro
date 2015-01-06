package allegro

import (
	"image"
)

// This file contains tools for making the library more idiomatic by
// doing things like grouping common functionality into interfaces.

// EventGenerator represents anything that can be registered with an
// event queue.
type EventGenerator interface {
	EventSource() *EventSource
}

// ImageToBitmap() converts any image.Image to an Allegro bitmap.
//
// This method is experimental and hasn't been tested nor optimized.
func ImageToBitmap(img image.Image) (*Bitmap, error) {
	var (
		bounds = img.Bounds()
		bmp    = CreateBitmap(bounds.Max.X, bounds.Max.Y)
	)

	err := bmp.WithLockedTarget(bmp.BitmapFormat(), LOCK_WRITEONLY, func() {
		const max = 0xFFFF // the highest color value that RGBA() can return
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()
				col := MapRGBAf(float32(r)/max, float32(g)/max, float32(b)/max, float32(a)/max)
				PutPixel(x, y, col)
			}
		}
	})

	if err != nil {
		bmp.Destroy()
		return nil, err
	}

	return bmp, nil
}
