package intenik

import "time"

type Meta struct {
	Objekt     int    `json:"objekt"`
	Gemeinde   int    `json:"gemeinde"`
	RequestID  string `json:"requestId,omitempty"`
	Idempotent bool   `json:"idempotent,omitempty"`
}

type Meldescheine []Meldeschein

type Meldeschein struct {
	ID       int      `json:"id,omitempty"`
	Personen Personen `json:"personen"`
	// TravelSpan TravelSpan `json:"travel_span"`
	// Objekt     Objekt     `json:"objekt"`
	MSType string `json:"ms_type"`

	BeherbergerID         int         `json:"beherberger_id,omitempty"`
	VermieterAbweichendID interface{} `json:"vermieter_abweichend_id,omitempty"`
	ObjektAbweichendID    interface{} `json:"objekt_abweichend_id,omitempty"`
	VerwendeteKurtaxe     int         `json:"verwendete_kurtaxe,omitempty"`
	// CreationUser          struct {
	// 	ID       int    `json:"id,omitempty"`
	// 	Username string `json:"username,omitempty"`
	// 	Alias    string `json:"alias,omitempty"`
	// } `json:"creation_user,omitempty"`
	PrecheckErrors []interface{} `json:"precheck_errors,omitempty"`
	Invoice        interface{}   `json:"invoice,omitempty"`
	StateDisplay   string        `json:"state_display,omitempty"`
	FakturaDate    interface{}   `json:"faktura_date,omitempty"`
	Csid           string        `json:"CSID,omitempty"`
	FormID         interface{}   `json:"Form_Id,omitempty"`
	BatchNo        interface{}   `json:"BatchNo,omitempty"`
	BatchRDate     string        `json:"BatchRDate,omitempty"`
	BatchTrack     interface{}   `json:"BatchTrack,omitempty"`
	State          string        `json:"state,omitempty"`
	// Created             time.Time     `json:"created,omitempty"`
	UserNote            interface{} `json:"user_note,omitempty"`
	IsStorniert         bool        `json:"is_storniert,omitempty"`
	Source              interface{} `json:"source,omitempty"`
	StornoBestaetigt    bool        `json:"storno_bestaetigt,omitempty"`
	KurtaxeIst          interface{} `json:"kurtaxe_ist,omitempty"`
	ManuellerBetrag     int         `json:"manueller_betrag,omitempty"`
	KzManuellerBetrag   bool        `json:"kz_manueller_betrag,omitempty"`
	KurtaxeCalc         int         `json:"kurtaxe_calc,omitempty"`
	StatAlterKind1      interface{} `json:"stat_alter_kind_1,omitempty"`
	StatAlterKind2      interface{} `json:"stat_alter_kind_2,omitempty"`
	StatAlterKind3      interface{} `json:"stat_alter_kind_3,omitempty"`
	StatAlterKind4      interface{} `json:"stat_alter_kind_4,omitempty"`
	VermieterAbweichend interface{} `json:"vermieter_abweichend,omitempty"`
	ObjektAbweichend    interface{} `json:"objekt_abweichend,omitempty"`
}

type Personen []Person

type Person struct {
	ID                   int    `json:"id,omitempty"`
	TarifID              int    `json:"tarif_id,omitempty"`
	Anzahl               int    `json:"anzahl,omitempty"`
	GastType             string `json:"gast_type,omitempty"`
	ArrivalDate          string `json:"arrival_date,omitempty"`
	DepartureDate        string `json:"departure_date,omitempty"`
	GastVorname          string `json:"Gast_Vorname,omitempty"`
	GastName             string `json:"Gast_Name,omitempty"`
	Staatsangehoerigkeit string `json:"Gast_Staatsangehoerigkeit,omitempty"`

	GastGeburtsdatum string `json:"Gast_Geburtsdatum,omitempty"`
	GastStrasse      string `json:"Gast_Strasse,omitempty"`
	GastPostleitzahl string `json:"Gast_Postleitzahl,omitempty"`
	GastWohnort      string `json:"Gast_Wohnort,omitempty"`
	GastLand         string `json:"Gast_Land,omitempty"`
	GastEmail        string `json:"Gast_Email,omitempty"`

	AddressID    interface{} `json:"address_id,omitempty"`
	AddressLat   interface{} `json:"address_lat,omitempty"`
	AddressLng   interface{} `json:"address_lng,omitempty"`
	AddressLabel interface{} `json:"address_label,omitempty"`

	Delete bool `json:"delete,omitempty"`
}

type TravelSpan struct {
	ArrivalDate        string `json:"arrival_date"`
	DepartureDate      string `json:"departure_date"`
	TravelDays         int    `json:"travel_days"`
	UebernachtungenSum int    `json:"uebernachtungen_sum"`
}

type Objekt struct {
	ID            int `json:"id"`
	BeherbergerID int `json:"beherberger_id"`
	Beherberger   struct {
		Firma           string `json:"firma"`
		Vorname         string `json:"vorname"`
		Nachname        string `json:"nachname"`
		Comment         string `json:"comment"`
		DebitorNr       string `json:"debitor_nr"`
		Alias           string `json:"alias"`
		ID              int    `json:"id"`
		Ansprechpartner int    `json:"ansprechpartner"`
		GemeindeID      int    `json:"gemeinde_id"`
	} `json:"beherberger"`
	BeherbergerAlias string `json:"beherberger_alias"`
	Ortsteil         struct {
		ID       int    `json:"id"`
		Alias    string `json:"alias"`
		Gemeinde int    `json:"gemeinde"`
	} `json:"ortsteil"`
	Created        time.Time `json:"created"`
	Modified       time.Time `json:"modified"`
	AnzahlBetten   int       `json:"anzahl_betten"`
	Alias          string    `json:"alias"`
	Strasse        string    `json:"strasse"`
	Plz            string    `json:"plz"`
	Ort            string    `json:"ort"`
	Tarifzone      int       `json:"tarifzone"`
	Unterkunftsart int       `json:"unterkunftsart"`
}
