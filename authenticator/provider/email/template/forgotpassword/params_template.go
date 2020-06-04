package forgotpassword

// TemplateParam represents the params used by the forgot_password template
type TemplateParam struct {
	Code string
}

func NewTempateParam(code string) *TemplateParam {
	param := &TemplateParam{Code: code}

	return param
}
