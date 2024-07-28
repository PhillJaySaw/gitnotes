package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func getCurrentBranchName() (string, error) {
	head, err := os.ReadFile(".git/HEAD")

	if err != nil {
		return "", errors.New("no git repo found or head missing")
	}

	headString := string(head)

	if strings.Contains(headString, "refs/heads") == false {
		return "", errors.New("on detached HEAD")
	}

	s := strings.Split(string(head), "refs/heads/")
	branchName := s[len(s)-1]

	return strings.TrimSpace(branchName), nil
}

func createNotesDirForCurrentBranch() error {
	currentBranch, err := getCurrentBranchName()

	if err != nil {
		return err
	}

	fmt.Println("Your current branch is: ", currentBranch)
	fmt.Println("I will try to create a dir for it")

	err = os.Mkdir(currentBranch, 0755)

	if err != nil {
		return err
	}

	return nil
}
