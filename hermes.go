package main

import (
	"strings"
	"time"

	"github.com/ktrysmt/go-bitbucket"
)

// Hermes lists out the possible causes of anomalies.
type Hermes interface {
	CodeChanges(time.Time) ([]CommitInfo, error)
	// SystemChanges(string) string
}

// hermes is a concrete implementation of Hermes
type hermes struct {
	config Config
}

func (svc hermes) CodeChanges(date time.Time) ([]CommitInfo, error) {

	var commits []CommitInfo
	c := bitbucket.NewBasicAuth(svc.config.WorkSpace, svc.config.AppPassword)

	opt := &bitbucket.CommitsOptions{
		Owner:    svc.config.Owner,
		RepoSlug: svc.config.RepoSlug,
	}

	res, err := c.Repositories.Commits.GetCommits(opt)
	if err != nil {
		return commits, nil
	}
	allCommits := res.(map[string]interface{})["values"].([]interface{})

	for i := range allCommits {
		var temp CommitInfo
		temp.Author = allCommits[i].(map[string]interface{})["author"].(map[string]interface{})["raw"].(string)
		temp.Message = strings.Split(allCommits[i].(map[string]interface{})["message"].(string), "\n")[0]
		temp.Date, _ = time.Parse(time.RFC3339, allCommits[i].(map[string]interface{})["date"].(string))

		if temp.Date.After(date.AddDate(0, 0, -2)) && temp.Date.Before(date.AddDate(0, 0, 1)) {
			commits = append(commits, temp)
		}
	}

	return commits, nil
}

// func (hermes) SystemChanges(date string) string {
// 	return date
// }

func newHermesService(config Config) Hermes {
	svc := &hermes{
		config: config,
	}
	return svc
}
