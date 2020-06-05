package template

import (
	"bytes"
	htmlParser "html/template"
	"strings"

	forgotPass "github.com/bixlabs/authentication/authenticator/provider/email/template/forgotpassword"

	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	// autoload is not working in main
	//_ "github.com/joho/godotenv/autoload"
)

// Builder represents a template loader and builder for the emails
type Builder struct {
	TemplatePath string `env:"AUTH_SERVER_EMAIL_TEMPLATE_PATH"`
	custom       bool
}

func NewTemplateBuilder() *Builder {
	loader := &Builder{custom: true}
	err := env.Parse(loader)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the template build failed", err)
	}

	if loader.TemplatePath == "" {
		loader.custom = false
	}

	return loader
}

// Build generates html and text templates using the templateName with the params
func (tb *Builder) Build(defaultHTML forgotPass.TemplateHTML, params interface{}) (string, error) {
	var (
		htmlMessage string
		err         error
	)

	if !tb.custom {
		htmlMessage, err = tb.defaultTemplateBuild(defaultHTML, params)
	} else {
		htmlMessage, err = tb.customTemplateBuild(params)

		if err != nil {
			htmlMessage, err = tb.defaultTemplateBuild(defaultHTML, params)
			tools.Log().Info("The custom template provided is not valid, default template used instead")
		}
	}

	return htmlMessage, err
}

func (tb *Builder) defaultTemplateBuild(defaultHTML forgotPass.TemplateHTML, params interface{}) (string, error) {
	htmlMessage, err := tb.buildDefaultTemplate(defaultHTML, params)

	if err != nil {
		return "", err
	}

	return htmlMessage, nil
}

func (tb *Builder) customTemplateBuild(params interface{}) (string, error) {
	pathSplit := strings.Split(tb.TemplatePath, "/")
	customHTMLTemplateName := pathSplit[len(pathSplit)-1]

	htmlMessage, err := tb.buildHTMLTemplate(customHTMLTemplateName, tb.TemplatePath, params)

	if err != nil {
		return "", err
	}

	return htmlMessage, nil
}

func (tb *Builder) buildHTMLTemplate(templateName, templatePath string, templateValues interface{}) (string, error) {
	t := htmlParser.New(templateName)

	var err error
	t, err = t.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, templateValues); err != nil {
		return "", err
	}

	return tpl.String(), nil
}

func (tb *Builder) buildDefaultTemplate(template forgotPass.TemplateHTML, templateValues interface{}) (string, error) {
	t := htmlParser.New(template.Name)

	var err error
	t, err = t.Parse(template.HTMLTemplate)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, templateValues); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
