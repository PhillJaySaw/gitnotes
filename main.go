package main

import (
	"fmt"
)

func main() {

	config, err := initConfig()

	if err != nil {
		fmt.Println(err)
	}

	err = createNotesDirForCurrentBranch(config.NotesDir)	

	if err != nil {
		fmt.Println(err)
	}
}
