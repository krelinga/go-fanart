package fanart

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/krelinga/go-jsonflex"
)

type Client interface {
	GetObject(ctx context.Context, path string, options ...RequestOption) (jsonflex.Object, error)
}

type ClientOptions struct {
	APIKey     string
	HttpClient *http.Client
}

func NewClient(options ClientOptions) Client {
	if options.HttpClient == nil {
		options.HttpClient = http.DefaultClient
	}
	return &clientImpl{
		options: options,
	}
}

type HttpStatusCodeError struct {
	StatusCode int
}

func (e HttpStatusCodeError) Error() string {
	return fmt.Sprintf("unexpected status code: %d", e.StatusCode)
}

type clientImpl struct {
	options ClientOptions
}

func (c *clientImpl) getRaw(ctx context.Context, path string, options ...RequestOption) (io.ReadCloser, error) {
	if c.options.APIKey != "" {
		options = append(options, WithQueryParam("api_key", c.options.APIKey))
	}
	urlValues := url.Values{}
	for _, opt := range options {
		if opt.ChangeValues != nil {
			opt.ChangeValues(&urlValues)
		}
	}
	reqUrl := &url.URL{
		Scheme:   "http",
		Host:     "webservice.fanart.tv",
		Path:     path,
		RawQuery: urlValues.Encode(),
	}
	reqHeader := http.Header{}
	for _, opt := range options {
		if opt.ChangeHeader != nil {
			opt.ChangeHeader(&reqHeader)
		}
	}
	req := &http.Request{
		Method: http.MethodGet,
		URL:    reqUrl,
		Header: reqHeader,
	}
	for _, opt := range options {
		if opt.ChangeRequest != nil {
			opt.ChangeRequest(req)
		}
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	response, err := c.options.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	for _, opt := range options {
		if opt.ChangeResponse != nil {
			opt.ChangeResponse(response)
		}
	}
	if response.StatusCode != http.StatusOK {
		return nil, HttpStatusCodeError{StatusCode: response.StatusCode}
	}
	contentType := response.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return nil, fmt.Errorf("unexpected content type: %s", contentType)
	}
	return response.Body, nil
}

func (c *clientImpl) GetObject(ctx context.Context, path string, options ...RequestOption) (jsonflex.Object, error) {
	body, err := c.getRaw(ctx, path, options...)
	if err != nil {
		return nil, err
	}
	defer body.Close()
	o := jsonflex.Object{}
	if err := json.NewDecoder(body).Decode(&o); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return o, nil
}
