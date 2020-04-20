package gittrack

import git "gopkg.in/src-d/go-git.v4"

// GitOperations will be responsible to clone the git repository
type GitOperations struct {
	RepositoryName string
	Repository     *git.Repository
	Branch         string
	SubPath        string
	Username       string
	Password       string
	Type           string
}
