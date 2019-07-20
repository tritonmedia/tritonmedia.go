package tritonmedia

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// BaseURL is the URL used for all requests
	BaseURL = "https://app.tritonjs.com"
)

// Client is a client used for all requests to Triton Media
type Client struct {
	// BaseURL is the use used for all requests for this client
	BaseURL string

	// Client is the underlying HTTP client
	Client *http.Client

	// token is the API token used for authentication
	token string

	// throttler is a throttler implementation that allows us to throttle
	// requests on the client side
	throttler <-chan time.Time

	// ctx is the context used for this client
	ctx context.Context
}

// NewClient returns an instance of the TritonMedia client
func NewClient(baseURL, apiToken string) (*Client, error) {
	if baseURL == "" {
		baseURL = BaseURL
	}

	if apiToken == "" {
		return nil, fmt.Errorf("Missing an api token")
	}

	return &Client{
		Client:    http.DefaultClient,
		token:     apiToken,
		throttler: time.Tick(time.Second / 60),
		ctx:       context.Background(),
	}, nil
}

// waitForThrottle waits for the throttler to says it's OK to send a request
func (c *Client) waitForThrottle() {
	<-c.throttler
}

func (c *Client) errorParser(data []byte) error {
	var e *V1Error
	err := json.Unmarshal(data, &e)
	if err != nil {
		return nil
	}

	if !e.Success {
		return fmt.Errorf(e.Message)
	}

	return nil
}

// Get makes a GET request to a resource. Updates target interface with contents
func (c *Client) Get(path string, target interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.BaseURL, path), nil)
	if err != nil {
		return err
	}

	req = req.WithContext(c.ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %v", resp.StatusCode)
	}

	// TODO(jaredallard): investigate alternatives to reading the entire into mem twice
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	// TODO(jaredallard): merge this under the 200 status code check
	// check if the backend reported any errors
	err = c.errorParser(b)
	if err != nil {
		return fmt.Errorf("error reported by backend: %v", err)
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return nil
}

// Post makes a POST request to a resource. Updates target interface with contents
func (c *Client) Post(path string, data interface{}, target interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshall data to json: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.BaseURL, path), bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	req = req.WithContext(c.ctx)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code %v", resp.StatusCode)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %v", err)
	}

	// check if the backend reported any errors
	err = c.errorParser(b)
	if err != nil {
		return fmt.Errorf("error reported by backend: %v", err)
	}

	err = json.Unmarshal(b, target)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return nil
}
