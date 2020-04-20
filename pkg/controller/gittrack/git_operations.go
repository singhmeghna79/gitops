package gittrack

import (
	"bufio"
	"os"
	"os/exec"
	"path"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing"

	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"

	git "gopkg.in/src-d/go-git.v4"
)

func (g *GitOperations) clone(clonePath string) error {
	gitCloneOptions := g.getCloneOptions()
	directory := path.Join("/tmp", clonePath)
	repository, err := git.PlainClone(directory, false, gitCloneOptions)
	if err != nil {
		return err
	}
	g.Repository = repository
	return nil
}

func (g *GitOperations) getCloneOptions() *git.CloneOptions {
	return &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: g.Username,
			Password: g.Password,
		},
		URL:      g.RepositoryName,
		Progress: os.Stdout,
	}
}

func (g *GitOperations) checkoutBranch() error {
	w, err := g.Repository.Worktree()
	if err != nil {
		return err
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(g.Branch),
	})

	if err != nil {
		return err
	}
	return nil
}

// func (g *GitOperations) getChangedFilePaths(subPath string, sha1, sha2 [20]byte) ([]string, error) {
// 	commit1, err := g.Repository.CommitObject(plumbing.Hash(sha1))
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	commit2, err := g.Repository.CommitObject(plumbing.Hash(sha2))
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	patch, err := commit1.Patch(commit2)
// 	if err != nil {
// 		return []string{}, err
// 	}
// 	fileStats := patch.Stats()
// 	var changedFileList []string
// 	for _, v := range fileStats {
// 		changedFileList = append(changedFileList, v.Name)
// 	}
// 	return changedFileList, nil
// }

func (g *GitOperations) getChangedFilePaths(repoPath, sha1, sha2 string) ([]string, error) {
	directory := path.Join("/tmp", repoPath)
	dir, _ := os.Getwd()
	os.Chdir(directory)
	out, err := exec.Command("git", "diff", "--name-only", sha1, sha2).Output()
	os.Chdir(dir)
	if err != nil {
		return []string{}, err
	}
	var fileList []string
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		fileList = append(fileList, scanner.Text())
	}
	return fileList, nil
}

func getDiff(sha1, sha2 string) {
	plumbing.Patch(sha1)
}
