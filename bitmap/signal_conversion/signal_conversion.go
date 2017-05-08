package signal_conversion

import (
	"SimpleGraphics/bitmap"
	"SimpleGraphics/signal"
)

func ConvertSignalToPixel(colSig *signal.ColorSignal) *bitmap.Pixel {
	if colSig == nil {
		return nil
	}

	pixel := bitmap.NewPixel(uint8((*colSig).Red()/256), uint8((*colSig).Green()/256), uint8((*colSig).Blue()/256), 255)
	return &pixel
}
