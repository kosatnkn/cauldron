package cli

import (
	"fmt"

	"github.com/kosatnkn/cauldron/config"
	"github.com/kosatnkn/cauldron/log"
)

// ShowConfig prints config values to the cli.
func ShowConfig(cfg *config.Config) {

	log.Info("Project Configurations")

	tag := cfg.Base.Version
	if tag == "" {
		if cfg.Base.MaxVersion == "" {
			tag = "latest"
		} else {
			tag = cfg.Base.MaxVersion
		}
	}

	m := fmt.Sprintf(` Project Name      : %s
 Project Namespace : %s
 Base Repository   : %s
 Base Repo Version : %s`, cfg.Project.Name, cfg.Project.Namespace, cfg.Base.Repo, tag)

	log.Default(m)
}

// ShowComplete shows the completed message.
func ShowComplete(cfg *config.Config) {

	m := fmt.Sprintf("Project '%s' has been created successfully\n", cfg.Project.Name)
	log.Info(m)
}
