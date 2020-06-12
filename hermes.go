package main

import (
	"time"

	"github.com/ktrysmt/go-bitbucket"
)

// Hermes lists out the possible causes of anomalies.
type Hermes interface {
	CodeChanges(string) ([]commitInfo, error)
	// SystemChanges(string) string
}

// hermes is a concrete implementation of Hermes
type hermes struct {
	config Config
}

func (svc hermes) CodeChanges(date string) ([]commitInfo, error) {

	var commits []commitInfo
	c := bitbucket.NewBasicAuth(svc.config.WorkSpace, svc.config.AppPassword)

	opt := &bitbucket.CommitsOptions{
		Owner:    svc.config.Owner,
		RepoSlug: svc.config.RepoSlug,
	}

	res, err := c.Repositories.Commits.GetCommits(opt) //c.GetCommts(opt)
	if err != nil {
		panic(err)
	}
	allCommits := res.(map[string]interface{})["values"].([]interface{})

	for i := range commits {
		var temp commitInfo
		temp.Author = allCommits[i].(map[string]interface{})["author"].(map[string]interface{})["raw"].(string)
		temp.Message = allCommits[i].(map[string]interface{})["message"].(string)
		temp.Date = allCommits[i].(map[string]interface{})["message"].(time.Time)
	}

	return commits, nil
}

// func (hermes) SystemChanges(date string) string {
// 	return date
// }

func newHermes(config Config) Hermes {
	svc := &hermes{
		config: config,
	}
	return svc
}
