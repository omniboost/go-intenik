package intenik

import (
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-intenik/utils"
)

func (c *Client) NewGetCurrentUserGetRequest() GetCurrentUserGetRequest {
	r := GetCurrentUserGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type GetCurrentUserGetRequest struct {
	client      *Client
	queryParams *GetCurrentUserGetRequestQueryParams
	pathParams  *GetCurrentUserGetRequestPathParams
	method      string
	headers     http.Header
	requestBody GetCurrentUserGetRequestBody
}

func (r GetCurrentUserGetRequest) NewQueryParams() *GetCurrentUserGetRequestQueryParams {
	return &GetCurrentUserGetRequestQueryParams{}
}

type GetCurrentUserGetRequestQueryParams struct{}

func (p GetCurrentUserGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetCurrentUserGetRequest) QueryParams() *GetCurrentUserGetRequestQueryParams {
	return r.queryParams
}

func (r GetCurrentUserGetRequest) NewPathParams() *GetCurrentUserGetRequestPathParams {
	return &GetCurrentUserGetRequestPathParams{}
}

type GetCurrentUserGetRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *GetCurrentUserGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *GetCurrentUserGetRequest) PathParams() *GetCurrentUserGetRequestPathParams {
	return r.pathParams
}

func (r *GetCurrentUserGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetCurrentUserGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetCurrentUserGetRequest) Method() string {
	return r.method
}

func (r GetCurrentUserGetRequest) NewRequestBody() GetCurrentUserGetRequestBody {
	return GetCurrentUserGetRequestBody{}
}

type GetCurrentUserGetRequestBody struct {
}

func (r *GetCurrentUserGetRequest) RequestBody() *GetCurrentUserGetRequestBody {
	return nil
}

func (r *GetCurrentUserGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetCurrentUserGetRequest) SetRequestBody(body GetCurrentUserGetRequestBody) {
	r.requestBody = body
}

func (r *GetCurrentUserGetRequest) NewResponseBody() *GetCurrentUserGetResponseBody {
	return &GetCurrentUserGetResponseBody{}
}

