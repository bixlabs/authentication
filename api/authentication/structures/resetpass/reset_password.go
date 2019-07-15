package passreset

//TODO: We could use nested struct promoted fields here but the swaggo
// library is not generating correct documentation using that.
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
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"newPassword"`
}
