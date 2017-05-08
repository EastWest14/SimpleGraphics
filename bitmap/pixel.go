package bitmap

import (
	"fmt"
)

type Pixel struct {
	red   uint8
	green uint8
	blue  uint8
	alpha uint8
}

func NewPixel(red, green, blue, alpha uint8) Pixel {
	return Pixel{red: red, green: green, blue: blue, alpha: alpha}
}

func (p *Pixel) String() string {
	return fmt.Sprintf("Red: %d, Green: %d, Blue: %d, Alpha: %d", p.Red(), p.Green(), p.Blue(), p.Alpha())
}

func (p Pixel) Red() uint8 {
	return p.red
}

func (p Pixel) Green() uint8 {
	return p.green
}

func (p Pixel) Blue() uint8 {
	return p.blue
}

func (p Pixel) Alpha() uint8 {
	return p.alpha
}

func (p1 Pixel) Equal(p2 Pixel) (equal bool, inequalityMessage string) {
	p1Description := p1.String()
	p2Description := p2.String()

	if p1.Red() != p2.Red() || p1.Green() != p2.Green() || p1.Blue() != p2.Blue() || p1.Alpha() != p2.Alpha() {
		return false, fmt.Sprintf("Pixel 1 not equal to Pixel 2. P1: %s, P2: %s", p1Description, p2Description)
	}

	return true, ""
}