type GetCurrentUserGetResponseBody struct {
	User struct {
		ID                 int       `json:"id"`
		Username           string    `json:"username"`
		Email              string    `json:"email"`
		Alias              string    `json:"alias"`
		FirstName          string    `json:"first_name"`
		LastName           string    `json:"last_name"`
		UserType           string    `json:"user_type"`
		IsActive           bool      `json:"is_active"`
		LastLogin          time.Time `json:"last_login"`
		IsStaff            bool      `json:"is_staff"`
		IsSuperuser        bool      `json:"is_superuser"`
		MustChangePassword bool      `json:"must_change_password"`
		HasAPIAccess       bool      `json:"has_api_access"`
		DateJoined         time.Time `json:"date_joined"`
		PasswordChangedAt  time.Time `json:"password_changed_at"`
	} `json:"user"`
	Beherberger []struct {
		ID      int `json:"id"`
		Objekte []struct {
			ID            int `json:"id"`
			BeherbergerID int `json:"beherberger_id"`
			Beherberger   struct {
				ID             int    `json:"id"`
				Alias          string `json:"alias"`
				CreateInvoices bool   `json:"create_invoices"`
			} `json:"beherberger"`
			Alias            string `json:"alias"`
			BeherbergerAlias string `json:"beherberger_alias"`
		} `json:"objekte"`
		Alias           string `json:"alias"`
		Gemeinde        int    `json:"gemeinde"`
		Ansprechpartner struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Alias    string `json:"alias"`
		} `json:"ansprechpartner"`
		AnsprechpartnerID int `json:"ansprechpartner_id"`
		Ortsteil          struct {
			ID       int    `json:"id"`
			Alias    string `json:"alias"`
			Gemeinde int    `json:"gemeinde"`
		} `json:"ortsteil"`
		OrtsteilID int `json:"ortsteil_id"`
	} `json:"beherberger"`
	Gemeinden []struct {
		ID        int `json:"id"`
		Ortsteile []struct {
			ID       int    `json:"id"`
			Alias    string `json:"alias"`
			Gemeinde int    `json:"gemeinde"`
		} `json:"ortsteile"`
		Unterkunftsarten []struct {
			ID       int    `json:"id"`
			Alias    string `json:"alias"`
			Gemeinde int    `json:"gemeinde"`
		} `json:"unterkunftsarten"`
		Tarifzonen []struct {
			ID       int       `json:"id"`
			Created  time.Time `json:"created"`
			Modified time.Time `json:"modified"`
			Alias    string    `json:"alias"`
			Gemeinde int       `json:"gemeinde"`
		} `json:"tarifzonen"`
		Saisonzeiten []struct {
			ID         int       `json:"id"`
			Created    time.Time `json:"created"`
			Modified   time.Time `json:"modified"`
			Alias      string    `json:"alias"`
			ValidFrom  string    `json:"valid_from"`
			ValidUntil string    `json:"valid_until"`
			Gemeinde   int       `json:"gemeinde"`
		} `json:"saisonzeiten"`
		Tarife []struct {
			ID                  int       `json:"id"`
			FullAlias           string    `json:"full_alias"`
			Created             time.Time `json:"created"`
			Modified            time.Time `json:"modified"`
			IsActive            bool      `json:"is_active"`
			TarifCode           string    `json:"tarif_code"`
			Alias               string    `json:"alias"`
			AllowBeherberger    bool      `json:"allow_beherberger"`
			PrintKurkarte       bool      `json:"print_kurkarte"`
			DefaultSelected     bool      `json:"default_selected"`
			Gemeinde            int       `json:"gemeinde"`
			MeldescheinTemplate int       `json:"meldeschein_template"`
		} `json:"tarife"`
		MeldescheinFields []struct {
			ID             int         `json:"id"`
			Created        time.Time   `json:"created"`
			Modified       time.Time   `json:"modified"`
			Alias          string      `json:"alias"`
			ColumnName     string      `json:"column_name"`
			DataType       string      `json:"data_type"`
			Size           int         `json:"size"`
			Transferable   bool        `json:"transferable"`
			ListVisible    bool        `json:"list_visible"`
			Required       bool        `json:"required"`
			Priority       string      `json:"priority"`
			ViewOrder      int         `json:"view_order"`
			DetailSearch   bool        `json:"detail_search"`
			FieldType      string      `json:"field_type"`
			FulltextSearch bool        `json:"fulltext_search"`
			ShowInDetail   bool        `json:"show_in_detail"`
			HintText       interface{} `json:"hint_text"`
			Gemeinde       int         `json:"gemeinde"`
		} `json:"meldeschein_fields"`
		Settings struct {
			Gemeinde    int `json:"gemeinde"`
			ErrorConfig struct {
				MsNotFound struct {
					Color     string `json:"color"`
					Icon      string `json:"icon"`
					CanIgnore bool   `json:"canIgnore"`
				} `json:"MS_NOT_FOUND"`
				KurtaxeIstMissing struct {
					Color     string `json:"color"`
					Icon      string `json:"icon"`
					CanIgnore bool   `json:"canIgnore"`
				} `json:"KURTAXE_IST_MISSING"`
			} `json:"error_config"`
			Created                        time.Time `json:"created"`
			Modified                       time.Time `json:"modified"`
			LetterheadSender               string    `json:"letterhead_sender"`
			LetterheadRightSide            string    `json:"letterhead_right_side"`
			LetterheadFooter               string    `json:"letterhead_footer"`
			LetterheadExtra                string    `json:"letterhead_extra"`
			InvoiceFooter                  string    `json:"invoice_footer"`
			InvoiceFooterWithLastschrift   string    `json:"invoice_footer_with_lastschrift"`
			InvoiceFooterCredit            string    `json:"invoice_footer_credit"`
			InvoiceNumberTemplate          string    `json:"invoice_number_template"`
			InvoiceNumberCurrent           int       `json:"invoice_number_current"`
			CalcStartOffset                int       `json:"calc_start_offset"`
			PrecheckAllowHigherTax         bool      `json:"precheck_allow_higher_tax"`
			PrecheckAllowZeroTax           bool      `json:"precheck_allow_zero_tax"`
			PrecheckAllowMissingKurtaxeIst bool      `json:"precheck_allow_missing_kurtaxe_ist"`
			EmailUserCredentialsFooter     string    `json:"email_user_credentials_footer"`
			AusgabebelegAdditionalText     string    `json:"ausgabebeleg_additional_text"`
			AusgabebelegRightSide          string    `json:"ausgabebeleg_right_side"`
			InvoiceShowTax                 bool      `json:"invoice_show_tax"`
			InvoiceTaxPercentage           int       `json:"invoice_tax_percentage"`
			InvoiceTaxString               string    `json:"invoice_tax_string"`
			InvoiceRestforderungString     string    `json:"invoice_restforderung_string"`
			InvoiceTitle                   string    `json:"invoice_title"`
			InvoiceStornoTitle             string    `json:"invoice_storno_title"`
			InvoiceHasAccountingExport     bool      `json:"invoice_has_accounting_export"`
			InvoiceTemplateHTML            string    `json:"invoice_template_html"`
			InvoiceRunSortBy               string    `json:"invoice_run_sort_by"`
		} `json:"settings"`
		Created                time.Time   `json:"created"`
		Modified               time.Time   `json:"modified"`
		Alias                  string      `json:"alias"`
		LetterheadTemplateHTML string      `json:"letterhead_template_html"`
		UsingPaper             bool        `json:"using_paper"`
		UsingEms               bool        `json:"using_ems"`
		DatatableName          string      `json:"datatable_name"`
		IsActive               bool        `json:"is_active"`
		ModelClass             string      `json:"model_class"`
		Gemeindeschluessel     interface{} `json:"gemeindeschluessel"`
		OverrideImagePath      interface{} `json:"override_image_path"`
	} `json:"gemeinden"`
	Deleted  interface{} `json:"deleted"`
	Created  time.Time   `json:"created"`
	Modified time.Time   `json:"modified"`
}

func (r *GetCurrentUserGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/users/get_current_user/", r.PathParams())
	return &u
}

func (r *GetCurrentUserGetRequest) Do() (GetCurrentUserGetResponseBody, error, *http.Response) {
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
