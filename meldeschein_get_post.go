package intenik

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewMeldescheinGetPostRequest() MeldescheinGetPostRequest {
	r := MeldescheinGetPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MeldescheinGetPostRequest struct {
	client      *Client
	queryParams *MeldescheinGetPostRequestQueryParams
	pathParams  *MeldescheinGetPostRequestPathParams
	method      string
	headers     http.Header
	requestBody MeldescheinGetPostRequestBody
}

func (r MeldescheinGetPostRequest) NewQueryParams() *MeldescheinGetPostRequestQueryParams {
	return &MeldescheinGetPostRequestQueryParams{}
}

type MeldescheinGetPostRequestQueryParams struct{}

func (p MeldescheinGetPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MeldescheinGetPostRequest) QueryParams() *MeldescheinGetPostRequestQueryParams {
	return r.queryParams
}

func (r MeldescheinGetPostRequest) NewPathParams() *MeldescheinGetPostRequestPathParams {
	return &MeldescheinGetPostRequestPathParams{}
}

type MeldescheinGetPostRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *MeldescheinGetPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *MeldescheinGetPostRequest) PathParams() *MeldescheinGetPostRequestPathParams {
	return r.pathParams
}

func (r *MeldescheinGetPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MeldescheinGetPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *MeldescheinGetPostRequest) Method() string {
	return r.method
}

func (r MeldescheinGetPostRequest) NewRequestBody() MeldescheinGetPostRequestBody {
	return MeldescheinGetPostRequestBody{}
}

type MeldescheinGetPostRequestBody struct {
	Meta         Meta  `json:"meta"`
	Meldescheine []int `json:"meldescheine"`
}

func (r *MeldescheinGetPostRequest) RequestBody() *MeldescheinGetPostRequestBody {
	return &r.requestBody
}

func (r *MeldescheinGetPostRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *MeldescheinGetPostRequest) SetRequestBody(body MeldescheinGetPostRequestBody) {
	r.requestBody = body
}

func (r *MeldescheinGetPostRequest) NewResponseBody() *MeldescheinGetPostResponseBody {
	return &MeldescheinGetPostResponseBody{}
}

type MeldescheinGetPostResponseBody struct {
}

func (r *MeldescheinGetPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/import/meldeschein/get/", r.PathParams())
	return &u
}

func (r *MeldescheinGetPostRequest) Do() (MeldescheinGetPostResponseBody, error, *http.Response) {
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
