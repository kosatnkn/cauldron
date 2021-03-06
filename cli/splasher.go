package cli

import (
	"github.com/kosatnkn/cauldron/v2/config"
	"github.com/kosatnkn/cauldron/v2/log"
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

	usage := `
Usage:
  cauldron -n Sample -s github.com/username [-t v1.0.0]
  cauldron --name Sample --namespace github.com/username [--tag v1.0.0]

  This will create a new project with go.mod module
  github.com/username/sample`

	log.Default(splash)
	log.Default("  " + cfg.Cauldron.Version)
	log.Default(usage)
	log.Default("")
}
