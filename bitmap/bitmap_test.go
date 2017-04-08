package bitmap

import (
	"testing"
)

func TestNewBitmap(t *testing.T) {
	/*bm := NewBitmap()
	if bm == nil {
		t.Errorf("Failed to initialize Bitmap.")
	}*/
}

func TestPixelAtSetGet(t *testing.T) {
	bm := NewBitmap()
	p := NewPixel(uint8(100), uint8(0), uint8(0), uint8(0))

	bm.SetPixelAt(100, 100, p)
	pReturned := bm.PixelAt(100, 100)

	if pReturned.Red() != uint8(100) {
		t.Errorf("Pixel returned incorrectly.")
	}

	pNotSet := bm.PixelAt(12, 34)
	if pNotSet.Red() != uint8(0) || pNotSet.Green() != uint8(0) || pNotSet.Blue() != uint8(0) || pNotSet.Blue() != uint8(0) {
		t.Errorf("Expected an unset pixel to be 0 for all colors and alpha, got: %s.", pNotSet.String())
	}
}
