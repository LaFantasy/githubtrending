package githubtrending

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	clientTimeOut   = 10 * time.Second
	apiEndpoint     = `https://github-trending-api.now.sh`
	devApiEndpoint  = apiEndpoint + "/developers"
	repoApiEndpoint = apiEndpoint + "/repositories"
)

var (
	client = http.Client{
		Timeout: clientTimeOut,
	}
)

// GetDevelopers retrieves a list of GitHub trending developers from the API.
func GetDevelopers() (devs []*Developer, err error) {
	var resp *http.Response
	if resp, err = client.Get(devApiEndpoint); err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&devs)
	if err == nil && len(devs) == 0 {
		err = fmt.Errorf("got no trending developers")
	}
	return
}

// GetRepositories retrieves a list of GitHub trending repositories from the API.
func GetRepositories() (repos []*Repository, err error) {
	var resp *http.Response
	if resp, err = client.Get(repoApiEndpoint); err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err == nil && len(repos) == 0 {
		err = fmt.Errorf("got no trending repositories")
	}
	return
}
