package kong2openapi

import (
	"github.com/kong/go-apiops/logbasics"
)

type K2oOptions struct {
}

func (opts *K2oOptions) setDefaults() {
}

// Convert converts a Kong declarative file to an OpenAPI spec.
func Convert(content []byte, opts K2oOptions) (map[string]interface{}, error) {
	opts.setDefaults()
	logbasics.Debug("received OpenAPI2Kong options", "options", opts)

	// set up output document
	result := make(map[string]interface{})

	// we're done!
	logbasics.Debug("finished processing document")
	return result, nil
}
