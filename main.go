package main

import (
	"fmt"
)

func main() {

	err := createNotesDirForCurrentBranch()

	if err != nil {
		fmt.Println(err)
	}
}
