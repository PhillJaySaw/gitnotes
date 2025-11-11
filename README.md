# git notes

## user flow

1. create new branch in git project
2. git notes automatically creates a notes file based on config

## config

- notes template
- location of notes dir
- locations of projects to listen to


## notes dir structure

| notes dir
|\ wip (work in progress, working branches)
| | project name
| |\ branch_name_1.md
| | project name
| |\ branch_name_1.md
| |\ branch_name_2.md
| |\ branch_name_2.md
|
|\
| | project name
| |\ branch_name_1.md
| |\ branch_name_2.md
|
|\ archive (notes archive for completed branches)


## features

- set status in note file to auto archive (write DONE at the top for example)
- runs in the background
- execute command in project dir to automatically listen for branch changes and new notes

## milestones / brainstorm

- [ ] figure out a way to detect branch checkouts in repo
- [ ] figure out a way for git-notes to run in the background without having to launch it each time
- [ ] which git projects should be watche? config? maybe running command in project to start listening to it
- [ ] what if you checkout someone elses branch, should it be ignored? how to detect this?
- [ ] create notes files based on templates in config 
- [ ] search capability for notes based on content/branch name/commits
- [ ] maybe there could be a section in the notes which is being updated with changes made to the branch
- [ ] different note types based on branch template 'feature/...', 'test/...', etc.

## POC

- [ ] go project
- [ ] run go program in repo
- [ ] read basic setup from config file in `./.config/git-notes.config`
- [ ] detect branch changes 
