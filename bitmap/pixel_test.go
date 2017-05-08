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

func TestPixelEqual(t *testing.T) {
	cases := []struct {
		p1                        Pixel
		p2                        Pixel
		expectedEqual             bool
		expectedInequalityMessage string
	}{
		//Equal cases
		{NewPixel(0, 0, 0, 0), NewPixel(0, 0, 0, 0), true, ""},
		{NewPixel(255, 255, 255, 255), NewPixel(255, 255, 255, 255), true, ""},
		{NewPixel(100, 150, 200, 250), NewPixel(100, 150, 200, 250), true, ""},
		//Not equal
		{NewPixel(0, 0, 0, 0), NewPixel(1, 0, 0, 0), false, "Pixel 1 not equal to Pixel 2. P1: Red: 0, Green: 0, Blue: 0, Alpha: 0, P2: Red: 1, Green: 0, Blue: 0, Alpha: 0"},
		{NewPixel(0, 0, 0, 0), NewPixel(0, 1, 0, 0), false, "Pixel 1 not equal to Pixel 2. P1: Red: 0, Green: 0, Blue: 0, Alpha: 0, P2: Red: 0, Green: 1, Blue: 0, Alpha: 0"},
		{NewPixel(0, 0, 0, 0), NewPixel(0, 0, 1, 0), false, "Pixel 1 not equal to Pixel 2. P1: Red: 0, Green: 0, Blue: 0, Alpha: 0, P2: Red: 0, Green: 0, Blue: 1, Alpha: 0"},
		{NewPixel(0, 0, 0, 0), NewPixel(0, 0, 0, 1), false, "Pixel 1 not equal to Pixel 2. P1: Red: 0, Green: 0, Blue: 0, Alpha: 0, P2: Red: 0, Green: 0, Blue: 0, Alpha: 1"},
		{NewPixel(0, 0, 0, 0), NewPixel(255, 255, 255, 255), false, "Pixel 1 not equal to Pixel 2. P1: Red: 0, Green: 0, Blue: 0, Alpha: 0, P2: Red: 255, Green: 255, Blue: 255, Alpha: 255"},
	}

	for i, aCase := range cases {
		equal, inequalityMessage := aCase.p1.Equal(aCase.p2)

		if aCase.expectedEqual && !equal {
			t.Errorf("Error in case %d. Exepected equal, got not equal", i)
		}
		if !aCase.expectedEqual && equal {
			t.Errorf("Error in case %d. Exepected not equal, got equal", i)
		}
		if inequalityMessage != aCase.expectedInequalityMessage {
			t.Errorf("Error in case %d. Exepected inequality message %s, got %s", i, aCase.expectedInequalityMessage, inequalityMessage)
		}
	}
}
