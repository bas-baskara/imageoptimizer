package imageoptimizer

import (
	"os"
	"testing"
)

func TestResizeImage(t *testing.T) {
	ResizeImage("https://loremflickr.com/cache/resized/65535_52476915038_190d3e82ef_z_600_400_nofilter.jpg", "/home/ayok/Pictures/cat-resize-image", 0.5, "webp")

	_, err := os.Stat("/home/ayok/Pictures/cat-resize-image.webp")

	if err != nil {
		t.Errorf("error: file is not created: %v\n", err)
	}

}

func TestImageCrop(t *testing.T) {
	ImageCrop("https://loremflickr.com/cache/resized/65535_52476915038_190d3e82ef_z_600_400_nofilter.jpg", "/home/ayok/Pictures/cat-image-crop-1.0-100-0", 1.0, 100, 0, "png")

	_, err := os.Stat("/home/ayok/Pictures/cat-image-crop-1.0-100-0.png")

	if err != nil {
		t.Errorf("error: file is not created: %v\n", err)
	}
}

func TestCreateThumbnail(t *testing.T) {
	CreateThumbnail("https://loremflickr.com/cache/resized/65535_52476915038_190d3e82ef_z_600_400_nofilter.jpg", "/home/ayok/Pictures/create-thumb-75-150", 75, 150, "gif")

	_, err := os.Stat("/home/ayok/Pictures/create-thumb-75-150.gif")

	if err != nil {
		t.Errorf("error: file is not created: %v\n", err)
	}
}

func TestCreateThumbnailWithSize(t *testing.T) {
	CreateThumbnailWithSize("https://loremflickr.com/cache/resized/65535_52476915038_190d3e82ef_z_600_400_nofilter.jpg", "/home/ayok/Pictures/create-thumb-withsize-200-35", 200, 35, "webp")

	_, err := os.Stat("/home/ayok/Pictures/create-thumb-withsize-200-35.webp")

	if err != nil {
		t.Errorf("error: file is not created %v\n", err)
	}
}
