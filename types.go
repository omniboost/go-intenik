package intenik

type Meta struct {
	Objekt     int    `json:"objekt"`
	Gemeinde   int    `json:"gemeinde"`
	RequestID  string `json:"requestId,omitempty"`
	Idempotent bool   `json:"idempotent,omitempty"`
}
