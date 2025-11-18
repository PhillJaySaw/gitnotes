# git notes

## user flow

1. create new branch in git project
2. git notes automatically creates a notes file based on config after executing `gitnotes` command

## config

The config will be created when gitnotes is execute for the first time, if it does not exist.
The location of the config is `~/.config/gitnotes`. For now the only thing that this file configures is the location of the notes dir used by `gitnotes` (default is set to `~/gitnotes/notes`).    

## notes dir structure

gitnotes/
    notes/
        project-1/
            branch-name.md
            branch2-name.md
            branch3-name.md
        project-2/
            branch-name.md
            branch2-name.md
        project-3/


## features

### implemented 

- create notes file for a specific branch in notes folder
- add custom NOTE_TEMPLATE.md file which can be use when creating notes for a project

### no implemented 

- set status in note file to auto archive (ex. write DONE at the top of the note)
- runs in the background, no need to run command in git repo, note will be created automagically when branch is created
- add more command for listing all notes, listing notes in project
- different notes for different branch types (ex. `feature/branch`, `bugfix/branch`)
- update note with info/metadata about the branch, like commits, dates etc. 
- improved logging with `github.com/charmbracelet/log`

## POC

- [x] go project
- [x] run go program in repo
- [x] read basic setup from config file in `./.config/git-notes.config`
- [ ] detect branch changes 
