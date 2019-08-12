package template

import (
	"bytes"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	// autoload is not working in main
	_ "github.com/joho/godotenv/autoload"
	htmlTemplate "html/template"
	"os"
	"path"
	"strings"
	textTemplate "text/template"
)

type Builder struct {
	TemplatePath string `env:"EMAIL_TEMPLATE_PATH" envDefault:"authenticator/provider/email/template/"`
}

func NewTemplateBuilder() *Builder {
	const templateRelativePath = "authenticator/provider/email/template/"

	loader := &Builder{}
	err := env.Parse(loader)

	if err != nil {
		tools.Log().Panic("Parsing the env variables for the template build failed", err)
	}

	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if loader.TemplatePath == "" {
		loader.TemplatePath = path.Join(rootDir, templateRelativePath)
	}

	return loader
}

func (tb Builder) Build(templateName string, templateParams interface{}) (htmlTemplate, textTemplate string, error error) {
	currentTemplateDirName := strings.Replace(templateName, "_", "", 1)
	currentTemplateDirPath := path.Join(tb.TemplatePath, currentTemplateDirName)

	htmlTemplateName := templateName + ".html"
	htmlTemplatePath := path.Join(currentTemplateDirPath, htmlTemplateName)
	htmlMessage, err := tb.buildHTMLTemplate(htmlTemplateName, htmlTemplatePath, templateParams)
	if err != nil {
		return "", "", err
	}

	textTemplateName := templateName + ".txt"
	textTemplatePath := path.Join(currentTemplateDirPath, textTemplateName)
	textMessage, err := tb.buildTextTemplate(textTemplateName, textTemplatePath, templateParams)
	if err != nil {
		return "", "", err
	}

	return htmlMessage, textMessage, nil
}

func (tb Builder) buildHTMLTemplate(templateName, templatePath string, templateValues interface{}) (string, error) {
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

func (tb Builder) buildTextTemplate(templateName, templatePath string, templateValues interface{}) (string, error) {
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