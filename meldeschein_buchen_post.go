package intenik

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewMeldescheinBuchenPostRequest() MeldescheinBuchenPostRequest {
	r := MeldescheinBuchenPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MeldescheinBuchenPostRequest struct {
	client      *Client
	queryParams *MeldescheinBuchenPostRequestQueryParams
	pathParams  *MeldescheinBuchenPostRequestPathParams
	method      string
	headers     http.Header
	requestBody MeldescheinBuchenPostRequestBody
}

func (r MeldescheinBuchenPostRequest) NewQueryParams() *MeldescheinBuchenPostRequestQueryParams {
	return &MeldescheinBuchenPostRequestQueryParams{}
}

type MeldescheinBuchenPostRequestQueryParams struct{}

func (p MeldescheinBuchenPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MeldescheinBuchenPostRequest) QueryParams() *MeldescheinBuchenPostRequestQueryParams {
	return r.queryParams
}

func (r MeldescheinBuchenPostRequest) NewPathParams() *MeldescheinBuchenPostRequestPathParams {
	return &MeldescheinBuchenPostRequestPathParams{}
}

type MeldescheinBuchenPostRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *MeldescheinBuchenPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *MeldescheinBuchenPostRequest) PathParams() *MeldescheinBuchenPostRequestPathParams {
	return r.pathParams
}

func (r *MeldescheinBuchenPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MeldescheinBuchenPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *MeldescheinBuchenPostRequest) Method() string {
	return r.method
}

func (r MeldescheinBuchenPostRequest) NewRequestBody() MeldescheinBuchenPostRequestBody {
	return MeldescheinBuchenPostRequestBody{}
}

type MeldescheinBuchenPostRequestBody struct {
	Meta         Meta         `json:"meta"`
	Meldescheine Meldescheine `json:"meldescheine"`
}

func (r *MeldescheinBuchenPostRequest) RequestBody() *MeldescheinBuchenPostRequestBody {
	return &r.requestBody
}

func (r *MeldescheinBuchenPostRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *MeldescheinBuchenPostRequest) SetRequestBody(body MeldescheinBuchenPostRequestBody) {
	r.requestBody = body
}

func (r *MeldescheinBuchenPostRequest) NewResponseBody() *MeldescheinBuchenPostResponseBody {
	return &MeldescheinBuchenPostResponseBody{}
}

type MeldescheinBuchenPostResponseBody struct {
	Meta         Meta   `json:"meta"`
	Status       string `json:"status"`
	Meldescheine []struct {
		ID          int `json:"id,omitempty"`
		Meldeschein `json:"meldeschein"`
	} `json:"meldescheine"`
}

func (r *MeldescheinBuchenPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/import/meldeschein/buchen/", r.PathParams())
	return &u
}

func (r *MeldescheinBuchenPostRequest) Do() (MeldescheinBuchenPostResponseBody, error, *http.Response) {
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
