package main

import (
	"fmt"
	"path"
	"testing"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/vcaesar/gcv"
)

func findInDesktop(root string, imagePath string) (*gcv.Result, error) {
	robotgo.SaveCapture(path.Join(root, "tmp_desktop.png"))
	results := gcv.FindAllImgFile(path.Join(root, imagePath), path.Join(root, "tmp_desktop.png"))
	if len(results) == 0 {
		return nil, fmt.Errorf("image %v not found in desktop screenshot", imagePath)
	}
	return &results[0], nil
}

func mustFindInDesktop(root string, imagePath string) *gcv.Result {
	result, err := findInDesktop(root, imagePath)
	if err != nil {
		panic(err)
	}
	return result
}

func clickGCV(result *gcv.Result) {
	robotgo.MoveClick(result.Middle.X, result.Middle.Y)
}

func TestGreeting(t *testing.T) {
	root := "test/TestGreeting"

	// Click the input
	input := mustFindInDesktop(root, "001_input.png")
	clickGCV(input)

	// Type our name
	robotgo.TypeStr("AlbinoDrought")

	// Click the greeting button
	button := mustFindInDesktop(root, "002_button.png")
	clickGCV(button)

	// Must see our greetings text
	time.Sleep(1 * time.Second)
	mustFindInDesktop(root, "003_greetings.png")
}
