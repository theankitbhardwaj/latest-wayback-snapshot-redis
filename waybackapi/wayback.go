package waybackapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Closest struct {
	Status    string `json:"status"`
	Available bool   `json:"available"`
	Url       string `json:"url"`
	Timestamp string `json:"timestamp"`
}

type ArchivedSnapshots struct {
	Closest Closest `json:"closest"`
}
type Response struct {
	Url               string            `json:"url"`
	ArchivedSnapshots ArchivedSnapshots `json:"archived_snapshots"`
}

func GetSnapshotUrl(url string, timestamp string) string {
	apiUrl := "https://archive.org/wayback/available?url=%v&timestamp=%v"

	reqUrl := fmt.Sprintf(apiUrl, url, timestamp)

	req, _ := http.NewRequest("GET", reqUrl, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	response := &Response{}

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return ""
	}

	return response.ArchivedSnapshots.Closest.Url
}
