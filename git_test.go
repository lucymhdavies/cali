package cali

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContainerName(t *testing.T) {
	t.Log("Simple example")
	name, err := GitCheckoutConfig{
		Repo:    "repo",
		Branch:  "branch",
		RelPath: ".",
	}.GetContainerName()
	if assert.NoError(t, err) {
		assert.Equal(t, "data_repo_branch_89bb6900736548ebd6455d0ab07aa5fe", name)
	}

	t.Log("Simple example, with empty path")
	name, err = GitCheckoutConfig{
		Repo:    "repo",
		Branch:  "branch",
		RelPath: "",
	}.GetContainerName()
	if assert.NoError(t, err) {
		assert.Equal(t, "data_repo_branch_89bb6900736548ebd6455d0ab07aa5fe", name)
	}

	t.Log("Simple example, with unspecified path")
	name, err = GitCheckoutConfig{
		Repo:   "repo",
		Branch: "branch",
	}.GetContainerName()
	if assert.NoError(t, err) {
		assert.Equal(t, "data_repo_branch_89bb6900736548ebd6455d0ab07aa5fe", name)
	}

	t.Log("Branches and paths with slashes")
	name, err = GitCheckoutConfig{
		Repo:    "repo",
		Branch:  "branch/with-slash",
		RelPath: "path/containing/slashes",
	}.GetContainerName()
	if assert.NoError(t, err) {
		assert.Equal(t, "data_repo_path-containing-slashes_branch-with-slash_89bb6900736548ebd6455d0ab07aa5fe", name)
	}
}

func TestRepoNameFromURL(t *testing.T) {
	t.Log("Example git url")
	name, err := repoNameFromURL("git@github.com:skybet/cali.git")
	if assert.NoError(t, err) {
		assert.Equal(t, "github-com-skybet-cali", name)
	}

	t.Log("Example https url")
	name, err = repoNameFromURL("https://github.com/skybet/cali.git")
	if assert.NoError(t, err) {
		assert.Equal(t, "github-com-skybet-cali", name)
	}

	t.Log("Example git url with ssh protocol")
	name, err = repoNameFromURL("ssh://git@github.com/skybet/cali.git")
	if assert.NoError(t, err) {
		assert.Equal(t, "github-com-skybet-cali", name)
	}
}
