package intenik

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewTokenPostRequest() TokenPostRequest {
	r := TokenPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TokenPostRequest struct {
	client      *Client
	queryParams *TokenPostRequestQueryParams
	pathParams  *TokenPostRequestPathParams
	method      string
	headers     http.Header
	requestBody TokenPostRequestBody
}

func (r TokenPostRequest) NewQueryParams() *TokenPostRequestQueryParams {
	return &TokenPostRequestQueryParams{}
}

type TokenPostRequestQueryParams struct {
}

func (p TokenPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TokenPostRequest) QueryParams() *TokenPostRequestQueryParams {
	return r.queryParams
}

func (r TokenPostRequest) NewPathParams() *TokenPostRequestPathParams {
	return &TokenPostRequestPathParams{}
}

type TokenPostRequestPathParams struct {
}

func (p *TokenPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TokenPostRequest) PathParams() *TokenPostRequestPathParams {
	return r.pathParams
}

func (r *TokenPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TokenPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *TokenPostRequest) Method() string {
	return r.method
}

func (r TokenPostRequest) NewRequestBody() TokenPostRequestBody {
	return TokenPostRequestBody{}
}

type TokenPostRequestBody struct {
	// Username or e-mail address
	Username string `json:"username"`
	// the associated password
	Password string `json:"password"`
}

func (r *TokenPostRequest) RequestBody() *TokenPostRequestBody {
	return &r.requestBody
}

func (r *TokenPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *TokenPostRequest) SetRequestBody(body TokenPostRequestBody) {
	r.requestBody = body
}

func (r *TokenPostRequest) NewResponseBody() *TokenPostResponseBody {
	return &TokenPostResponseBody{}
}

type TokenPostResponseBody struct {
	Refresh string `json:"refresh"`
	Token   string `json:"token"`
}

func (r *TokenPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/token/", r.PathParams())
	return &u
}

func (r *TokenPostRequest) Do() (TokenPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
