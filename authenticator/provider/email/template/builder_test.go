package template

import "github.com/cucumber/godog"

func anEmptyEnviromentVariable() error {
	return godog.ErrPending
}

func anCorrectEnvironmentVariable() error {
	return godog.ErrPending
}

func aWrongEnviromentVariable() error {
	return godog.ErrPending
}

func theSystemSendsAnEmail() error {
	return godog.ErrPending
}

func theEmailShouldArriveWithTheDefaultTemplate() error {
	return godog.ErrPending
}

func theEmailShouldArriveWithTheTemplateProvided() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^an empty enviroment variable$`, anEmptyEnviromentVariable)
	s.Step(`^the systems sends an email$`, theSystemSendsAnEmail)
	s.Step(`^the email should arrive with the default template$`, theEmailShouldArriveWithTheDefaultTemplate)

	s.Step(`^an correct environment variable$`, anCorrectEnvironmentVariable)
	s.Step(`^the system sends an email$`, theSystemSendsAnEmail)
	s.Step(`^the email should arrive with the template provided$`, theEmailShouldArriveWithTheTemplateProvided)

	s.Step(`^a wrong enviroment variable$`, aWrongEnviromentVariable)
	s.Step(`^the system sends an email$`, theSystemSendsAnEmail)
	s.Step(`^the email should arrive with the default template$`, theEmailShouldArriveWithTheDefaultTemplate)
}
