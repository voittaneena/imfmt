package converter

import (
	"fmt"
	"image"
	"io"
)

// EncodeOptions holds shared settings for all image encoders.
type EncodeOptions struct {
	Quality int
}

// Encoder is the Strategy interface for saving images.
type Encoder interface {
	Encode(w io.Writer, img image.Image, opts EncodeOptions) error
}

// GetEncoder acts as a Factory to resolve the correct Strategy based on format.
func GetEncoder(format string) (Encoder, error) {
	switch format {
	case "jpg", "jpeg":
		return &JPEGEncoder{}, nil
	case "png":
		return &PNGEncoder{}, nil
	default:
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}
}
