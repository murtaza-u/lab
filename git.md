# Learning Git

## Git config

* `git config --global user.name "Murtaza Udaipurwala"`
* `git config --global user.email "murtaza.u6e24@gmail.com"`
* `git config --list`

## Basic workflow

* `git init`
* `git status`
* `git add [file|.|-A]`
* `git diff`
* `git diff hash1 hash2`
* `git commit [-m]`
* `git log`
* `git remote -v`
* `git remote set-url origin url`
* `git push [-f] [branch:-origin]`
* `git pull [branch:-origin]`
* `git branch -l`
* `git branch --show-current`
* `git branch name`
* `git checkout [-b] [branch|hash]`
* `git switch branch`
* `git clone url`
* `git branch [-d|-D] branch`
* `git push origin --delete branch`

## Fixing mistakes and Undoing bad commit

* `git checkout file`
* `git commit --amend -m "implement subtract function"`
* `git commit --amend`
* `git log --stat`
* `git cherry-pick latest-hash`
* Reset
    1. soft - staged (`git reset --soft [hash]`)
    2. mix(default) - unstaged (`git reset [hash]`)
    3. hard - changes gone (`git reset --hard [hash]`)
* `git reset` (unstage all staged files)
* `git clean -df` (removes any untracked files/directory)
* `git reflog`
* `git restore file`
* `git restore --staged file`

## Stash

* `git stash save 'message'`
* `git stash list`
* `git stash apply stash@{0}`
* `git stash pop`
* `git stash drop stash@{0}`
* `git stash clear`

## Diff Tool

* `git difftool`
* `.gitconfig`

```toml
[difftool]
    prompt = false
[diff]
    tool = nvimdiff
```

## add subcommand

* `git add -A` (default behaviour)
* `git add -A` != `git add .` (at least not always. Hint: a subdirectory)
* `git add --no-all` (doesn't stage deleted files)
* `git add -u` (doesn't stage new files)

## Rebase

* `git rebase -i HEAD~3`

# GitHub CLI

* `gh auth login`
* `gh config list`
* `gh repo create`
* `gh repo delete`
* `gh repo clone`
* `gh repo fork`
* `gh gist create`
* `gh gist delete <id>`
* `gh gist edit <id>`
* `gh gist list [--public|--private]`
* `gh issue list`
* `gh issue create`
* `gh issue status`
