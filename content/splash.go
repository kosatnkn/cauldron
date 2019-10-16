package content

import (
	"bytes"
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// GenerateSplashStyle generates the content of the splash style file using the project name.
func GenerateSplashStyle(name string, style string) string {

	splash := generateSplashASCII(name, style)

	content := `// The ASCII art is generated using the go-figure package
// https://github.com/common-nighthawk/go-figure
//
// More styles are available at following links.
// - https://www.kammerl.de/ascii/AsciiSignature.php
// - http://www.figlet.org/

package splash

// StyleShadow shadow style
const StyleShadow string = `

	return fmt.Sprintf("%s`\n%s`", content, splash)
}

// generateSplashASCII generates the ASCII art.
func generateSplashASCII(name string, style string) string {

	splash := figure.NewFigure(name, style, true).String()

	// replace backtick (`) with single quote (')
	// so that back tick string notation will not break
	splash = string(bytes.Replace([]byte(splash), []byte("`"), []byte("'"), -1))

	return splash
}
