package cmd_test

import (
	"strings"
	"testing"

	"github.com/lonepeon/golib/testutils"
	"github.com/lonepeon/tailog/internal/cmd"
)

func TestRunVersion(t *testing.T) {
	var stdout strings.Builder

	err := cmd.Run([]string{"-v"}, nil, &stdout)
	testutils.RequireNoError(t, err, "unexpected error while running the command")

	testutils.AssertContainsString(t, "branch=", stdout.String(), "missing branch information")
	testutils.AssertContainsString(t, "state=", stdout.String(), "missing branch information")
	testutils.AssertContainsString(t, "date=", stdout.String(), "missing branch information")
}
