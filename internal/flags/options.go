package flags

import (
	"context"
	"io"
	"io/fs"
	"testing"
)

// Options are suite run options
// flags are mapped to these options.
//
// It can also be used together with godog.RunWithOptions
// to run test suite from go source directly
//
// See the flags for more details
type Options struct {
	// Print step definitions found and exit
	ShowStepDefinitions bool

	// Randomize, if not `0`, will be used to run scenarios in a random order.
	//
	// Randomizing scenario order is especially helpful for detecting
	// situations where you have state leaking between scenarios, which can
	// cause flickering or fragile tests.
	//
	// The default value of `0` means "do not randomize".
	//
	// The magic value of `-1` means "pick a random seed for me", and godog will
	// assign a seed on it's own during the `RunWithOptions` phase, similar to if
	// you specified `--random` on the command line.
	//
	// Any other value will be used as the random seed for shuffling. Re-using the
	// same seed will allow you to reproduce the shuffle order of a previous run
	// to isolate an error condition.
	Randomize int64

	// Stops on the first failure
	StopOnFailure bool

	// Fail suite when there are pending or undefined or ambiguous steps
	Strict bool

	// Forces ansi color stripping
	NoColors bool

	// Various filters for scenarios parsed
	// from feature files
	Tags string

	// Dialect to be used to parse feature files. If not set, default to "en".
	Dialect string

	// The formatter name
	Format string

	// Concurrency rate, not all formatters accepts this
	Concurrency int

	// All feature file paths
	Paths []string

	// Where it should print formatter output
	Output io.Writer

	// DefaultContext is used as initial context instead of context.Background().
	DefaultContext context.Context

	// TestingT runs scenarios as subtests.
	TestingT *testing.T

	// FeatureContents allows passing in each feature manually
	// where the contents of each feature is stored as a byte slice
	// in a map entry
	FeatureContents []Feature

	// FS allows passing in an `fs.FS` to read features from, such as an `embed.FS`
	// or os.DirFS(string).
	FS fs.FS

	// ShowHelp enables suite to show CLI flags usage help and exit.
	ShowHelp bool

	// MaxRetries is the number of times you can retry failing tests. Default is 0 retry.
	MaxRetries int
}

type Feature struct {
	Name     string
	Contents []byte
}
