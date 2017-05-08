package signal_conversion

import (
	"SimpleGraphics/bitmap"
	"SimpleGraphics/signal"
	"testing"
)

func TestConvertSignalToPixel(t *testing.T) {
	//zeroPixel := bitmap.NewPixel(0, 0, 0, 255)
	fullPixel := bitmap.NewPixel(255, 255, 255, 255)

	cases := []struct {
		signal        signal.ColorSignal
		expectedPixel *bitmap.Pixel
	}{
		//{signal: signal.NewColorSignal(0, 0, 0), expectedPixel: &zeroPixel},
		{signal: signal.NewColorSignal(65535, 65535, 65535), expectedPixel: &fullPixel},
	}

	for i, aCase := range cases {
		signal := aCase.signal
		pixel := ConvertSignalToPixel(&signal)

		if pixel == nil {
			if aCase.expectedPixel != nil {
				t.Errorf("Error in case %d. Expected non-nil Pixel, got nil.", i)
				continue
			}
		}

		if aCase.expectedPixel == nil {
			t.Errorf("Error in case %d. Expected nil, got not-nil Pixel", i)
			continue
		}

		equal, inequalityMessage := (*pixel).Equal(*aCase.expectedPixel)
		if !equal {
			t.Errorf("Error in case %d. Pixels not equal: %s", i, inequalityMessage)
		}
	}
}
