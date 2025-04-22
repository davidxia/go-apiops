package cmd

import (
	"fmt"
	"log"

	"github.com/kong/go-apiops/deckformat"
	"github.com/kong/go-apiops/filebasics"
	"github.com/kong/go-apiops/kong2openapi"
	"github.com/kong/go-apiops/logbasics"
	"github.com/spf13/cobra"
)

func executeKong2OpenAPI(cmd *cobra.Command, _ []string) error {
	verbosity, _ := cmd.Flags().GetInt("verbose")
	logbasics.Initialize(log.LstdFlags, verbosity)

	inputFilename, err := cmd.Flags().GetString("spec")
	if err != nil {
		return fmt.Errorf("failed getting cli argument 'spec'; %w", err)
	}

	options := kong2openapi.K2oOptions{}

	trackInfo := deckformat.HistoryNewEntry("kong2openapi")
	trackInfo["input"] = inputFilename

	// do the work: read/convert/write
	content, err := filebasics.ReadFile(inputFilename)
	if err != nil {
		return err
	}
	result, err := kong2openapi.Convert(content, options)
	if err != nil {
		return fmt.Errorf("failed converting OpenAPI spec '%s'; %w", inputFilename, err)
	}
	deckformat.HistoryAppend(result, trackInfo)
	fmt.Fprintf(cmd.OutOrStdout(), "%s", result)

	return nil
}

var kong2openAPICmd = &cobra.Command{
	Use:   "kong2openapi",
	Short: "Convert Kong's decK format to OpenAPI",
	Long: `Convert Kong's decK format to OpenAPI.

The example file has extensive annotations explaining the conversion
process, as well as all supported custom annotations (x-kong-... directives).
See: https://github.com/Kong/go-apiops/blob/main/docs/learnservice_oas.yaml`,
	RunE: executeKong2OpenAPI,
	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(kong2openAPICmd)
	kong2openAPICmd.Flags().StringP("spec", "s", "-", "Kong decK file to process. Use - to read from stdin")
}
