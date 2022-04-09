package cmd

import "fmt"

var (
	gitBranch = "unknown"
	gitCommit = "unknown"
	gitState  = "unknown"
	buildDate = "unknown"
)

func Version() string {
	return fmt.Sprintf(
		"tailog(branch=%s commit=%s state=%s date=%s)",
		gitBranch, gitCommit, gitState, buildDate,
	)
}
