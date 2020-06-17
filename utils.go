package main

import (
	"time"
)

// Config contains the env config parameters to run the service
type Config struct {
	UserID      string   `required:"true"`
	AuthToken   string   `required:"true"`
	Endpoint    string   `required:"true"`
	WorkSpace   string   `required:"true"`
	AppPassword string   `required:"true"`
	Owner       string   `required:"true"`
	DAU         []string `required:"true"`
	Impressions []string `required:"true"`
	Requests    []string `required:"true"`
	Responses   []string `required:"true"`
}

// App contains all fields of app
type App struct {
	Date        time.Time `json:"date"`
	ID          string    `json:"app"`
	Dau         int       `json:"dau"`
	Requests    int       `json:"requests"`
	Responses   int       `json:"responses"`
	Impressions int       `json:"impressions"`
}

type appNumbers struct {
	app             App
	meanDau         float64
	stdDau          float64
	meanRequests    float64
	stdRequests     float64
	meanResponses   float64
	stdResponses    float64
	meanImpressions float64
	stdImpressions  float64
}

// CommitInfo contains attributes of commit required
type CommitInfo struct {
	Date     time.Time `json:"date"`
	Message  string    `json:"message"`
	Author   string    `json:"author"`
	RepoSlug string    `json:"repo_slug"`
}

type activityResponse struct {
	Results []Activity `json:"results"`
}

// Activity contains the activity log for a date
type Activity struct {
	Callee    string `json:"callee"`
	Text      string `json:"text"`
	Level     string `json:"level"`
	Tag       string `json:"tag"`
	CreatedAt string `json:"created_at"`
}

// IsAnomaly tells whether an anomaly exists for a field
type IsAnomaly struct {
	Dau         bool
	Impressions bool
	Requests    bool
	Responses   bool
}

func signum(x float64) float64 {
	if x >= 0 {
		return 1.0
	}
	return -1.0
}
