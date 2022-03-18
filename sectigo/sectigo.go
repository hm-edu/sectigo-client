package sectigo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

const (
	defaultBaseURL = "https://cert-manager.com/api"
)

type Client struct {
	login       string
	password    string
	customerUri string

	// httpClient is the underlying HTTP client used to communicate with the API.
	httpClient *http.Client

	// BaseURL for API requests.
	BaseURL string

	Debug                   bool
	ClientService           *ClientService
	DomainService           *DomainService
	AcmeService             *AcmeService
	DomainValidationService *DomainValidationService
	SslService              *SslService
	PersonService           *PersonService
	OrganizationService     *OrganizationService
}

func NewClient(httpClient *http.Client, login, password, customerUri string) *Client {
	c := &Client{httpClient: httpClient, BaseURL: defaultBaseURL, login: login, password: password, customerUri: customerUri}
	c.ClientService = &ClientService{Client: c}
	c.DomainService = &DomainService{Client: c}
	c.AcmeService = &AcmeService{Client: c}
	c.DomainValidationService = &DomainValidationService{Client: c}
	c.SslService = &SslService{Client: c}
	c.PersonService = &PersonService{Client: c}
	c.OrganizationService = &OrganizationService{Client: c}
	return c
}

func Get[T any](c *Client, ctx context.Context, path string) (*T, *http.Response, error) {
	return makeRequest[T](c, ctx, http.MethodGet, path, nil, true)
}

func GetWithoutJsonResponse(c *Client, ctx context.Context, path string) (*http.Response, error) {
	_, resp, err := makeRequest[any](c, ctx, http.MethodGet, path, nil, false)
	return resp, err
}
func Post[T any](c *Client, ctx context.Context, path string, payload interface{}) (*T, *http.Response, error) {
	return makeRequest[T](c, ctx, http.MethodPost, path, payload, true)
}

func PostWithoutJsonResponse(c *Client, ctx context.Context, path string, payload interface{}) (*http.Response, error) {
	_, resp, err := makeRequest[any](c, ctx, http.MethodPost, path, payload, false)
	return resp, err
}

func Delete(c *Client, ctx context.Context, path string) (*http.Response, error) {
	_, resp, err := makeRequest[any](c, ctx, http.MethodDelete, path, nil, false)
	return resp, err
}

func sendRequestAndParse[T any](c *Client, ctx context.Context, req *http.Request, response bool) (*T, *http.Response, error) {
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
		err := Body.Close()
		if err != nil {
			log.Warn().Err(err)
		}
	}(resp.Body)
	obj := new(T)
	if obj != nil {
		err = json.NewDecoder(resp.Body).Decode(obj)
	}

	return obj, resp, err
}

type Response struct {
	// HTTP response
	HTTPResponse *http.Response
}

type ErrorResponse struct {
	Response
	// human-readable message
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
func makeRequest[T any](c *Client, ctx context.Context, method, path string, payload interface{}, response bool) (*T, *http.Response, error) {
	req, err := c.buildRequest(method, path, payload)
	if err != nil {
		return nil, nil, err
	}

	if c.Debug {
		log.Printf("Request (%v): %#v", req.URL, req)
	}

	obj, resp, err := sendRequestAndParse[T](c, ctx, req, response)
	if err != nil {
		return nil, nil, err
	}

	if c.Debug {
		log.Printf("Response: %#v", resp)
	}

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
	req.Header.Set("customerUri", c.customerUri)
	return req, err
}
