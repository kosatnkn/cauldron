package cli

import (
	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/log"
)

// ShowSplash shows a splash message.
//
// More is available at following link.
// https://www.kammerl.de/ascii/AsciiSignature.php
func ShowSplash(cfg *config.Config) {

	splash := `
  ___           _    _              
 / __|__ _ _  _| |__| |_ _ ___ _ _  
| (__/ _' | || | / _' | '_/ _ \ ' \ 
 \___\__,_|\_,_|_\__,_|_| \___/_||_|`

	log.Default(splash)
	log.Default("  " + cfg.Cauldron.Version)
	log.Default("")
}
