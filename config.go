package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/philljaysaw/gitnotes/logger"
)

type Config struct {
	NotesDir string `toml:"notes_dir"`
}

const CONFIG_LOCATION = ".config/gitnote"
const CONFIG_FILE_NAME = "gitnotes.toml"

func initConfig() (*Config, error) {
	log := logger.New()
	homedir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	fullPathToConfigDir := filepath.Join(homedir, CONFIG_LOCATION)
	fullPathToConfigFile := filepath.Join(homedir, CONFIG_LOCATION, CONFIG_FILE_NAME)

	_, err = os.Stat(fullPathToConfigFile)

	if os.IsNotExist(err) {
		log.Info("No config file found, creating one in " + fullPathToConfigDir)

		err = os.MkdirAll(fullPathToConfigDir, 0755)

		if err != nil {
			return nil, err
		}

		configFile, err := os.Create(fullPathToConfigFile)
		defer configFile.Close()
		if err != nil {
			return nil, err
		}

		s := fmt.Sprintf(`notes_dir = "%s/gitnotes/notes"`, homedir)
		_, err = fmt.Fprintln(configFile, s)
		if err != nil {
			return nil, err
		}

		err = configFile.Sync()
		if err != nil {
			return nil, err
		}

		log.Info("Config created successfully")
	}

	var config Config

	_, err = toml.DecodeFile(fullPathToConfigFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
