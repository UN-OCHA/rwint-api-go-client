package rwapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Base URL for the ReliefWeb API.
const RWAPIURL = "https://api.reliefweb.int/v1/"

// A Client is used to query the ReliefWeb API.
type Client struct {
	AppName string
	Client  *http.Client
}

// NewClient creates a new ReliefWeb API client.
// Appname identifies the application using querying the API.
// Timeout sepcifies the time limit for the request.
// That includes the connection, redirects and reading the response body.
func NewClient(appname string, timeout time.Duration) *Client {
	client := &http.Client{Timeout: timeout}
	return &Client{AppName: appname, Client: client}
}

// Query queries the ReliefWeb API. The resource can be a full resource
// or resource + item ID in the form "resource/ID".
// It returns the raw JSON response payload.
func (c *Client) QueryRaw(resource string, query *Query) ([]byte, error) {
	url := RWAPIURL + resource

	// Add the appname to know who is using the API.
	if c.AppName != "" {
		url += "?appname=" + c.AppName
	}

	data, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("Unable to serialize ReliefWeb API query payload: %s", err)
	}

	response, err := c.Client.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Unable to query ReliefWeb API with request %s", url)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read ReliefWeb API response body for request %s with payload %s: %s", url, string(data), err)
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected ReliefWeb API response for request %s with payload %s: %s", url, string(data), string(body))
	}
	return body, nil
}

// Query queries the ReliefWeb API. The resource can be a full resource
// or resource + item ID in the form "resource/ID".
// It returns an unserialized version of the response payload.
func (c *Client) Query(resource string, query *Query) (*Result, error) {
	body, err := c.QueryRaw(resource, query)
	if err != nil {
		return nil, err
	}

	var result *Result
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("Unable to unserialize ReliefWeb API response %s: %s", string(body), err)
	}
	return result, nil
}
