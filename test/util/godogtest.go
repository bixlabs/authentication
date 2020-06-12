package util

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestMainWrapper(m *testing.M, featureContext func(*godog.Suite)) {
	var opts = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "pretty",
	}

	flag.Parse()
	opts.Paths = flag.Args()

	// godog v0.9.0 (latest) and earlier
	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		featureContext(s)
	}, opts)

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
