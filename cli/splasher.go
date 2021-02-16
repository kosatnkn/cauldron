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
	 _____          __   __            
	/ ___/__ ___ __/ /__/ /______  ___ 
   / /__/ _ '/ // / / _  / __/ _ \/ _ \
   \___/\_,_/\_,_/_/\_,_/_/  \___/_//_/`

	splash = `
    ___           _    _              
   / __|__ _ _  _| |__| |_ _ ___ _ _  
  | (__/ _' | || | / _' | '_/ _ \ ' \ 
   \___\__,_|\_,_|_\__,_|_| \___/_||_|`

	splash = `
    __|            |     |            
   (     _' | |  | |  _' |  _|_ \   \ 
  \___|\__,_|\_,_|_|\__,_|_|\___/_| _|`

	log.Default(splash)
	log.Default(cfg.Cauldron.Version + "\n")
	log.Default("")
}
