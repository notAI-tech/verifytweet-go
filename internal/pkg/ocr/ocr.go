package ocr

import (
	"github.com/otiai10/gosseract"
	"gopkg.in/gographics/imagick.v3/imagick"
)

// Rescale ...
func Rescale(image []byte) ([]byte, error) {
	wand := imagick.NewMagickWand()
	err := wand.ReadImageBlob(image)
	if err != nil {
		return nil, err
	}
	err = wand.ResampleImage(300, 300, imagick.FILTER_LAGRANGE)
	if err != nil {
		return nil, err
	}
	err = wand.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	if err != nil {
		return nil, err
	}
	err = wand.SetImageColorspace(imagick.COLORSPACE_GRAY)
	if err != nil {
		return nil, err
	}
	err = wand.BlurImage(1, 65000)
	if err != nil {
		return nil, err
	}
	err = wand.NormalizeImage()
	if err != nil {
		return nil, err
	}
	convertedImage := wand.GetImageBlob()
	return convertedImage, nil
}

// ConvertToText ...
func ConvertToText(image []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()
	err := client.SetImageFromBytes(image)
	if err != nil {
		return "", err
	}
	text, err := client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
