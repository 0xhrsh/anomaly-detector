package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Nostalgia contains tools required to get data from Service Nostalgia
type Nostalgia interface {
	FetchAppDataForRange(appID string, date time.Time, window int) (*NostalgiaResponse, error)
}

func newNostalgiaService(config Config) Nostalgia {
	return &nostalgiaImpl{
		config: config,
		client: &http.Client{
			Timeout: time.Minute * 5,
		},
	}
}

type nostalgiaImpl struct {
	config Config
	client *http.Client
}

func (n *nostalgiaImpl) FetchAppDataForRange(appID string, date time.Time, window int) (*NostalgiaResponse, error) {

	requestURL, err := url.Parse(n.config.Endpoint)
	if err != nil {
		return nil, err
	}
	requestURL.Path = path.Join(requestURL.Path, "v3", "nostalgia", "report")

	query := requestURL.Query()
	query.Add("app_id", appID)
	query.Add("from", date.AddDate(0, 0, -1*window).Format("2006-01-02"))
	query.Add("to", date.AddDate(0, 0, -1).Format("2006-01-02"))
	query.Add("dim", "date,app")
	query.Add("metrics", "ad_responses,ad_requests,impressions,dau")

	requestURL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Id", n.config.UserID)
	req.Header.Set("Auth-Token", n.config.AuthToken)
	res, err := n.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var nResp NostalgiaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&nResp)

	if err != nil {
		return nil, err
	}

	return &nResp, nil
}

// NostalgiaResponse contains struct to unmarshal Nostalgia response
type NostalgiaResponse struct {
	Result []App `json:"result"`
}

func (nResp NostalgiaResponse) Len() int {
	return len(nResp.Result)
}

func (nResp NostalgiaResponse) Swap(i, j int) {
	nResp.Result[i], nResp.Result[j] = nResp.Result[j], nResp.Result[i]
}

func (nResp NostalgiaResponse) Less(i, j int) bool {
	return nResp.Result[i].Date.After(nResp.Result[j].Date)
}
