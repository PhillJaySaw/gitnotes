package main

import "github.com/philljaysaw/gitnotes/logger"

func init() {
	log := logger.New()
	log.Info("Initalizing GitNotes")
}

func main() {
	log := logger.New()

	// TODO handle `gitnotes` - create note for branch

	// TODO handle `gitnotes help` - show commands

	// TODO handle `gitnotes list` - list of all notes for current project

	config, err := initConfig()

	if err != nil {
		log.Error(err)
		return
	}

	err = createNotesDirForCurrentBranch(config.NotesDir)

	if err != nil {
		log.Error(err)
		return
	}
}
