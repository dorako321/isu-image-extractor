package Binary

import (
	"testing"
	"github.com/dorako321/isu-image-extractor/modules/Binary"
)

func TestGetExtensionName(t *testing.T) {
	// jpgが返ること
	jpg := Binary.GetExtensionName([]byte{255, 216, 0, 0, 0, 0, 0, 0, 0})
	if jpg != ".jpg" {
		t.Fatalf("failed test %#v", jpg)
	}
	// pngが返ること
	png := Binary.GetExtensionName([]byte{137, 80, 78, 71, 13, 10, 26, 10})
	if png != ".png" {
		t.Fatalf("failed test %#v", png)
	}
	// gifが返ること
	gif := Binary.GetExtensionName([]byte{71, 73, 70, 0, 0, 0, 0, 0, 0, 0})
	if gif != ".gif" {
		t.Fatalf("failed test %#v", gif)
	}

}
