package template

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/bixlabs/authentication/authenticator/provider/email/template/forgotpassword"
	"github.com/bixlabs/authentication/tools"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

const relativeGoodPath = "forgotpassword/custom_template_test.html"
const relativeBadPath = "forgotpassword/custom_template_test.htm"

type BuilderTest struct {
	envVariable           string
	testerDefaultTemplate forgotpassword.TemplateHTML
	testerParam           *forgotpassword.TemplateParam
	testerTemplate        string
	tester                *Builder
	code                  string
	fullGoodPath          string
	fullBadPath           string
}

func newBuilderTest() *BuilderTest {
	tester := &BuilderTest{code: "99999"}
	_, filename, _, _ := runtime.Caller(0)

	tester.fullGoodPath = path.Join(path.Dir(filename), relativeGoodPath)
	tester.fullBadPath = path.Join(path.Dir(filename), relativeBadPath)
	tester.tester = NewTemplateBuilder()
	tester.envVariable = "AUTH_SERVER_EMAIL_TEMPLATE_PATH"
	tester.testerParam = forgotpassword.NewTempateParam(tester.code)
	tester.testerDefaultTemplate = forgotpassword.NewTemplateHTML()

	return tester
}

func (bd *BuilderTest) anEmptyEnviromentVariable() error {
	err := os.Setenv(bd.envVariable, "")

	if err != nil {
		err = fmt.Errorf("failed seting up an empty environment variable")
	}
	bd.tester = NewTemplateBuilder()

	return err
}

func (bd *BuilderTest) aCorrectEnvironmentVariable() error {
	err := os.Setenv("AUTH_SERVER_EMAIL_TEMPLATE_PATH", bd.fullGoodPath)

	if err != nil {
		err = fmt.Errorf("failed seting up a correct environment variable")
	}
	bd.tester = NewTemplateBuilder()

	return err
}

func (bd *BuilderTest) aWrongEnviromentVariable() error {
	err := os.Setenv(bd.envVariable, bd.fullBadPath)

	if err != nil {
		err = fmt.Errorf("failed seting up a wrong environment variable")
	}
	bd.tester = NewTemplateBuilder()

	return err
}

func (bd *BuilderTest) theSystemSendsAnEmail() error {
	htmlResponse, err := bd.tester.Build(bd.testerDefaultTemplate, bd.testerParam)
	bd.testerTemplate = htmlResponse

	if err != nil {
		err = fmt.Errorf("failed building the template")
	}
	return err
}

func (bd *BuilderTest) theEmailShouldArriveWithTheDefaultTemplate() error {
	templateComparator, err := bd.tester.defaultTemplateBuild(bd.testerDefaultTemplate, bd.testerParam)

	if err != nil {
		err = fmt.Errorf("failed on building comparator template")
	}

	if strings.Compare(templateComparator, bd.testerTemplate) != 0 {
		err = fmt.Errorf("arrived email not with default template")
	}

	return err
}

func (bd *BuilderTest) theEmailShouldArriveWithTheTemplateProvided() error {
	templateComparator, err := bd.tester.customTemplateBuild(bd.testerParam)

	if err != nil {
		err = fmt.Errorf("failed on building comparator template")
	}

	if strings.Compare(templateComparator, bd.testerTemplate) != 0 {
		err = fmt.Errorf("arrived email not with custom template")
	}

	return err
}

func FeatureContext(s *godog.Suite) {
	tools.InitializeLogger()
	builder := newBuilderTest()

	s.Step(`^an empty enviroment variable$`, builder.anEmptyEnviromentVariable)
	s.Step(`^the systems sends an email$`, builder.theSystemSendsAnEmail)
	s.Step(`^the email should arrive with the default template$`, builder.theEmailShouldArriveWithTheDefaultTemplate)

	s.Step(`^an correct environment variable$`, builder.aCorrectEnvironmentVariable)
	s.Step(`^the system sends an email$`, builder.theSystemSendsAnEmail)
	s.Step(`^the email should arrive with the template provided$`, builder.theEmailShouldArriveWithTheTemplateProvided)

	s.Step(`^a wrong enviroment variable$`, builder.aWrongEnviromentVariable)
	s.Step(`^the system sends an email$`, builder.theSystemSendsAnEmail)
	s.Step(`^the email should arrive with the default template$`, builder.theEmailShouldArriveWithTheDefaultTemplate)
}

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	// godog v0.9.0 (latest) and earlier
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opts)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}