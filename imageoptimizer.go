package imageoptimizer

import (
	"fmt"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

type dimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func returnError(err error) error {
	fmt.Println("error: ", err)
	return fmt.Errorf("error:  %v", err)
}

func setDimension(width *int, height *int, aspectRatio float64) (*dimension, error) {
	newDimension := &dimension{}

	if width == nil && height == nil {
		return nil, fmt.Errorf("width and height can be nil")
	}

	if width == nil {
		newDimension.Width = int(float64(*height) * aspectRatio)
		newDimension.Height = *height
	} else {
		newDimension.Height = int(float64(*width) / aspectRatio)
		newDimension.Width = *width
	}

	return newDimension, nil
}

func exportAs(fileType string, image *vips.ImageRef) ([]byte, error) {

	var newImage []byte
	var err error

	switch fileType {
	case "webp":
		newImage, _, err = image.ExportWebp(vips.NewWebpExportParams())
	case "png":
		newImage, _, err = image.ExportPng(vips.NewPngExportParams())
	case "jpeg":
		newImage, _, err = image.ExportJpeg(vips.NewJpegExportParams())
	case "jp2":
		newImage, _, err = image.ExportJp2k(vips.NewJp2kExportParams())
	case "gif":
		newImage, _, err = image.ExportGIF(vips.NewGifExportParams())
	default:
	}

	return newImage, err

}

func getImage(inputFile string) (*vips.ImageRef, error) {
	var image *vips.ImageRef
	var err error

	if strings.HasPrefix(inputFile, "http") {
		resp, _ := http.Get(inputFile)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		image, err = vips.NewImageFromReader(resp.Body)

	} else {
		image, err = vips.NewImageFromFile(inputFile)

	}

	return image, err
}

// Inputfile can be file path or url
// Outputfile should be like /path/filename without the extension name
// Scale: 0.5 for 50%
// filetype can be web, jpeg, png, jp2
func ResizeImage(inputFile string, outputFile string, scale float64, fileType string) error {

	image, err := getImage(inputFile)

	if err != nil {
		return returnError(err)
	}

	defer image.Close()

	err = image.Resize(scale, vips.KernelAuto)

	if err != nil {
		return returnError(err)
	}

	newImage, err := exportAs(fileType, image)

	if err != nil {
		return returnError(err)
	}

	if fileType == "jpeg" {
		fileType = "jpg"
	}

	file, err := os.Create(outputFile + "." + fileType)

	if err != nil {
		return returnError(err)
	}

	defer file.Close()

	_, err = file.Write(newImage)

	if err != nil {
		return returnError(err)
	}

	err = file.Sync()

	if err != nil {
		return returnError(err)
	}

	return nil

}

// Inputfile can be file path or url
// Outputfile should be like /path/filename without the extension name
// aspectRatio is the proportional relationship between width and height and must be in float. For square aspect ratio can be write 1 or 1.0
// targetHeight will be ignore if targetWidth not equal to zero 0
// filetype can be web, jpeg, png, jp2
func ImageCrop(inputFile string, outputFile string, aspectRatio float64, targetWidth int, targetHeight int, fileType string) error {
	image, err := getImage(inputFile)

	if err != nil {
		return returnError(err)
	}

	defer image.Close()

	var newWidth int
	var newHeight int

	if targetWidth == 0 && targetHeight == 0 {
		newWidth = image.Width()
	}

	if targetWidth != 0 {
		newWidth = targetWidth
	}

	if targetWidth == 0 && targetHeight != 0 {
		newHeight = image.Height()
	}

	newDimension, err := setDimension(&newWidth, &newHeight, aspectRatio)

	if err != nil {
		return returnError(err)
	}

	// check if new dimension is greater than original dimension
	if newDimension.Width > image.Width() || newDimension.Height > image.Height() {
		// add/expand white canvas to image
		left := int(math.Abs(float64(newWidth)-float64(image.Width())) / 2)
		top := int(math.Abs(float64(newHeight)-float64(image.Height())) / 2)

		err := image.EmbedBackgroundRGBA(left, top, newDimension.Width, newDimension.Height, &vips.ColorRGBA{R: 255, G: 255, B: 255, A: 0})

		if err != nil {
			return returnError(err)
		}
	}

	err = image.SmartCrop(newDimension.Width, newDimension.Height, vips.InterestingAttention)

	if err != nil {
		return returnError(err)
	}

	newImage, err := exportAs(fileType, image)

	if err != nil {
		return returnError(err)
	}

	if fileType == "jpeg" {
		fileType = "jpg"
	}

	file, err := os.Create(outputFile + "." + fileType)

	if err != nil {
		return returnError(err)
	}

	defer file.Close()

	_, err = file.Write(newImage)

	if err != nil {
		return returnError(err)
	}

	err = file.Sync()

	if err != nil {
		return returnError(err)
	}

	return nil
}

// Inputfile can be file path or url
// Outputfile should be like /path/filename without the extension name
// filetype can be web, jpeg, png, jp2
func CreateThumbnail(inputFile string, outputFile string, width, height int, fileType string) error {
	image, err := getImage(inputFile)

	if err != nil {
		return returnError(err)
	}

	defer image.Close()

	err = image.Thumbnail(width, height, vips.InterestingAttention)

	if err != nil {
		return returnError(err)
	}

	newImage, err := exportAs(fileType, image)

	if err != nil {
		return returnError(err)
	}

	if fileType == "jpeg" {
		fileType = "jpg"
	}

	file, err := os.Create(outputFile + "." + fileType)

	if err != nil {
		return returnError(err)
	}

	defer file.Close()

	_, err = file.Write(newImage)

	if err != nil {
		return returnError(err)
	}

	err = file.Sync()

	if err != nil {
		return returnError(err)
	}

	return nil

}

// Inputfile can be file path or url
// Outputfile should be like /path/filename without the extension name
// filetype can be web, jpeg, png, jp2
func CreateThumbnailWithSize(inputFile string, outputFile string, width, height int, fileType string) error {
	image, err := getImage(inputFile)

	if err != nil {
		return returnError(err)
	}

	defer image.Close()

	err = image.ThumbnailWithSize(width, height, vips.InterestingAttention, vips.SizeBoth)

	if err != nil {
		return returnError(err)
	}

	newImage, err := exportAs(fileType, image)

	if err != nil {
		return returnError(err)
	}

	if fileType == "jpeg" {
		fileType = "jpg"
	}

	file, err := os.Create(outputFile + "." + fileType)

	if err != nil {
		return returnError(err)
	}

	defer file.Close()

	_, err = file.Write(newImage)

	if err != nil {
		return returnError(err)
	}

	err = file.Sync()

	if err != nil {
		return returnError(err)
	}

	return nil

}

func AddWaterMark(inputFile string, watermarkFile string, outputFile string, fileType string) error {
	image, err := getImage(inputFile)

	if err != nil {
		return returnError(err)
	}

	defer image.Close()

	wmarkImage, err := vips.NewImageFromFile(watermarkFile)

	if err != nil {
		return returnError(err)
	}

	// get bottom position
	x := (image.Width() - wmarkImage.Width()) / 2
	y := (image.Height() - wmarkImage.Height()) / 2

	image.Insert(wmarkImage, x, y, false, &vips.ColorRGBA{R: 0, G: 0, B: 0, A: 0})

	newImage, err := exportAs(fileType, image)

	if err != nil {
		return returnError(err)
	}

	if fileType == "jpeg" {
		fileType = "jpg"
	}

	file, err := os.Create(outputFile + "." + fileType)

	if err != nil {
		return returnError(err)
	}

	defer file.Close()

	_, err = file.Write(newImage)

	if err != nil {
		return returnError(err)
	}

	err = file.Sync()

	if err != nil {
		return returnError(err)
	}

	return nil
}
