package repository

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const catalystURL string = "https://github.com/kosatnkn/catalyst.git"

// GetBase creates the project base by cloning Catalyst
// and checking out either the given tag if one is provided
// or the latest tag if a tag is not provided.
// This will panic if the provided tag is not found.
func GetBase(name string, tag string) {

	// get clone dir
	dir, err := getCloningDir(name)
	if err != nil {
		panic(err)
	}

	// clone
	r, err := clone(dir)
	if err != nil {
		panic(err)
	}

	// get latest tag
	tag, err = getLatestTag(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(tag)
}

func getCloningDir(name string) (string, error) {

	// get current location
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// define directory to clone in to
	dir := path + string(os.PathSeparator) + name

	return dir, nil
}

// clone repository
func clone(dir string) (*git.Repository, error) {

	opts := &git.CloneOptions{
		URL:      catalystURL,
		Progress: os.Stdout,
	}

	r, err := git.PlainClone(dir, false, opts)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// get latest tag
func getLatestTag(repository *git.Repository) (string, error) {

	tagRefs, err := repository.Tags()
	if err != nil {
		return "", err
	}

	var latestTagCommit *object.Commit
	var latestTagName string

	err = tagRefs.ForEach(func(tagRef *plumbing.Reference) error {

		revision := plumbing.Revision(tagRef.Name().String())

		tagCommitHash, err := repository.ResolveRevision(revision)
		if err != nil {
			return err
		}

		commit, err := repository.CommitObject(*tagCommitHash)
		if err != nil {
			return err
		}

		if latestTagCommit == nil {
			latestTagCommit = commit
			latestTagName = tagRef.Name().String()
		}

		if commit.Committer.When.After(latestTagCommit.Committer.When) {
			latestTagCommit = commit
			latestTagName = tagRef.Name().String()
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return latestTagName, nil
}
