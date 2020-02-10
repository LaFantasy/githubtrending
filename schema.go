package githubtrending

import (
	"strings"

	"github.com/1set/gut/ystring"
)

// Repository consists of basic info of a GitHub trending repository and its developer.
type Repository struct {
	Author             string `json:"author"`
	Name               string `json:"name"`
	Avatar             string `json:"avatar"`
	URL                string `json:"url"`
	Description        string `json:"description"`
	Language           string `json:"language,omitempty"`
	LanguageColor      string `json:"languageColor,omitempty"`
	Stars              int    `json:"stars"`
	Forks              int    `json:"forks"`
	CurrentPeriodStars int    `json:"currentPeriodStars"`
	Owner              []struct {
		Username string `json:"username"`
		Href     string `json:"href"`
		Avatar   string `json:"avatar"`
	} `json:"builtBy"`
}

func (r Repository) String() string {
	sb := strings.Builder{}
	sb.Grow(len(r.URL)*2 + len(r.Description))

	if !ystring.IsBlank(r.Language) {
		sb.WriteString("[")
		sb.WriteString(r.Language)
		sb.WriteString("]")
	}

	sb.WriteString(r.Name)

	if !ystring.IsBlank(r.Description) {
		sb.WriteString(": ")
		sb.WriteString(strings.TrimSpace(r.Description))
	}

	sb.WriteString("|")
	sb.WriteString(r.URL)

	return sb.String()
}

// Developer consists of basic info of a GitHub trending developer and his/her trending repository.
type Developer struct {
	Username    string `json:"username"`
	DisplayName string `json:"name"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Avatar      string `json:"avatar"`
	Repo        struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"repo"`
}

func (d Developer) String() string {
	sb := strings.Builder{}
	sb.Grow(len(d.DisplayName) + len(d.Username) + len(d.URL) + 5)

	sb.WriteString(d.DisplayName)
	sb.WriteString(" (")
	sb.WriteString(d.Username)
	sb.WriteString(") ")
	sb.WriteString(d.URL)

	return sb.String()
}
