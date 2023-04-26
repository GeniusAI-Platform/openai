package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoFarsi/openai/errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
)

var _ Transporter = (*Client)(nil)

const (
	_default_base_url               = "https://api.openai.com/v1"
	_defaultEmptyMessagesLimit uint = 300
)

type Client struct {
	httpClient *http.Client
	validator  *validator.Validate
	rate       *rate.Limiter

	apiKey             string
	baseURL            string
	organizationID     string
	emptyMessagesLimit uint
}

type Response struct {
	resp *http.Response
}

type Transporter interface {
	GetClient() *http.Client
	GetValidator() *validator.Validate
	GetAPIKey() string
	GetOrganizationID() string
	Get(ctx context.Context, apiConfig *APIConfig) (*Response, error)
	Post(ctx context.Context, apiConfig *APIConfig, apiRequest any) (*Response, error)
}

// New create openai client
func New(apiKey string, opts ...Option) (Transporter, error) {
	if len(apiKey) == 0 {
		return nil, errors.ErrAPIKeyIsEmpty
	}

	client := &Client{
		validator:          validator.New(),
		apiKey:             apiKey,
		baseURL:            _default_base_url,
		emptyMessagesLimit: _defaultEmptyMessagesLimit,
	}

	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		client.httpClient = http.DefaultClient
	}

	return client, nil
}

func (c *Client) GetClient() *http.Client {
	return c.httpClient
}

func (c *Client) GetValidator() *validator.Validate {
	return c.validator
}

func (c *Client) GetAPIKey() string {
	return c.apiKey
}

func (c *Client) GetOrganizationID() string {
	return c.organizationID
}

// Get do get request and return response
func (c *Client) Get(ctx context.Context, apiConfig *APIConfig) (*Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, errors.New(http.StatusTooManyRequests, "", err.Error(), "", "")
	}

	req, err := http.NewRequest(GET.String(), c.baseURL+apiConfig.Path, nil)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "", err.Error(), "", "")
	}

	if len(apiConfig.Headers) != 0 {
		for k, v := range apiConfig.Headers {
			req.Header.Add(k, v)
		}
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	if len(apiConfig.Query) != 0 {
		req.URL.RawQuery = c.queryBuilder(apiConfig.Query)
	}

	if c.organizationID != "" {
		req.Header.Set("OpenAI-Organization", c.organizationID)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "", err.Error(), "", "")
	}

	return &Response{resp}, nil
}

// Post do post request and return response
func (c *Client) Post(ctx context.Context, apiConfig *APIConfig, apiRequest any) (*Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, errors.New(http.StatusTooManyRequests, "", err.Error(), "", "")
	}

	body, err := json.Marshal(apiRequest)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "", err.Error(), "", "")
	}

	url, err := url.JoinPath(c.baseURL, apiConfig.Path)
	if err != nil {
		return nil, errors.New(http.StatusTooManyRequests, "", err.Error(), "", "")
	}

	req, err := http.NewRequest(POST.String(), url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(http.StatusTooManyRequests, "", err.Error(), "", "")
	}

	if len(apiConfig.Headers) != 0 {
		for k, v := range apiConfig.Headers {
			req.Header.Add(k, v)
		}
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	if len(apiConfig.Query) != 0 {
		req.URL.RawQuery = c.queryBuilder(apiConfig.Query)
	}

	if c.organizationID != "" {
		req.Header.Set("OpenAI-Organization", c.organizationID)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, errors.New(http.StatusTooManyRequests, "", err.Error(), "", "")
	}

	return &Response{resp}, nil
}

func (c *Client) Stream(ctx context.Context, apiConfig *APIConfig, method Method, apiRequest any) (*Response, error) {
	if err := c.awaitRateLimiter(ctx); err != nil {
		return nil, err
	}

	body, err := json.Marshal(apiRequest)
	if err != nil {
		return nil, err
	}

	url, err := url.JoinPath(c.baseURL, apiConfig.Path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method.String(), url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	if c.organizationID != "" {
		req.Header.Set("OpenAI-Organization", c.organizationID)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}

	return &Response{resp}, nil
}

// GetJSON decode response body to your response
func (r *Response) GetJSON(response any) error {
	defer r.resp.Body.Close()
	if err := json.NewDecoder(r.resp.Body).Decode(response); err != nil {
		return errors.New(http.StatusInternalServerError, "", errors.ErrFailedToUnmarshalJSON.Error(), "", "")
	}
	return nil
}

// GetHttpResponse return http response
func (r *Response) GetHttpResponse() *http.Response {
	return r.resp
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}

	return client.Do(req.WithContext(ctx))
}

func (c *Client) awaitRateLimiter(ctx context.Context) error {
	if c.rate == nil {
		return nil
	}
	return c.rate.Wait(ctx)
}

func (c *Client) queryBuilder(params map[string]string) string {
	query := url.Values{}
	for k, v := range params {
		query[k] = []string{v}
	}
	return query.Encode()
}
