package cmd

import "fmt"

var (
	gitBranch = "unknown"
	gitCommit = "unknown"
	gitState  = "unknown"
	buildDate = "unknown"
)

func Version(appName string) string {
	return fmt.Sprintf(
		"%s(branch=%s commit=%s state=%s date=%s)",
		appName, gitBranch, gitCommit, gitState, buildDate,
	)
}
