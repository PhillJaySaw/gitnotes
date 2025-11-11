package main

import (
	"fmt"
)

func init() {
	fmt.Println("Initalizing GitNotes")
}
	
func main() {

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
