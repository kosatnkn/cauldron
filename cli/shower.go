package cli

import (
	"fmt"

	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/log"
)

// ShowConfig prints config values to the cli.
func ShowConfig(cfg *config.Config) {

	log.Info("Configurations")

	m := fmt.Sprintf(`
	Name     : %s
	Namespace: %s
	Tag      : %s
	`, cfg.Name, cfg.Namespace, cfg.Tag)

	log.Default(m)
}
