package converter

import (
	"image"
	"image/jpeg"
	"io"
)

// JPEGEncoder implements the Strategy pattern for JPEG.
type JPEGEncoder struct{}

func (e *JPEGEncoder) Encode(w io.Writer, img image.Image, opts EncodeOptions) error {
	options := &jpeg.Options{
		Quality: opts.Quality,
	}
	return jpeg.Encode(w, img, options)
}
