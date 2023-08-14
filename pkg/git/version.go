package git

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLatestReleaseTag(repositoryURL string) (string, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repositoryURL)
	response, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch release information: %s", response.Status)
	}

	var releaseInfo struct {
		TagName string `json:"tag_name"`
	}
	err = json.NewDecoder(response.Body).Decode(&releaseInfo)
	if err != nil {
		return "", err
	}

	return releaseInfo.TagName, nil
}
