package sectigo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

const (
	defaultBaseURL = "https://cert-manager.com/api"
)

// Client is the main wrapper around the different sectigo services.
type Client struct {
	login       string
	password    string
	customerURI string
	logger      *zap.Logger

	// httpClient is the underlying HTTP client used to communicate with the API.
	httpClient *http.Client

	// BaseURL for API requests.
	BaseURL string

	ClientService           *ClientService
	DomainService           *DomainService
	AcmeService             *ACMEService
	DomainValidationService *DomainValidationService
	SslService              *SSLService
	PersonService           *PersonService
	OrganizationService     *OrganizationService
}

// NewClient creates a new client against the sectigo API using the provided credentials and the http.Client.
func NewClient(httpClient *http.Client, logger *zap.Logger, login, password, customerURI string) *Client {
	c := &Client{httpClient: httpClient, BaseURL: defaultBaseURL, login: login, password: password, customerURI: customerURI, logger: logger}
	c.ClientService = &ClientService{Client: c}
	c.DomainService = &DomainService{Client: c}
	c.AcmeService = &ACMEService{Client: c}
	c.DomainValidationService = &DomainValidationService{Client: c}
	c.SslService = &SSLService{Client: c}
	c.PersonService = &PersonService{Client: c}
	c.OrganizationService = &OrganizationService{Client: c}
	return c
}

// Get executes a GET-Requests and deserializes the returned JSON information using the provided type.
func Get[T any](ctx context.Context, c *Client, path string) (*T, *http.Response, error) {
	return makeRequest[T](ctx, c, http.MethodGet, path, nil, true)
}

// GetWithoutJSONResponse executes a GET-Request without expecting a JSON response.
// Custom handling of the response can be done using the returned http.Response.
func GetWithoutJSONResponse(ctx context.Context, c *Client, path string) (*http.Response, error) {
	_, resp, err := makeRequest[any](ctx, c, http.MethodGet, path, nil, false)
	return resp, err
}

// Post executes a POST-Requests and deserializes the returned JSON information using the provided type.
func Post[T any](ctx context.Context, c *Client, path string, payload interface{}) (*T, *http.Response, error) {
	return makeRequest[T](ctx, c, http.MethodPost, path, payload, true)
}

// PostWithoutJSONResponse executes a POST-Request without expecting a JSON response.
// Custom handling of the response can be done using the returned http.Response.
func PostWithoutJSONResponse(ctx context.Context, c *Client, path string, payload interface{}) (*http.Response, error) {
	_, resp, err := makeRequest[any](ctx, c, http.MethodPost, path, payload, false)
	return resp, err
}

// Delete executes a DELETE-Request and deserializes the returned JSON information using the provided type.
func Delete[T any](ctx context.Context, c *Client, path string, payload interface{}) (*T, *http.Response, error) {
	data, resp, err := makeRequest[T](ctx, c, http.MethodDelete, path, payload, true)
	return data, resp, err
}

// DeleteWithoutJSONResponse executes a DELETE-Request without expecting a JSON response.
// Custom handling of the response can be done using the returned http.Response.
func DeleteWithoutJSONResponse(ctx context.Context, c *Client, path string) (*http.Response, error) {
	_, resp, err := makeRequest[any](ctx, c, http.MethodDelete, path, nil, false)
	return resp, err
}

func sendRequestAndParse[T any](ctx context.Context, c *Client, req *http.Request, response bool) (*T, *http.Response, error) {
	if ctx == nil {
		return nil, nil, errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	err = checkResponse(resp)
	if err != nil {
		return nil, nil, err
	}
	if !response {
		return nil, resp, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	obj := new(T)
	if obj != nil {
		err = json.NewDecoder(resp.Body).Decode(obj)
	}

	return obj, resp, err
}

// Response is a wrapper around the normal http.Response.
type Response struct {
	// HTTP response.
	HTTPResponse *http.Response
}

// ErrorResponse provides the error information returned by sectigo.
type ErrorResponse struct {
	Response
	// human-readable message.
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

// Error implements the error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v %v %v", r.HTTPResponse.Request.Method, r.HTTPResponse.Request.URL, r.HTTPResponse.StatusCode, r.Code, r.Description)
}

func checkResponse(resp *http.Response) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	errorResponse := &ErrorResponse{Response: Response{
		HTTPResponse: resp,
	}}

	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

// makeRequest executes an API request and returns the HTTP response.
func makeRequest[T any](ctx context.Context, c *Client, method, path string, payload interface{}, response bool) (*T, *http.Response, error) {
	if c == nil {
		return nil, nil, fmt.Errorf("no client passed")
	}
	req, err := c.buildRequest(method, path, payload)
	if err != nil {
		c.logger.Warn("Building request failed", zap.Any("method", method), zap.Any("path", path), zap.Any("payload", payload))
		return nil, nil, err
	}

	c.logger.Debug("Request", zap.Any("url", req.URL), zap.Any("req", req))

	obj, resp, err := sendRequestAndParse[T](ctx, c, req, response)
	if err != nil {
		c.logger.Warn("Request failed", zap.Any("resp", err))
		return nil, nil, err
	}

	c.logger.Debug("Response", zap.Any("resp", resp))

	return obj, resp, nil
}

// buildRequest creates an API-Request.
func (c *Client) buildRequest(method, path string, payload interface{}) (*http.Request, error) {
	url := c.BaseURL + path

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Set("login", c.login)
	req.Header.Set("password", c.password)
	req.Header.Set("customerUri", c.customerURI)
	return req, err
}
