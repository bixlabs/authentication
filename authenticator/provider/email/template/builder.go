package template

import (
	"bytes"
	"runtime"

	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"

	// autoload is not working in main
	htmlTemplate "html/template"
	"path"
	"strings"
	textTemplate "text/template"
	//_ "github.com/joho/godotenv/autoload"
)

// Builder represents a template loader and builder for the emails
type Builder struct {
	TemplatePath string `env:"AUTH_SERVER_EMAIL_TEMPLATE_PATH"`
}

func NewTemplateBuilder() *Builder {
	loader := &Builder{}
	err := env.Parse(loader)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the template build failed", err)
	}

	if loader.TemplatePath == "" {
		_, filename, _, ok := runtime.Caller(1)

		if !ok {
			tools.Log().Panic("Getting the current directory of email templates", err)
		}

		const templateRelativePath = "template"
		loader.TemplatePath = path.Join(path.Dir(filename), templateRelativePath)
	}

	return loader
}

// Build generates html and text templates using the templateName with the params
func (tb *Builder) Build(templateName string, params interface{}) (htmlTemplate, textTemplate string, error error) {
	currentTemplateDirName := strings.Replace(templateName, "_", "", 1)
	currentTemplateDirPath := path.Join(tb.TemplatePath, currentTemplateDirName)

	htmlTemplateName := templateName + ".html"
	htmlTemplatePath := path.Join(currentTemplateDirPath, htmlTemplateName)
	htmlMessage, err := tb.buildHTMLTemplate(htmlTemplateName, htmlTemplatePath, params)
	if err != nil {
		return "", "", err
	}

	textTemplateName := templateName + ".txt"
	textTemplatePath := path.Join(currentTemplateDirPath, textTemplateName)
	textMessage, err := tb.buildTextTemplate(textTemplateName, textTemplatePath, params)
	if err != nil {
		return "", "", err
	}

	return htmlMessage, textMessage, nil
}

func (tb *Builder) buildHTMLTemplate(templateName, templatePath string, templateValues interface{}) (string, error) {
	t := htmlTemplate.New(templateName)

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

func (tb *Builder) buildTextTemplate(templateName, templatePath string, templateValues interface{}) (string, error) {
	t := textTemplate.New(templateName)

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
