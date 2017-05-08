package signal

import (
	"testing"
)

func TestNewColorSignal(t *testing.T) {
	colorSignal := NewColorSignal(uint16(100), uint16(5000), uint16(65535))

	expectedColorSignal := ColorSignal{red: uint16(100), green: uint16(5000), blue: uint16(65535)}
	if !colorSignal.Equal(expectedColorSignal) {
		t.Errorf("ColorSignal initialized to wrong value. Expected: %s, got %s", expectedColorSignal.String(), colorSignal.String())
	}
}

func TestRedGreenBlue(t *testing.T) {
	colorSignal := NewColorSignal(0, 5000, 65535)
	red := colorSignal.Red()
	green := colorSignal.Green()
	blue := colorSignal.Blue()

	if red != 0 {
		t.Errorf("Red extracted incorrectly. Expected %d, got %d.", 0, red)
	}
	if green != 5000 {
		t.Errorf("Green extracted incorrectly. Expected %d, got %d.", 0, green)
	}
	if blue != 65535 {
		t.Errorf("Blue extracted incorrectly. Expected %d, got %d.", 0, blue)
	}
}

func TestColorSignalString(t *testing.T) {
	colorSignal := NewColorSignal(100, 5000, 65535)

	description := colorSignal.String()
	if description != "Color signal red: 100, green: 5000, blue: 65535" {
		t.Errorf("Expected description to be: %s, but got %s", "_", description)
	}
}

func (colSig1 ColorSignal) Equal(colSig2 ColorSignal) bool {
	return colSig1.red == colSig2.red && colSig1.green == colSig2.green && colSig1.blue == colSig2.blue
}

func TestColorSignalEqual(t *testing.T) {
	cases := []struct {
		colSig1       ColorSignal
		colSig2       ColorSignal
		expectedEqual bool
	}{
		{NewColorSignal(uint16(100), uint16(5000), uint16(55000)), NewColorSignal(uint16(100), uint16(5000), uint16(55000)), true},
		{NewColorSignal(uint16(101), uint16(5000), uint16(55000)), NewColorSignal(uint16(100), uint16(5000), uint16(55000)), false},
		{NewColorSignal(uint16(100), uint16(5000), uint16(55001)), NewColorSignal(uint16(100), uint16(5000), uint16(55000)), false},
		{NewColorSignal(uint16(0), uint16(0), uint16(0)), NewColorSignal(uint16(100), uint16(5000), uint16(55000)), false},
		{NewColorSignal(uint16(0), uint16(0), uint16(0)), NewColorSignal(uint16(0), uint16(0), uint16(0)), true},
		{NewColorSignal(uint16(0), uint16(1), uint16(0)), NewColorSignal(uint16(0), uint16(0), uint16(0)), false},
	}

	for i, aCase := range cases {
		equal := aCase.colSig1.Equal(aCase.colSig2)
		equalReverse := aCase.colSig2.Equal(aCase.colSig1)

		if equal != equalReverse {
			t.Errorf("Color Signal equality method is not symmetric for case: %d.", i)
		}

		if equal && !aCase.expectedEqual {
			t.Errorf("Expected color signals %d to be not equal, got equal. Signal I: %s, II: %s", i, aCase.colSig1.String(), aCase.colSig2.String())
		}
		if !equal && aCase.expectedEqual {
			t.Errorf("Expected color signals %d to be  equal, got not equal. Signal I: %s, II: %s", i, aCase.colSig1.String(), aCase.colSig2.String())
		}
	}
}
