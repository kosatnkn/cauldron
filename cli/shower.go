package cli

import (
	"fmt"

	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/log"
)

// ShowConfig prints config values to the cli.
func ShowConfig(cfg *config.Config) {

	log.Info("Project Configurations")

	tag := "latest"
	if cfg.Tag != "" {
		tag = cfg.Tag
	}

	m := fmt.Sprintf(`
 Name     : %s
 Namespace: %s
 Tag      : %s
`, cfg.Name, cfg.Namespace, tag)

	log.Default(m)
}

// ShowComplete shows the completed message.
func ShowComplete(cfg *config.Config) {

	m := fmt.Sprintf("Project `%s` has been created successfully\n", cfg.Name)
	log.Info(m)
}
