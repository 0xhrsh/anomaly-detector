package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/ktrysmt/go-bitbucket"
)

// Hermes lists out the possible causes of anomalies.
type Hermes interface {
	CodeChanges(time.Time) ([]CommitInfo, error)
	SystemChanges(time.Time) ([]Activity, error)
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

func (svc hermes) SystemChanges(date time.Time) ([]Activity, error) {
	var nResp activityResponse
	requestURL, err := url.Parse(svc.config.Endpoint)

	if err != nil {
		return nResp.Results, nil
	}

	requestURL.Path = path.Join(requestURL.Path, "v3", "logs", "search")

	query := requestURL.Query()
	query.Add("limit", "20")
	query.Add("offset", "0")
	query.Add("from", date.AddDate(0, 0, -2).Format("2006-01-02"))
	query.Add("to", date.Format("2006-01-02"))

	requestURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", requestURL.String(), nil)

	if err != nil {
		return nResp.Results, nil
	}

	req.Header.Set("User-Id", svc.config.UserID)
	req.Header.Set("Auth-Token", svc.config.AuthToken)
	req.Header.Set("DNT", "1")

	fmt.Println(req)
	var client http.Client

	res, err := client.Do(req)

	if err != nil {
		return nResp.Results, nil
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&nResp)

	return nResp.Results, err
}

func newHermesService(config Config) Hermes {
	svc := &hermes{
		config: config,
	}
	return svc
}
