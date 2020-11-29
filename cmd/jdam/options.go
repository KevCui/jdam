package main

import (
	"errors"
	"flag"
)

type options struct {
	Mutators  *string
	Ignore    *string
	Output    *string
	Count     *int
	Rounds    *int
	NilChance *float64
	MaxDepth  *int
	Seed      *int64
	List      *bool
	Verbose   *bool
	Version   *bool
}

func parseOptions() (options, error) {
	options := options{
		Mutators:  flag.String("mutators", "", "Comma-separated list of mutator IDs to use (default: all)"),
		Ignore:    flag.String("ignore", "", "Comma-separated list of fields to exclude from fuzzing"),
		Output:    flag.String("output", "", "Output file pattern to use for results (e.g. /tmp/jdam-%d.json)"),
		Rounds:    flag.Int("rounds", 1, "Number of times to fuzz object"),
		Count:     flag.Int("count", 1, "Number of fuzzed objects to generate"),
		NilChance: flag.Float64("nil-chance", .75, "Probability of value being set to nil (between 0 (no nils) and 1 (all nils))"),
		MaxDepth:  flag.Int("max-depth", 100, "Maximum object depth to fuzz"),
		Seed:      flag.Int64("seed", 0, "Seed to use for pseudo-random number generator (default: current UNIX timestamp)"),
		List:      flag.Bool("list", false, "List available mutators"),
		Verbose:   flag.Bool("verbose", false, "Print activity information"),
		Version:   flag.Bool("version", false, "Print current jdam version"),
	}

	flag.Parse()

	if *options.NilChance < 0 || *options.NilChance > 1 {
		return options, errors.New("nil chance must be between 0 and 1")
	}

	return options, nil
}
