package converter

import (
	"image"
	"image/png"
	"io"
)

// PNGEncoder implements the Strategy pattern for PNG.
type PNGEncoder struct{}

func (e *PNGEncoder) Encode(w io.Writer, img image.Image, opts EncodeOptions) error {
	encoder := &png.Encoder{
		CompressionLevel: png.DefaultCompression,
	}

	if opts.Quality >= 90 {
		encoder.CompressionLevel = png.BestCompression
	} else if opts.Quality <= 50 {
		encoder.CompressionLevel = png.BestSpeed
	}

	return encoder.Encode(w, img)
}
