package esi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sethgrid/pester"
)

func fetchURL(ctx context.Context, client *pester.Client, url string, r interface{}) (http.Header, error) {
	// log.Printf("Fetching %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "go-evepraisal")
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case 200:
	case 404:
		return nil, nil
	default:
		return nil, fmt.Errorf("Error talking to esi: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(r)
	defer resp.Body.Close()
	return resp.Header, nil
}
