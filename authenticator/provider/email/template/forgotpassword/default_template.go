package forgotpassword

type TemplateHTML struct {
	HTMLTemplate string
	Name         string
}

func NewTemplateHTML() TemplateHTML {
	template := TemplateHTML{}
	template.HTMLTemplate = "<p>" +
		"You told us you forgot your password. " +
		"The code to reset your password is:" +
		"<strong>{{.Code}}</strong>" +
		"</p>"
	template.Name = "defaultTemplate.html"

	return template
}
