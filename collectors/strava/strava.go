package strava

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/edznux/lyfe/storage"
)

type Config struct {
	ClientID     string
	ClientSecret string
}

type Collector struct {
	Config Config
	store  *storage.Store
}

func NewCollector(store *storage.Store) *Collector {
	return &Collector{
		store: store,
	}
}

func RequestActivities(authToken string) ([]Activity, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+authToken)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	activities := []Activity{}
	err = json.NewDecoder(resp.Body).Decode(&activities)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

// FIXME, use the userID instead of my hardcoded username
func (c *Collector) Collect(ctx context.Context, userID int64) error {
	authToken, err := c.store.Querier.GetAuthTokenForUsername(ctx, "edznux")
	if err != nil {
		return err
	}
	activities, err := RequestActivities(authToken.AccessToken)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", activities)
	return nil
}

func (c *Collector) Init() error {
	return nil
}
