package oauth

// TokenValidation contains token validation data
type TokenValidation struct {
	Scope       []interface{} `json:"scope"`
	Exp         int           `json:"exp"`
	Authorities []struct {
		Authority string `json:"authority"`
	} `json:"authorities"`
	ClientID string `json:"client_id"`
}
