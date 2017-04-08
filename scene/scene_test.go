package scene

import (
	"testing"
)

func TestNewScene(t *testing.T) {
	sc := NewScene()
	if sc == nil {
		t.Errorf("Failed to initialize scene")
	}
}

func TestSceneAtIndex(t *testing.T) {
	sc := NewScene()

	_, err := sc.ColorAt(256, 255)
	if err == nil {
		t.Errorf("Expected out of bounds error on first index, got no error.")
	}

	_, err = sc.ColorAt(255, 256)
	if err == nil {
		t.Errorf("Expected out of bounds error on second index, got no error.")
	}

	_, err = sc.ColorAt(-1, 0)
	if err == nil {
		t.Errorf("Expected out of bounds error on first index, got no error.")
	}

	_, err = sc.ColorAt(10, -1)
	if err == nil {
		t.Errorf("Expected out of bounds error on second index, got no error.")
	}

	_, err = sc.ColorAt(0, 0)
	if err != nil {
		t.Errorf("Expected no error, got error: %s.", err.Error())
	}
}
