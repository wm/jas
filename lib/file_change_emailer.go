package jas

import (
	"log"

	"github.com/google/go-github/github"
)

// FileChangeEmailer is a handler that captures the files changed and if it
// finds a match emails the relevant people.
type FileChangeEmailer struct {
	opt FileChangeEmailerOptions
}

// Options is a struct for specifying configuration options for the
// FileChangeEmailer jas.Handler
type FileChangeEmailerOptions struct {
	// Emails is the list of email addresses to mail when a change is
	// detected
	Emails *[]string

	// Files is the list of files to monitor for change
	Files *[]string
}

// NewFileChangeEmailer returns a new FileChangeEmailer instance
func NewFileChangeEmailer(opt FileChangeEmailerOptions) *FileChangeEmailer {
	return &FileChangeEmailer{
		opt: opt,
	}
}

func (fce *FileChangeEmailer) HandlePayload(event []string, payload *github.WebHookPayload) {
	var files []string
	for _, commit := range payload.Commits {
		files = append(files, commit.Modified...)
		files = append(files, commit.Added...)
		files = append(files, commit.Removed...)
	}
	fce.emailIfWatchedFiles(&files)
}

func (fce *FileChangeEmailer) emailIfWatchedFiles(files *[]string) {
	// To be implemented - duh!
	log.Println("Files changed:", files)
	log.Println("Files to watch:", fce.opt.Files)
	log.Println("Peeps to email:", fce.opt.Emails)
}
