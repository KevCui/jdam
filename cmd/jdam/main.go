package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitlab.com/michenriksen/jdam/pkg/jdam"
	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"

	"github.com/olekukonko/tablewriter"
)

func main() {
	logger := logger{}
	options, err := parseOptions()
	if err != nil {
		logger.Fatal(err.Error() + "\n")
	}
	logger.SetVerbose(*options.Verbose)

	if *options.Version {
		fmt.Printf("jdam v%s\n", version)
		os.Exit(0)
	}

	if *options.List {
		printMutatorList()
		os.Exit(0)
	}

	logger.Important("jdam v%s started at %s\n", version, time.Now().Format(time.RFC3339))

	if *options.Output != "" {
		outDir := filepath.Dir(*options.Output)
		stat, err := os.Stat(outDir)
		if err != nil {
			if os.IsNotExist(err) {
				logger.Fatal("Invalid output pattern: output directory does not exist\n")
			} else {
				logger.Fatal("Invalid output path: %s\n", err.Error())
			}
		}
		if !stat.IsDir() {
			logger.Fatal("Invalid output pattern: output directory is not a directory")
		}
	}

	var mutators mutation.MutatorList
	if *options.Mutators != "" {
		ids := strings.Split(*options.Mutators, ",")
		mutators = mutatorListFromIDs(ids)
	} else {
		mutators = mutation.Mutators
	}

	logger.Info("Fuzzing with %d mutators\n", len(mutators))

	reader := bufio.NewReader(os.Stdin)
	var input []rune
	for {
		char, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		input = append(input, char)
	}

	subject := map[string]interface{}{}
	err = json.Unmarshal([]byte(strings.TrimSpace(string(input))), &subject)
	if err != nil {
		logger.Fatal("Unable to parse input as JSON object: %s\n", err.Error())
	}

	fuzzed := map[string]interface{}{}

	var fuzzer *jdam.Fuzzer
	if *options.Seed != 0 {
		logger.Info("Using PRNG seed: %d\n", *options.Seed)
		fuzzer = jdam.NewWithSeed(*options.Seed, mutators)
	} else {
		fuzzer = jdam.New(mutators)
	}
	fuzzer.IgnoreFields(strings.Split(*options.Ignore, ",")).
		NilChance(*options.NilChance).
		MaxDepth(*options.MaxDepth)

	for i := 0; i < *options.Count; i++ {
		fuzzed = subject
		for j := 0; j < *options.Rounds; j++ {
			logger.Info("[OBJECT #%d] Fuzzing round: %d\n", i+1, j+1)
			fuzzed = fuzzer.Fuzz(fuzzed)
		}

		fuzzedJSON, err := json.Marshal(fuzzed)
		if err != nil {
			logger.Fatal("[OBJECT #%d] Unable to marshal fuzzed object as JSON: %s\n", i+1, err.Error())
		}

		if *options.Output != "" {
			outFile := fmt.Sprintf(*options.Output, i+1)
			err := ioutil.WriteFile(outFile, fuzzedJSON, 0644)
			if err != nil {
				logger.Fatal("[OBJECT #%d] Unable to write fuzzed json object to %s: %s\n", i+1, outFile, err.Error())
			}
			logger.Info("[OBJECT #%d] Wrote fuzzed object to %s\n", i+1, outFile)
		} else {
			fmt.Println(string(fuzzedJSON))
		}
	}
}

func printMutatorList() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", ""})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	for _, m := range mutation.Mutators {
		table.Append([]string{m.ID(), m.Name(), m.Description()})
	}

	table.Render()
}

func mutatorListFromIDs(ids []string) mutation.MutatorList {
	var mutators mutation.MutatorList
	for _, m := range mutation.Mutators {
		for _, id := range ids {
			if id == m.ID() {
				mutators = append(mutators, m)
				break
			}
		}
	}
	return mutators
}
