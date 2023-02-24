package intenik

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewMeldescheinAktualisierenPostRequest() MeldescheinAktualisierenPostRequest {
	r := MeldescheinAktualisierenPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MeldescheinAktualisierenPostRequest struct {
	client      *Client
	queryParams *MeldescheinAktualisierenPostRequestQueryParams
	pathParams  *MeldescheinAktualisierenPostRequestPathParams
	method      string
	headers     http.Header
	requestBody MeldescheinAktualisierenPostRequestBody
}

func (r MeldescheinAktualisierenPostRequest) NewQueryParams() *MeldescheinAktualisierenPostRequestQueryParams {
	return &MeldescheinAktualisierenPostRequestQueryParams{}
}

type MeldescheinAktualisierenPostRequestQueryParams struct{}

func (p MeldescheinAktualisierenPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MeldescheinAktualisierenPostRequest) QueryParams() *MeldescheinAktualisierenPostRequestQueryParams {
	return r.queryParams
}

func (r MeldescheinAktualisierenPostRequest) NewPathParams() *MeldescheinAktualisierenPostRequestPathParams {
	return &MeldescheinAktualisierenPostRequestPathParams{}
}

type MeldescheinAktualisierenPostRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *MeldescheinAktualisierenPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *MeldescheinAktualisierenPostRequest) PathParams() *MeldescheinAktualisierenPostRequestPathParams {
	return r.pathParams
}

func (r *MeldescheinAktualisierenPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MeldescheinAktualisierenPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *MeldescheinAktualisierenPostRequest) Method() string {
	return r.method
}

func (r MeldescheinAktualisierenPostRequest) NewRequestBody() MeldescheinAktualisierenPostRequestBody {
	return MeldescheinAktualisierenPostRequestBody{}
}

type MeldescheinAktualisierenPostRequestBody struct {
	Meta         Meta         `json:"meta"`
	Meldescheine Meldescheine `json:"meldescheine"`
}

func (r *MeldescheinAktualisierenPostRequest) RequestBody() *MeldescheinAktualisierenPostRequestBody {
	return &r.requestBody
}

func (r *MeldescheinAktualisierenPostRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *MeldescheinAktualisierenPostRequest) SetRequestBody(body MeldescheinAktualisierenPostRequestBody) {
	r.requestBody = body
}

func (r *MeldescheinAktualisierenPostRequest) NewResponseBody() *MeldescheinAktualisierenPostResponseBody {
	return &MeldescheinAktualisierenPostResponseBody{}
}

type MeldescheinAktualisierenPostResponseBody struct {
	Meta     Meta   `json:"meta"`
	Status   string `json:"status"`
	Response struct {
		Meldescheine []Meldeschein `json:"meldescheine"`
	} `json:"response"`
}

func (r *MeldescheinAktualisierenPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/import/meldeschein/aktualisieren/", r.PathParams())
	return &u
}

func (r *MeldescheinAktualisierenPostRequest) Do() (MeldescheinAktualisierenPostResponseBody, error, *http.Response) {
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
