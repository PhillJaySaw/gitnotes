package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	NotesDir string `toml:"notes_dir"`
}

const CONFIG_LOCATION = ".config/gitnote"
const CONFIG_FILE_NAME = "gitnotes.toml"

func initConfig() (*Config, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	fullPathToConfigDir := filepath.Join(homedir, CONFIG_LOCATION)
	fullPathToConfigFile := filepath.Join(homedir, CONFIG_LOCATION, CONFIG_FILE_NAME)

	_, err = os.Stat(fullPathToConfigFile)

	if os.IsNotExist(err) {
		fmt.Println("No config file found, creating one in " + fullPathToConfigDir)
		fmt.Println()

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

		fmt.Println("Config created successfully")
		fmt.Println()
	}

	var config Config

	_, err = toml.DecodeFile(fullPathToConfigFile, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
