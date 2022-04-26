package cmd

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
)

type Flags struct {
	History           int
	Fieldnames        string
	ErrorMappingField string
	DisplayVersion    bool
}

func Run(args []string, stdin io.Reader, stdout io.Writer) error {
	var cfg Flags

	fset := flag.NewFlagSet("tailog", flag.ExitOnError)
	fset.Usage = func() {
		var flagsHelp strings.Builder
		output := fset.Output()
		fset.SetOutput(&flagsHelp)
		fset.PrintDefaults()
		fset.SetOutput(output)

		fmt.Fprintf(fset.Output(), helmMessage, fset.Name(), flagsHelp.String())
	}
	fset.BoolVar(&cfg.DisplayVersion, "v", false, "display CLI version")
	fset.IntVar(&cfg.History, "history", 100, "numbers of line of logs to keep")
	fset.StringVar(&cfg.Fieldnames, "fields", "http.method,http.target,http.status_code,msg", "fields to display from the log stream")
	fset.StringVar(&cfg.ErrorMappingField, "errorfield", "msg", "field to use to display errors coming from parsing the log stream")

	if err := fset.Parse(args); err != nil {
		return err
	}

	if cfg.DisplayVersion {
		fmt.Fprintln(stdout, Version(fset.Name()))
		return nil
	}

	fieldNames := strings.Split(cfg.Fieldnames, ",")
	if len(fieldNames) == 0 {
		return errors.New("fields list must contain at least one field")
	}

	return StartTUI(fset.Name(), stdin, cfg.History, cfg.ErrorMappingField, fieldNames)
}

var helmMessage = `usage: %[1]s [flags]

Tailog is in charge of reading structued logs from STDIN and make them
accessible for display and filtering. The default structured log format is JSON

Flags

%[2]s

Examples

$> myapp | %[1]s -history 50
start myapp binary and display the 50 last JSON strucuted logs
`
