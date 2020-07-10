package template

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	forgotpass "github.com/bixlabs/authentication/authenticator/provider/email/template/forgotpassword/implementation"
	"github.com/bixlabs/authentication/authenticator/provider/email/template/structures"
	utilTest "github.com/bixlabs/authentication/test/util"
	"github.com/bixlabs/authentication/tools"
	"github.com/cucumber/godog"
)

const relativeGoodPath = "forgotpassword/custom_template_test.html"
const relativeBadPath = "forgotpassword/custom_template_test.htm"

type BuilderTest struct {
	envVariable           string
	testerDefaultTemplate structures.DefaultTemplate
	testerParam           structures.ParamsTemplate
	testerTemplate        string
	tester                Builder
	code                  string
	fullGoodPath          string
	fullBadPath           string
}

func newBuilderTest() *BuilderTest {
	tester := &BuilderTest{code: "99999"}
	_, filename, _, _ := runtime.Caller(0)

	tester.fullGoodPath = path.Join(path.Dir(filename), relativeGoodPath)
	tester.fullBadPath = path.Join(path.Dir(filename), relativeBadPath)
	tester.envVariable = "AUTH_SERVER_EMAIL_TEMPLATE_PATH"
	tester.testerParam = forgotpass.NewTempateParam(tester.code)
	tester.testerDefaultTemplate = forgotpass.NewTemplateHTML()
	tester.tester = NewTemplateBuilder(tester.testerDefaultTemplate)

	return tester
}

func (bd *BuilderTest) anEmptyEnviromentVariable() error {
	err := os.Setenv(bd.envVariable, "")

	if err != nil {
		err = fmt.Errorf("failed seting up an empty environment variable")
	}
	bd.tester = NewTemplateBuilder(bd.testerDefaultTemplate)

	return err
}

func (bd *BuilderTest) aCorrectEnvironmentVariable() error {
	err := os.Setenv(bd.envVariable, bd.fullGoodPath)

	if err != nil {
		err = fmt.Errorf("failed seting up a correct environment variable")
	}
	bd.tester = NewTemplateBuilder(bd.testerDefaultTemplate)

	return err
}

func (bd *BuilderTest) aWrongEnviromentVariable() error {
	err := os.Setenv(bd.envVariable, bd.fullBadPath)

	if err != nil {
		err = fmt.Errorf("failed seting up a wrong environment variable")
	}
	bd.tester = NewTemplateBuilder(bd.testerDefaultTemplate)

	return err
}

func (bd *BuilderTest) theSystemSendsAnEmail() error {
	htmlResponse, err := bd.tester.Build(bd.testerParam)
	bd.testerTemplate = htmlResponse

	if err != nil {
		err = fmt.Errorf("failed building the template")
	}
	return err
}

func (bd *BuilderTest) theEmailShouldArriveWithTheDefaultTemplate() error {
	templateComparator, err := buildTemplate(bd.tester.DefaultName,
		bd.tester.DefaultTemplate, bd.testerParam)

	if err != nil {
		err = fmt.Errorf("failed on building comparator template")
	}

	if strings.Compare(templateComparator, bd.testerTemplate) != 0 {
		err = fmt.Errorf("arrived email not with default template")
	}

	return err
}

func (bd *BuilderTest) theEmailShouldArriveWithTheTemplateProvided() error {
	templateComparator, err := buildTemplate(bd.tester.CustomName,
		bd.tester.CustomTemplate, bd.testerParam)
	if err != nil {
		err = fmt.Errorf("failed on building comparator template")
	}

	if strings.Compare(templateComparator, bd.testerTemplate) != 0 {
		err = fmt.Errorf("arrived email not with custom template")
	}

	return err
}

var builder *BuilderTest

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		tools.InitializeLogger()
		builder = newBuilderTest()
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		builder = newBuilderTest() // clean the state before every scenario
	})

	ctx.Step(`^an empty enviroment variable$`, builder.anEmptyEnviromentVariable)
	ctx.Step(`^the systems sends an email$`, builder.theSystemSendsAnEmail)
	ctx.Step(`^the email should arrive with the default template$`, builder.theEmailShouldArriveWithTheDefaultTemplate)

	ctx.Step(`^an correct environment variable$`, builder.aCorrectEnvironmentVariable)
	ctx.Step(`^the system sends an email$`, builder.theSystemSendsAnEmail)
	ctx.Step(`^the email should arrive with the template provided$`, builder.theEmailShouldArriveWithTheTemplateProvided)

	ctx.Step(`^a wrong enviroment variable$`, builder.aWrongEnviromentVariable)
	ctx.Step(`^the system sends an email$`, builder.theSystemSendsAnEmail)
	ctx.Step(`^the email should arrive with the default template$`, builder.theEmailShouldArriveWithTheDefaultTemplate)

}

func TestMain(m *testing.M) {
	utilTest.TestMainWrapper(m, InitializeTestSuite, InitializeScenario)
}
