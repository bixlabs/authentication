package structures

type LoginResponse struct {
	Token      string `json:"token"`
	IssuedAt   int    `json:"iat"`
	Expiration int    `json:"exp"`
}
