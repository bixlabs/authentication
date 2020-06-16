package implementation

import "github.com/bixlabs/authentication/authenticator/provider/email/template/structures"

type TemplateHTML struct {
	HTMLTemplate string
	Name         string
}

func NewTemplateHTML() structures.DefaultTemplate {
	template := TemplateHTML{}
	template.HTMLTemplate = "<p>" +
		"You told us you forgot your password. " +
		"The code to reset your password is:" +
		"<strong>{{.Code}}</strong>" +
		"</p>"
	template.Name = "defaultTemplate.html"

	return template
}

func (tpl TemplateHTML) GetTemplate() (string, string) {
	return tpl.HTMLTemplate, tpl.Name
}
