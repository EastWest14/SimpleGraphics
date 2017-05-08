package signal

import (
	"fmt"
)

type ColorSignal struct {
	red   uint16
	green uint16
	blue  uint16
}

//NewColorSignal is created with 3 integers in range 0 to 2^16 - 1.
func NewColorSignal(red, green, blue uint16) ColorSignal {
	return ColorSignal{red: red, green: green, blue: blue}
}

func (colSig ColorSignal) Red() uint16 {
	return colSig.red
}

func (colSig ColorSignal) Green() uint16 {
	return colSig.green
}

func (colSig ColorSignal) Blue() uint16 {
	return colSig.blue
}

func (colSig ColorSignal) String() string {
	return fmt.Sprintf("Color signal red: %d, green: %d, blue: %d", colSig.red, colSig.green, colSig.blue)
}
