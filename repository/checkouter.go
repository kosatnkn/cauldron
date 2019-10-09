package repository

import (
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// Checkout checks out the given tag or the latest tag when a tag is not given.
func Checkout(r *git.Repository, tag string) error {

	if tag == "" {
		return checkoutLatestTag(r)
	}

	return checkoutTag(tag)
}

// checkoutLatestTag checks out the latest tag.
func checkoutLatestTag(r *git.Repository) error {

	commit, tag, err := getLatestTag(r)
	if err != nil {
		return err
	}

	fmt.Printf("Tag not provided, checking out latest tag (%s)", tag)

	return checkoutCommit(r, commit)
}

// checkoutTag checks out the given tag.
func checkoutTag(r *git.Repository, tag string) error {

}

// checkoutCommit checks out the given commit.
func checkoutCommit(r *git.Repository, c *object.Commit) error {

	// get worktree
	w, err := r.Worktree()
	if err != nil {
		return err
	}

	// create checkout options
	opts := &git.CheckoutOptions{
		Hash: c.Hash,
	}

	// checkout
	return w.Checkout(opts)
}

// getLatestTag returns the commit and the tag name of the latest tag.
func getLatestTag(repository *git.Repository) (*object.Commit, string, error) {

	tagRefs, err := repository.Tags()
	if err != nil {
		return nil, "", err
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
		return nil, "", err
	}

	return latestTagCommit, latestTagName, nil
}

// getTag returns the commit of the given tag.
func getTag(repository *git.Repository, tag string) (*object.Commit, error) {

	tagRefs, err := repository.Tags()
	if err != nil {
		return nil, err
	}

	var tagCommit *object.Commit

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

		if tagCommit == nil {
			tagCommit = commit
		}

		if commit.Committer.When.After(tagCommit.Committer.When) {
			tagCommit = commit
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return tagCommit, nil
}
