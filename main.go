package main

import (
	"fmt"
)

func main() {

	config, err := initConfig()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
	fmt.Println(config.NotesDir)
}
