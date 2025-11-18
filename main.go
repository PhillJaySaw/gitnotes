package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initalizing GitNotes")
}

func main() {

	// TODO handle `gitnotes` - create note for branch

	// TODO handle `gitnotes help` - show commands

	// TODO handle `gitnotes list` - list of all notes for current project

	config, err := initConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = createNotesDirForCurrentBranch(config.NotesDir)

	if err != nil {
		fmt.Println(err)
	}
}
