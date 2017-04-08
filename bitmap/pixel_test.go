package bitmap

import (
	"testing"
)

func TestPixel(t *testing.T) {
	const (
		RED   = uint8(10)
		GREEN = uint8(20)
		BLUE  = uint8(30)
		ALPHA = uint8(40)
	)
	pixel := NewPixel(RED, GREEN, BLUE, ALPHA)

	if pixel.Red() != RED {
		t.Errorf("Expected Red to be %d, got %d", RED, pixel.Red())
	}
	if pixel.Green() != GREEN {
		t.Errorf("Expected Green to be %d, got %d", GREEN, pixel.Green())
	}
	if pixel.Blue() != BLUE {
		t.Errorf("Expected Blue to be %d, got %d", BLUE, pixel.Blue())
	}
	if pixel.Alpha() != ALPHA {
		t.Errorf("Expected Alpha to be %d, got %d", ALPHA, pixel.Alpha())
	}
}
