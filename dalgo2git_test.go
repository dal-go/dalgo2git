package dalgo2git

import "testing"
import git "github.com/go-git/go-git/v5"

func TestGit(t *testing.T) {
	// Open the repo (assumes current dir is a git repo)
	r, err := git.PlainOpen(".")
	if err != nil {
		panic(err)
	}

	// Get the working directory
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}
	if w == nil {
		t.Fatal("worktree is nil")
	}
}
