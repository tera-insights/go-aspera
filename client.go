package aspera

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	Client  *http.Client
	BaseURL *url.URL
}

func NewClient(httpClient *http.Client, baseUrl string) *Client {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil
	}
	return &Client{
		Client:  httpClient,
		BaseURL: url,
	}
}

func (c *Client) Do(ctx context.Context, req *http.Request, target interface{}) error {
	if ctx == nil {
		return errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)
	resp, err := c.Client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return fmt.Errorf("request cancelled: %w", ctx.Err())
		default:
		}
		// If the error type is *url.Error, sanitize its URL before returning.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = sanitizeURL(url).String()
				return e
			}
		}
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("response failed with code: %s", resp.Status)
	}

	// TODO: Handle error responses also

	// Read the response body
	err = json.NewDecoder(resp.Body).Decode(target)

	return err
}

func (c *Client) NewRequest(method, name string, body interface{}) (*http.Request, error) {
	endpoint, ok := endpoints[name]
	if !ok {
		return nil, fmt.Errorf("unknown endpoint %q", name)
	}

	request := &http.Request{
		Method: method,
		URL:    c.BaseURL.JoinPath(endpoint.URL()),
	}

	return request, nil
}

func (c *Client) NewRequestWithParameters(method, name string, params map[string]string, body interface{}) (*http.Request, error) {

	endpoint, ok := endpoints[name]
	if !ok {
		return nil, fmt.Errorf("unknown endpoint %q", name)
	}

	request := &http.Request{
		Method: method,
		URL:    c.BaseURL.JoinPath(endpoint.URLWithParams(params)),
	}

	return request, nil
}

// sanitizeURL redacts the client_secret parameter from the URL which may be exposed to the user.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return &url.URL{}
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}
