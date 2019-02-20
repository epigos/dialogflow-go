package dialogflow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	// DefaultVersion default api version
	DefaultVersion = "v2"
	// APIBase API base url
	APIBase = "https://dialogflow.googleapis.com"
)

// Client - Dialogflow client type
type Client struct {
	accessToken string
	projectID   string
	apiVersion  string
	apiURL      string
	httpClient  *http.Client
}

type errorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type errorResp struct {
	Error errorBody `json:"error"`
}

type httpParams struct {
	method   string
	resource string
	body     interface{}
}

// Operation represents the result of a network API call to dialogflow
type Operation struct {
	Name      string                 `json:"name"`
	Done      bool                   `json:"done"`
	Metadata  map[string]interface{} `json:"metadata"`
	Respoonse map[string]interface{} `json:"response"`
}

// NewClient - returns dialogflow.com client for API version 2
func NewClient(token, projectID string) *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	apiURL := fmt.Sprintf("%s/%s/projects/%s", APIBase, DefaultVersion, projectID)

	return &Client{
		accessToken: token,
		projectID:   projectID,
		apiVersion:  DefaultVersion,
		apiURL:      apiURL,
		httpClient:  httpClient,
	}
}

// GetAccessToken dialogflow.com access token
func (c *Client) GetAccessToken() string {
	return c.accessToken
}

// GetAPIVersion dialogflow.com version
func (c *Client) GetAPIVersion() string {
	if c.apiVersion != "" {
		return c.apiVersion
	}
	return DefaultVersion
}

// GetProjectID dialogflow project that the agent is associated with.
func (c *Client) GetProjectID() string {
	return c.projectID
}

func (c *Client) request(method, resource string, body io.Reader) (io.ReadCloser, error) {
	url := c.apiURL + resource
	log.Printf("Making request %s to %s\n", method, url)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.GetAccessToken())
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		defer resp.Body.Close()

		var e *errorResp
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&e)
		if err != nil {
			return nil, errors.New("Internal Error")
		}

		return nil, errors.New(e.Error.Message)
	}

	return resp.Body, nil
}

func (c *Client) processRequest(p *httpParams, dst interface{}) error {
	body, err := json.Marshal(p.body)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(body)
	resp, err := c.request(p.method, p.resource, reader)

	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp)
	if err = decoder.Decode(dst); err != nil {
		return err
	}

	return nil
}

func (c *Client) get(resource string, dst interface{}) error {
	return c.processRequest(&httpParams{
		method:   http.MethodGet,
		resource: resource,
	}, dst)
}

func (c *Client) post(resource string, body interface{}, dst interface{}) error {
	return c.processRequest(&httpParams{
		method:   http.MethodPost,
		resource: resource,
		body:     body,
	}, dst)
}

func (c *Client) put(resource string, body interface{}, dst interface{}) error {
	return c.processRequest(&httpParams{
		method:   http.MethodPut,
		resource: resource,
		body:     body,
	}, dst)
}

func (c *Client) delete(resource string) error {
	return c.processRequest(&httpParams{
		method:   http.MethodDelete,
		resource: resource,
	}, nil)
}
