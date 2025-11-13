package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
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

type NoteTemplate struct {
	ProjectName string
	BranchName  string
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
	branchNotesDirLocation := filepath.Join(notesDir, projectName, branchDirName)
	branchNotesFileLocation := filepath.Join(branchNotesDirLocation, "notes.md")

	fmt.Println("Your current branch is: ", currentBranch)

	if _, err := os.Stat(branchNotesFileLocation); errors.Is(err, os.ErrNotExist) {
		fmt.Println("Note file not found, I will try to create it in", branchNotesDirLocation)

		err = os.MkdirAll(branchNotesDirLocation, 0755)

		if err != nil {
			return err
		}

		noteTemplateLocation := filepath.Join(notesDir, projectName, "NOTE_TEMPLATE.md")

		noteTemplate, err := os.ReadFile(noteTemplateLocation)
		var templateString string

		if os.IsNotExist(err) {
			fmt.Println("No NOTE_TEMPLATE.md file found, creating notes file from default template")
			templateString = "# {{.ProjectName}} - {{.BranchName}}"
		} else if err != nil {
			return err
		} else {
			fmt.Println("Creating note file from template found in branch dir")
			templateString = string(noteTemplate[:])
		}

		templateArgs := NoteTemplate{projectName, currentBranch}

		tmpl, err := template.New("note").Parse(templateString)

		if err != nil {
			return err
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, templateArgs)

		if err != nil {
			return err
		}

		err = os.WriteFile(branchNotesFileLocation, buf.Bytes(), 0664)

		fmt.Println("Note file created in", branchNotesDirLocation)

		if err != nil {
			return err
		}

		return nil
	}

	fmt.Println("You already have a note for this branch in", branchNotesDirLocation)

	return nil
}
