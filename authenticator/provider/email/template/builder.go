package template

import (
	"bytes"
	htmlParser "html/template"
	"io/ioutil"
	"strings"

	forgotPass "github.com/bixlabs/authentication/authenticator/provider/email/template/forgotpassword"

	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	// autoload is not working in main
	//_ "github.com/joho/godotenv/autoload"
)

// Builder represents a template loader and builder for the emails
type Builder struct {
	TemplatePath    string `env:"AUTH_SERVER_EMAIL_TEMPLATE_PATH"`
	DefaultTemplate forgotPass.TemplateHTML
	CustomTemplate  string
	CustomName      string
}

func NewTemplateBuilder(defaultTemplate forgotPass.TemplateHTML) Builder {
	loader := Builder{}
	err := env.Parse(&loader)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the template build failed", err)
	}

	loader.DefaultTemplate = defaultTemplate

	if loader.TemplatePath != "" {
		loader.setCustomTemplate()
	}

	return loader
}

// Build generates html and text templates using the templateName with the params
func (tb Builder) Build(params interface{}) (string, error) {
	var (
		htmlMessage string
		err         error
	)

	if tb.TemplatePath != "" {
		htmlMessage, err = buildTemplate(tb.CustomName, tb.CustomTemplate, params)
	}

	if err != nil || tb.TemplatePath == "" {
		htmlMessage, err = buildTemplate(tb.DefaultTemplate.Name, tb.DefaultTemplate.HTMLTemplate, params)
	}

	return htmlMessage, err
}

func (tb *Builder) setCustomTemplate() {
	reader, err := ioutil.ReadFile(tb.TemplatePath)
	if err != nil {
		tools.Log().WithError(err).Info("Error parsing custom template provided, default template will be used")
		tb.TemplatePath = ""
		return
	}

	nameSplit := strings.Split(tb.TemplatePath, "/")
	tb.CustomName = nameSplit[len(nameSplit)-1]
	tb.CustomTemplate = string(reader)
}

func buildTemplate(templateName, templateHTML string, templateValues interface{}) (string, error) {
	t := htmlParser.New(templateName)

	var err error
	t, err = t.Parse(templateHTML)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, templateValues); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
