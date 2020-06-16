package implementation

import "github.com/bixlabs/authentication/authenticator/provider/email/template/structures"

// TemplateParam represents the params used by the forgot_password template
type TemplateParam struct {
	Code string
}

func NewTempateParam(code string) structures.ParamsTemplate {
	param := TemplateParam{Code: code}

	return param
}

func (tpl TemplateParam) GetParams() []string {
	return []string{"{{.Code}}"}
}
