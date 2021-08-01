package oauth

// TokenValidation contains token validation data
type TokenValidation struct {
	Exp         int      `json:"exp"`
	Username    string   `json:"user_name"`
	Authorities []string `json:"authorities"`
	ClientID    string   `json:"client_id"`
	Scope       []string `json:"scope"`
}
