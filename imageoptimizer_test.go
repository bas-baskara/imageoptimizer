package imageoptimizer

import (
	"os"
	"testing"
)

func TestResizeImage(t *testing.T) {
	ResizeImage("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg", "/home/ayok/Pictures/istockphoto-172759822-170667a.jpg", 0.5, "png")

	_, err := os.Stat("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg.png")

	if err != nil {
		t.Errorf("error: file is not created: %v\n", err)
	}

}

func TestImageCrop(t *testing.T) {
	ImageCrop("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg", "/home/ayok/Pictures/istockphoto-172759822-170667a.jpg-1.0-100-0", 1.0, 100, 0, "png")

	_, err := os.Stat("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg-1.0-100-0.png")

	if err != nil {
		t.Errorf("error: file is not created: %v\n", err)
	}
}

func TestCreateThumbnail(t *testing.T) {
	CreateThumbnail("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg", "/home/ayok/Pictures/istockphoto-172759822-170667a.jpg", 75, 150, "gif")

	_, err := os.Stat("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg.gif")

	if err != nil {
		t.Errorf("error: file is not created: %v\n", err)
	}
}

func TestCreateThumbnailWithSize(t *testing.T) {
	CreateThumbnailWithSize("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg", "/home/ayok/Pictures/istockphoto-172759822-170667a.jpg-200-35", 200, 35, "webp")

	_, err := os.Stat("/home/ayok/Pictures/istockphoto-172759822-170667a.jpg-200-35.webp")

	if err != nil {
		t.Errorf("error: file is not created %v\n", err)
	}
}
