package rest

//TODO: not being used yet with the whole potential of nested structs because of a limitation on swaggo library.
type ResponseWrapper struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Messages []string    `json:"messages"`
	Result   interface{} `json:"result"`
}
