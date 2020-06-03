package template

import (
	"bytes"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"runtime"

	// autoload is not working in main
	_ "github.com/joho/godotenv/autoload"
	htmlTemplate "html/template"
	"path"
	"strings"
	textTemplate "text/template"
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
		loader.TemplatePath = getDefaultPath()
		loader.custom = false
	}

	tools.Log().Printf("custom path set for template: %v\n", loader.custom)
	return loader
}

// Build generates html and text templates using the templateName with the params
func (tb *Builder) Build(defaultTemplateName string, params interface{}) (string, string, error) {
	var (
		htmlMessage string
		textMessage string
		err         error
	)

	if !tb.custom {
		htmlMessage, textMessage, err = tb.defaultTemplateBuild(defaultTemplateName, params)
	} else {
		htmlMessage, textMessage, err = tb.customTemplateBuild(params)

		if err != nil {
			tb.TemplatePath = getDefaultPath()
			htmlMessage, textMessage, err = tb.defaultTemplateBuild(defaultTemplateName, params)
		}
	}

	return htmlMessage, textMessage, err
}

func getDefaultPath() string {
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		tools.Log().Panic("Getting the current directory of email templates")
	}

	defaultPath := path.Join(path.Dir(filename))
	return defaultPath
}

func (tb *Builder) defaultTemplateBuild(defaultTemplateName string, params interface{}) (string, string, error) {
	currentDefaultTemplateDirName := strings.Replace(defaultTemplateName, "_", "", 1)
	currentDefaultTemplateDirPath := path.Join(tb.TemplatePath, currentDefaultTemplateDirName)

	defaultHTMLTemplateName := defaultTemplateName + ".html"
	defaultHTMLTemplatePath := path.Join(currentDefaultTemplateDirPath, defaultHTMLTemplateName)

	htmlMessage, err := tb.buildHTMLTemplate(defaultHTMLTemplateName, defaultHTMLTemplatePath, params)

	if err != nil {
		return "", "", err
	}

	defaultTextTemplateName := defaultTemplateName + ".txt"
	defaultTextTemplatePath := path.Join(currentDefaultTemplateDirPath, defaultTextTemplateName)

	textMessage, err := tb.buildTextTemplate(defaultTextTemplateName, defaultTextTemplatePath, params)

	if err != nil {
		return "", "", err
	}

	return htmlMessage, textMessage, nil
}

func (tb *Builder) customTemplateBuild(params interface{}) (string, string, error) {
	pathSplit := strings.Split(tb.TemplatePath, "/")
	customHTMLTemplateName := pathSplit[len(pathSplit)-1]

	htmlMessage, err := tb.buildHTMLTemplate(customHTMLTemplateName, tb.TemplatePath, params)

	if err != nil {
		return "", "", err
	}

	customTextTemplateName := pathSplit[len(pathSplit)-1]
	textMessage, err := tb.buildTextTemplate(customTextTemplateName, tb.TemplatePath, params)

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
