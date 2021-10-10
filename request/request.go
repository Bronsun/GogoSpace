package request

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Bronsun/GogoSpace/config"
	"github.com/Bronsun/GogoSpace/utils"
)

var (
	ErrFailedRequest = errors.New("Failed to create request")
)

type URLResponse struct {
	URL string `json:"url"`
}

const (
	baseURL = "https://api.nasa.gov/planetary/apod"
)

// CreateRequest creates custom request to extranl api
func CreateRequest(ctx context.Context, apiKey string, date time.Time) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}
	q := req.URL.Query()
	q.Add("api_key", apiKey)
	q.Add("date", date.Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()
	return req, nil
}

// GetImagesFromRequest main logic for getting all urls from extranal api
func GetImagesFromRequest(start, end time.Time) (urls []string, err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Get dates from which we will be taking photos
	dates, err := utils.GetDatesFromQuery(start, end)
	if err != nil {
		return nil, err
	}
	// Custom http.Client with 5sec cooldown - good bot
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	// Limiting active goroutines - could be better :(
	concurrentGoroutines := make(chan struct{}, config.GetConcurrentRequests())

	var wg sync.WaitGroup

	for _, date := range dates {
		concurrentGoroutines <- struct{}{}
		wg.Add(1)
		dateTime, _ := time.Parse("2006-01-02", date)
		go func(date time.Time) {
			defer wg.Done()

			// Make a custom request
			req, err := CreateRequest(ctx, config.GetApiKey(), dateTime)
			if err != nil {
				cancel()
				return
			}

			// Send a custom request to extranl api with limiter
			resp, err := client.Do(req)
			<-concurrentGoroutines
			if err != nil {
				cancel()
				return
			}
			if resp.StatusCode != http.StatusOK {
				cancel()
				return
			}
			defer resp.Body.Close()
			decoder := json.NewDecoder(resp.Body)
			var r URLResponse
			err = decoder.Decode(&r)
			if err != nil {

				cancel()
				return
			}
			// List of our urls
			urls = append(urls, r.URL)

		}(dateTime)
	}
	wg.Wait()
	return urls, nil
}
