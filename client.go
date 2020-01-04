package goxi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// client struct defines the URL of the Nagios instance, the user's API key, and embeds a pointer to a HTTP client
type client struct {
	apiURL     string
	apiKey     string
	httpClient *http.Client
}

// Constants related to the URL segments for the API
const (
	apiV1Segment   string = "/api/v1"
	hostSegment    string = "/objects/hoststatus"
	serviceSegment string = "/objects/servicestatus"
	apiParam       string = "apikey"
)

// NewClient returns a client struct with an embedded HTTP Client to interact with Nagios API
// apiURL: the URL of the Nagios instance (e.g.: https://nagios.example/nagiosxi
// apiKey: api key of the user generated from the User Dashboard
//TODO: allow user to configure a custom client (i.e.: timeout, TLS verification)
func NewClient(apiURL string, apiKey string) *client {
	client := client{
		apiURL: apiURL,
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout:   10 * time.Second,
		},
	}

	return &client
}

// GetHosts crafts the API URL and invokes the `get` method to query the API returning a slice of `service` structs
// A filter can be defined to only return a specific subset of data; passing in a an empty filter returns all results
func (c *client) GetHosts(hf filterInterface) (*host, error) {

	url := fmt.Sprintf("%s%s%s%s%s%s%s%s", c.apiURL, apiV1Segment, hostSegment, "?", apiParam, "=", c.apiKey, hf.build())

	resp := hostResponse{}
	host := host{}

	json, err := c.get(url)

	if err != nil {
		return &host, err
	}

	resp.unmarshal(&json)
	host.unmarshal(&resp)

	return &host, nil
}

// GetServices crafts the API URL and invokes the `get` method to query the API returning a slice of `service` structs
// A filter can be defined to only return a specific subset of data; passing in a an empty filter returns all results
func (c *client) GetServices(sf filterInterface) (*service, error) {

	url := fmt.Sprintf("%s%s%s%s%s%s%s%s", c.apiURL, apiV1Segment, serviceSegment, "?", apiParam, "=", c.apiKey, sf.build())

	resp := serviceResponse{}
	service := service{}

	json, err := c.get(url)

	if err != nil {
		return &service, err
	}

	resp.unmarshal(&json)
	service.unmarshal(&resp)

	return &service, nil
}

// get performs a HTTP GET request to the API and returns the results in JSON format
func (c *client) get(url string) ([]byte, error) {

	req, err := c.httpClient.Get(url)

	if err != nil {
		return *new([]byte), fmt.Errorf("%s%s", "HTTP request failed", err)
	}

	defer req.Body.Close()

	json, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return *new([]byte), fmt.Errorf("%s%s", "unable to parse results", err)
	}

	return json, nil
}
