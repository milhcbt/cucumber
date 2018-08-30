/*
This is a console application that prints Cucumber messages to
STDOUT. By default it prints them as protobuf, but the --json flag
will print them as JSON (useful for testing/debugging)
*/
package main

import (
	b64 "encoding/base64"
	"flag"
	"fmt"
	"github.com/cucumber/gherkin-go"
	"os"
)

var noSource = flag.Bool("no-source", false, "Skip GherkinSource messages")
var noAst = flag.Bool("no-ast", false, "Skip GherkinDocument messages")
var noPickles = flag.Bool("no-pickles", false, "Skip Pickle messages")
var printJson = flag.Bool("json", false, "Print messages as JSON instead of protobuf")
var versionFlag = flag.Bool("version", false, "print version")
var dialectsFlag = flag.Bool("dialects", false, "print dialects as JSON")
var fakeResultsFlag = flag.Bool("fake-results", false, "Emit fake results messages")
var defaultDialectFlag = flag.String("default-dialect", "en", "the default dialect")

// Set during build with -ldflags
var version string
var gherkinDialects string

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Printf("gherkin %s\n", version)
		os.Exit(0)
	}

	if *dialectsFlag {
		sDec, _ := b64.StdEncoding.DecodeString(gherkinDialects)
		fmt.Println(string(sDec))
		os.Exit(0)
	}

	paths := flag.Args()

	_, err := gherkin.Messages(
		paths,
		os.Stdin,
		*defaultDialectFlag,
		!*noSource,
		!*noAst,
		!*noPickles,
		os.Stdout,
		*printJson,
		*fakeResultsFlag,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse Gherkin: %+v\n", err)
		os.Exit(1)
	}
}
