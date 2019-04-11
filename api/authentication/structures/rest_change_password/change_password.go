package rest_change_password

//TODO: We could use nested struct promoted fields here but the swaggo library is not generating correct documentation using that.
type Response struct {
	Status   string   `json:"status"`
	Code     int      `json:"code"`
	Messages []string `json:"messages"`
	Result   Result   `json:"result"`
}

type Result struct {
	Success bool `json:"success"`
}

type Request struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
