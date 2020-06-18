package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/ktrysmt/go-bitbucket"
)

// Hermes lists out the possible causes of anomalies.
type Hermes interface {
	CodeChanges(time.Time, IsAnomaly) ([]CommitInfo, error)
	SystemChanges(time.Time, IsAnomaly) ([]Activity, error)
}

// hermes is a concrete implementation of Hermes
type hermes struct {
	config Config
}

func (svc hermes) CodeChanges(date time.Time, isAnomaly IsAnomaly) ([]CommitInfo, error) {

	var commits []CommitInfo
	c := bitbucket.NewBasicAuth(svc.config.WorkSpace, svc.config.AppPassword)

	repos, err := svc.GetRepos(isAnomaly)
	if err != nil {
		return commits, err
	}

	for i := 0; i < len(repos); i++ {
		opt := &bitbucket.CommitsOptions{
			Owner:    svc.config.Owner,
			RepoSlug: repos[i],
		}

		res, err := c.Repositories.Commits.GetCommits(opt)
		if err != nil {
			return commits, nil
		}

		repoCommits, err := GetCommitsForRepo(res, date, repos[i])
		if err != nil {
			return commits, err
		}

		commits = append(commits, repoCommits...)

	}

	return commits, nil
}

func (svc hermes) SystemChanges(date time.Time, isAnomaly IsAnomaly) ([]Activity, error) {
	var (
		nResp activityResponse
		ret   []Activity
	)
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

	var client http.Client

	res, err := client.Do(req)

	if err != nil {
		return nResp.Results, nil
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&nResp)

	for _, act := range nResp.Results {
		if isAnomaly.Dau {
			if screenActivity(act, svc.config.DAUSVC) {
				ret = append(ret, act)
				continue
			}
		}
		if isAnomaly.Impressions {
			if screenActivity(act, svc.config.ImpressionsSVC) {
				ret = append(ret, act)
				continue
			}
		}
		if isAnomaly.Responses {
			if screenActivity(act, svc.config.RequestsSVC) {
				ret = append(ret, act)
				continue
			}
		}
		if isAnomaly.Requests {
			if screenActivity(act, svc.config.ResponsesSVC) {
				ret = append(ret, act)
				continue
			}
		}
	}

	return ret, err
}

func newHermesService(config Config) Hermes {
	svc := &hermes{
		config: config,
	}
	return svc
}

// GetCommitsForRepo gets all commits in a repo given the repo slug
func GetCommitsForRepo(res interface{}, date time.Time, slug string) ([]CommitInfo, error) {
	var commits []CommitInfo

	if res, ok := res.(map[string]interface{}); ok {
		if allCommits, ok := res["values"].([]interface{}); ok {

			for i := range allCommits {
				var commit CommitInfo

				if authorInfo, ok := allCommits[i].(map[string]interface{}); ok {
					if commitAuthor, ok := authorInfo["author"].(map[string]interface{}); ok {
						commit.Author, _ = commitAuthor["raw"].(string)
					}
				}
				if commitMessage, ok := allCommits[i].(map[string]interface{}); ok {
					if temp, ok := commitMessage["message"].(string); ok {
						commit.Message = strings.Split(temp, "\n")[0]
					}
				}
				if commitDate, ok := allCommits[i].(map[string]interface{}); ok {
					if temp, ok := commitDate["date"].(string); ok {
						commit.Date, _ = time.Parse(time.RFC3339, temp)
					}
				}

				if commit.Date.After(date.AddDate(0, 0, -2)) && commit.Date.Before(date.AddDate(0, 0, 1)) {
					commit.RepoSlug = slug
					commits = append(commits, commit)
				}
			}
			return commits, nil
		}
	}

	return commits, errors.New("Error in unmarshalling bitbucket response")
}

func (svc hermes) GetRepos(isAnomaly IsAnomaly) ([]string, error) {
	var repos []string
	if isAnomaly.Dau {
		repos = append(repos, svc.config.DAURepos...)
	}
	if isAnomaly.Impressions {
		repos = append(repos, svc.config.ImpressionsRepos...)
	}
	if isAnomaly.Requests {
		repos = append(repos, svc.config.RequestsRepos...)
	}
	if isAnomaly.Responses {
		repos = append(repos, svc.config.ResponsesRepos...)
	}

	return repos, nil
}

func screenActivity(activity Activity, whiteListed []string) bool {
	for _, a := range whiteListed {
		if activity.Callee == a {
			return true
		}
	}
	return false
}
