package helmvalues

import (
	"fmt"

	"helm.sh/helm/v3/pkg/strvals"
)

// ParseValues parses a string representation of values in the format
// key1=value1,key2=value2 and returns a map of the parsed values.
// It's a wrapper around the strvals.Parse function from Helm v3.
func ParseValues(s string) (map[string]interface{}, error) {
	if s == "" {
		return map[string]interface{}{}, nil
	}

	values := map[string]interface{}{}
	err := strvals.ParseInto(s, values)
	if err != nil {
		return nil, fmt.Errorf("failed to parse values: %w", err)
	}

	return values, nil
}
