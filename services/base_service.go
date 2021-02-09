package services

import (
	"context"
	"net/http"
	"net/url"
)

// BaseService provides the interface between client and services
type BaseService struct {
	Client IClient
}

//IClient implements these methods to become a type of IClient used by the BaseService
type IClient interface {
	BuildURLFromPathParams(url.Values, string, string, interface{}) (url.URL, error)
	Get(ctx context.Context, requestParams RequestParams) (*http.Response, error)
	Post(ctx context.Context, requestParams RequestParams) (*http.Response, error)
	Patch(ctx context.Context, requestParams RequestParams) (*http.Response, error)
	Put(ctx context.Context, requestParams RequestParams) (*http.Response, error)
	Delete(ctx context.Context, requestParams RequestParams) (*http.Response, error)
}

// RequestParams contains all the optional request URL parameters
type RequestParams struct {
	// Method is the HTTP method of the request
	Method string
	// URL is the URL of the request
	URL url.URL
	// Body is the body of the request
	Body interface{}
	// Headers are additional headers to add to the request
	Headers http.Header
}

//HeaderParamsParser is an interface to implement parsing any types of headers including user-defined that can be sent along with the api request
type HeaderParamsParser interface {
	Parse() (http.Header)
}
