package content

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// GenerateSplash generates the splash.
func GenerateSplash(name string) string {

	splash := figure.NewFigure(name, "stop", true).String()

	content := `
	// More is available at following link.
	// https://www.kammerl.de/ascii/AsciiSignature.php

	package splash

	// StyleShadow shadow style
	const StyleShadow string = 
	`

	return fmt.Sprintf("%s `\n%s`", content, splash)
}
