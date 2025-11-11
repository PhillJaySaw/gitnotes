package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func getProjectName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	projectPath := strings.TrimSpace(string(output))
	result := filepath.Base((projectPath))

	return result, nil
}

func createNotesDirForCurrentBranch(notesDir string) error {
	currentBranch, err := getCurrentBranchName()

	if err != nil {
		return err
	}

	projectName, err := getProjectName()

	if err != nil {
		return err
	}

	branchDirName := strings.ReplaceAll(currentBranch, "/", "_")
	dirLocation := filepath.Join(notesDir, projectName, branchDirName)

	fmt.Println("Your current branch is: ", currentBranch)
	fmt.Println("I will try to create a directory for it at", dirLocation)

	err = os.MkdirAll(dirLocation, 0755)

	if err != nil {
		return err
	}

	return nil
}
