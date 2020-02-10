package githubtrending

import (
	"testing"
)

func TestGetDevelopers(t *testing.T) {
	gotDevs, err := GetDevelopers()
	if err != nil {
		t.Errorf("GetDevelopers() got error = %v", err)
	} else if len(gotDevs) <= 0 {
		t.Errorf("GetDevelopers() got no results")
	} else {
		for idx, dev := range gotDevs {
			t.Logf("#%d - %v", idx+1, dev)
		}
	}
}

func TestGetRepositories(t *testing.T) {
	gotRepos, err := GetRepositories()
	if err != nil {
		t.Errorf("GetRepositories() got error = %v", err)
	} else if len(gotRepos) <= 0 {
		t.Errorf("GetRepositories() got no results")
	} else {
		for idx, repo := range gotRepos {
			t.Logf("#%d - %v", idx+1, repo)
		}
	}
}
