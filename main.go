package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"codeium_helm/pkg/helmvalues"
)

func main() {
	// Check if command line arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <values_string> [nested_values_string]")
		fmt.Println("Example: go run main.go \"name=myapp,replicaCount=3\" \"config.data.key1=value1\"")
		os.Exit(1)
	}

	// Get values string from command line argument
	valuesStr := os.Args[1]

	// Parse the values using our helmvalues package
	values, err := helmvalues.ParseValues(valuesStr)
	if err != nil {
		log.Fatalf("Error parsing values: %v", err)
	}

	// Print the parsed values as JSON for better readability
	jsonValues, err := json.MarshalIndent(values, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling values to JSON: %v", err)
	}

	fmt.Println("Parsed values:")
	fmt.Println(string(jsonValues))

	// Check if a second argument is provided for nested values
	if len(os.Args) > 2 {
		nestedValuesStr := os.Args[2]
		
		nestedValues, err := helmvalues.ParseValues(nestedValuesStr)
		if err != nil {
			log.Fatalf("Error parsing nested values: %v", err)
		}

		jsonNestedValues, err := json.MarshalIndent(nestedValues, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling nested values to JSON: %v", err)
		}

		fmt.Println("\nParsed nested values:")
		fmt.Println(string(jsonNestedValues))
	}
}
