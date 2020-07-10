package util

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

// TestMainWrapper binds godog test with go test framework
func TestMainWrapper(m *testing.M, initializeTestSuite func(ctx *godog.TestSuiteContext),
	initializeScenario func(ctx *godog.ScenarioContext)) {
	opts := godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
	}

	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: initializeTestSuite,
		ScenarioInitializer:  initializeScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
