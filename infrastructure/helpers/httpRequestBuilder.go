package helpers

import (
	"fmt"
	"net/url"
)

func NewHttpRequestBuilder(baseUrl string) (HttpRequestBuilder, error)  {
	base, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return HttpRequestBuilder{}, fmt.Errorf("unable to parse base URL: %w", err)
	}
	if base.Scheme == "" {
		return HttpRequestBuilder{}, fmt.Errorf("no schem provided for base URL")
	}
	if base.Host == "" {
		return HttpRequestBuilder{}, fmt.Errorf("no host provided for base URL")
	}

	return HttpRequestBuilder{
		baseUrl:    base,
		params:     url.Values{},
	}, nil
}

type HttpRequestBuilder struct{
	baseUrl *url.URL
	params url.Values
}

func (request *HttpRequestBuilder) AddQueryParameter(key string, value string) {
	request.params.Add(key, value)
}

func (request *HttpRequestBuilder) GetUrl() string {
	request.baseUrl.RawQuery = request.params.Encode()
	return request.baseUrl.String()
}
