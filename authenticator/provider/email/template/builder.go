package template

import (
	"bytes"
	"errors"
	htmlParser "html/template"
	"io/ioutil"
	"strings"

	"github.com/bixlabs/authentication/authenticator/provider/email/template/structures"

	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
)

// Builder represents a template loader and builder for the emails
type Builder struct {
	TemplatePath    string `env:"AUTH_SERVER_EMAIL_TEMPLATE_PATH"`
	DefaultTemplate string
	DefaultName     string
	CustomTemplate  string
	CustomName      string
}

func NewTemplateBuilder(defaultTemplate structures.DefaultTemplate) Builder {
	loader := Builder{}
	err := env.Parse(&loader)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the template build failed", err)
	}

	if loader.TemplatePath != "" {
		loader.setCustomTemplate()
	}

	loader.DefaultTemplate, loader.DefaultName = defaultTemplate.GetTemplate()

	return loader
}

// Build generates html and text templates using the templateName with the params
func (tb Builder) Build(params structures.ParamsTemplate) (string, error) {
	var (
		htmlMessage string
		err         error
	)

	if tb.TemplatePath != "" {
		htmlMessage, err = buildTemplate(tb.CustomName, tb.CustomTemplate, params)
	}

	if err != nil || tb.TemplatePath == "" {
		htmlMessage, err = buildTemplate(tb.DefaultName, tb.DefaultTemplate, params)
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

func buildTemplate(templateName, templateHTML string, templateValues structures.ParamsTemplate) (string, error) {
	t := htmlParser.New(templateName)

	var err error
	t, err = t.Parse(templateHTML)
	if err != nil {
		return "", err
	}

	for _, v := range templateValues.GetParams() {
		if !strings.Contains(templateHTML, v) {
			err = errors.New("template does not contains one or more parameters")
			tools.Log().WithError(err).Warn("Custom html is not completely correct, default email used instead")
			return "", err
		}
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, templateValues); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
