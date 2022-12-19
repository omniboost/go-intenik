package intenik

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewMeldescheineGetPostRequest() MeldescheineGetPostRequest {
	r := MeldescheineGetPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MeldescheineGetPostRequest struct {
	client      *Client
	queryParams *MeldescheineGetPostRequestQueryParams
	pathParams  *MeldescheineGetPostRequestPathParams
	method      string
	headers     http.Header
	requestBody MeldescheineGetPostRequestBody
}

func (r MeldescheineGetPostRequest) NewQueryParams() *MeldescheineGetPostRequestQueryParams {
	return &MeldescheineGetPostRequestQueryParams{}
}

type MeldescheineGetPostRequestQueryParams struct{}

func (p MeldescheineGetPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MeldescheineGetPostRequest) QueryParams() *MeldescheineGetPostRequestQueryParams {
	return r.queryParams
}

func (r MeldescheineGetPostRequest) NewPathParams() *MeldescheineGetPostRequestPathParams {
	return &MeldescheineGetPostRequestPathParams{}
}

type MeldescheineGetPostRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *MeldescheineGetPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *MeldescheineGetPostRequest) PathParams() *MeldescheineGetPostRequestPathParams {
	return r.pathParams
}

func (r *MeldescheineGetPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MeldescheineGetPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *MeldescheineGetPostRequest) Method() string {
	return r.method
}

func (r MeldescheineGetPostRequest) NewRequestBody() MeldescheineGetPostRequestBody {
	return MeldescheineGetPostRequestBody{}
}

type MeldescheineGetPostRequestBody struct {
	Meta         Meta  `json:"meta"`
	Meldescheine []int `json:"meldescheine"`
}

func (r *MeldescheineGetPostRequest) RequestBody() *MeldescheineGetPostRequestBody {
	return nil
}

func (r *MeldescheineGetPostRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MeldescheineGetPostRequest) SetRequestBody(body MeldescheineGetPostRequestBody) {
	r.requestBody = body
}

func (r *MeldescheineGetPostRequest) NewResponseBody() *MeldescheineGetPostResponseBody {
	return &MeldescheineGetPostResponseBody{}
}

type MeldescheineGetPostResponseBody struct {
}

func (r *MeldescheineGetPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/import/meldeschein/get/", r.PathParams())
	return &u
}

func (r *MeldescheineGetPostRequest) Do() (MeldescheineGetPostResponseBody, error, *http.Response) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	err = r.client.InitToken(req)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)
	return *responseBody, err, resp
}
