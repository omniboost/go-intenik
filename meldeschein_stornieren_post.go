package intenik

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewMeldescheinStornierenPostRequest() MeldescheinStornierenPostRequest {
	r := MeldescheinStornierenPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MeldescheinStornierenPostRequest struct {
	client      *Client
	queryParams *MeldescheinStornierenPostRequestQueryParams
	pathParams  *MeldescheinStornierenPostRequestPathParams
	method      string
	headers     http.Header
	requestBody MeldescheinStornierenPostRequestBody
}

func (r MeldescheinStornierenPostRequest) NewQueryParams() *MeldescheinStornierenPostRequestQueryParams {
	return &MeldescheinStornierenPostRequestQueryParams{}
}

type MeldescheinStornierenPostRequestQueryParams struct{}

func (p MeldescheinStornierenPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MeldescheinStornierenPostRequest) QueryParams() *MeldescheinStornierenPostRequestQueryParams {
	return r.queryParams
}

func (r MeldescheinStornierenPostRequest) NewPathParams() *MeldescheinStornierenPostRequestPathParams {
	return &MeldescheinStornierenPostRequestPathParams{}
}

type MeldescheinStornierenPostRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *MeldescheinStornierenPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *MeldescheinStornierenPostRequest) PathParams() *MeldescheinStornierenPostRequestPathParams {
	return r.pathParams
}

func (r *MeldescheinStornierenPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MeldescheinStornierenPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *MeldescheinStornierenPostRequest) Method() string {
	return r.method
}

func (r MeldescheinStornierenPostRequest) NewRequestBody() MeldescheinStornierenPostRequestBody {
	return MeldescheinStornierenPostRequestBody{}
}

type MeldescheinStornierenPostRequestBody struct {
	Meta         Meta   `json:"meta"`
	Status       string `json:"status"`
	Meldescheine []int  `json:"meldescheine"`
}

func (r *MeldescheinStornierenPostRequest) RequestBody() *MeldescheinStornierenPostRequestBody {
	return &r.requestBody
}

func (r *MeldescheinStornierenPostRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *MeldescheinStornierenPostRequest) SetRequestBody(body MeldescheinStornierenPostRequestBody) {
	r.requestBody = body
}

func (r *MeldescheinStornierenPostRequest) NewResponseBody() *MeldescheinStornierenPostResponseBody {
	return &MeldescheinStornierenPostResponseBody{}
}

type MeldescheinStornierenPostResponseBody struct {
	Meta         Meta  `json:"meta"`
	Meldescheine []int `json:"meldescheine"`
}

func (r *MeldescheinStornierenPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/import/meldeschein/stornieren/", r.PathParams())
	return &u
}

func (r *MeldescheinStornierenPostRequest) Do() (MeldescheinStornierenPostResponseBody, error, *http.Response) {
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
