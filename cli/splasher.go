package cli

import "github.com/kosatnkn/cauldron/log"

// ShowSplash shows a splash message.
//
// More is available at following link.
// https://www.kammerl.de/ascii/AsciiSignature.php
func ShowSplash() {

	splash := `
	  ___|             |     |                
	 |      _' | |   | |  _' |  __| _ \  __ \ 
	 |     (   | |   | | (   | |   (   | |   |
	\____|\__,_|\__,_|_|\__,_|_|  \___/ _|  _|								 
	`

	caption := "REST API project generator using `Catalyst` as a base"

	log.Default(splash)
	log.Default(caption)
	log.Default("")
}
