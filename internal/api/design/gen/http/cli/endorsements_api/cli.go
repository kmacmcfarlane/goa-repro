// Code generated by goa v3.5.3, DO NOT EDIT.
//
// endorsements-api HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/kmacmcfarlane/goa-reference-bug-repro/internal/api/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	fooc "github.com/kmacmcfarlane/goa-reference-bug-repro/internal/api/design/gen/http/foo/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `foo get
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` foo get --body '{
      "fromProfileID": 658594,
      "fromProfileUUID": "91533209-d516-41a8-aa93-97e36b4ab195",
      "validateMe": "one"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		fooFlags = flag.NewFlagSet("foo", flag.ContinueOnError)

		fooGetFlags    = flag.NewFlagSet("get", flag.ExitOnError)
		fooGetBodyFlag = fooGetFlags.String("body", "REQUIRED", "")
	)
	fooFlags.Usage = fooUsage
	fooGetFlags.Usage = fooGetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "foo":
			svcf = fooFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "foo":
			switch epn {
			case "get":
				epf = fooGetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "foo":
			c := fooc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get":
				endpoint = c.Get()
				data, err = fooc.BuildGetPayload(*fooGetBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// fooUsage displays the usage of the foo command and its subcommands.
func fooUsage() {
	fmt.Fprintf(os.Stderr, `Service is the foo service interface.
Usage:
    %[1]s [globalflags] foo COMMAND [flags]

COMMAND:
    get: Get implements Get.

Additional help:
    %[1]s foo COMMAND --help
`, os.Args[0])
}
func fooGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] foo get -body JSON

Get implements Get.
    -body JSON: 

Example:
    %[1]s foo get --body '{
      "fromProfileID": 658594,
      "fromProfileUUID": "91533209-d516-41a8-aa93-97e36b4ab195",
      "validateMe": "one"
   }'
`, os.Args[0])
}
