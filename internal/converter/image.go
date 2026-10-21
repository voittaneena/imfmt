package converter

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// ConvertImage decodes the input and encodes it in the target format.
func ConvertImage(inputPath, outputPath string, quality int, targetFormat string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	encoder, err := GetEncoder(targetFormat)
	if err != nil {
		return err
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	opts := EncodeOptions{Quality: quality}
	if err := encoder.Encode(out, img, opts); err != nil {
		_ = out.Close()
		return err
	}

	return out.Close()
}
