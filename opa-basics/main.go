package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

//go:embed policy.rego
var policyFS embed.FS

func main() {
	ctx := context.Background()

	policy, err := policyFS.ReadFile("policy.rego")
	if err != nil {
		log.Fatalf("failed to read policy: %v", err)
	}

	query, err := rego.New(
		rego.Query("data.example.authz.allow"),
		rego.Module("policy.rego", string(policy)),
	).PrepareForEval(ctx)
	if err != nil {
		log.Fatalf("failed to prepare query: %v", err)
	}

	inputs := []map[string]any{
		{"method": "GET", "role": "whocares"},
		{"method": "GET", "role": "guest"},
		{"method": "POST", "role": "guest"},
		{"method": "GET", "role": "admin"},
		{"method": "POST", "role": "admin"},
	}
	fmt.Println("based on policy.rego")
	fmt.Println("everyone allowed to GET")
	fmt.Println("only admin allowed to POST")

	for _, input := range inputs {
		results, err := query.Eval(ctx, rego.EvalInput(input))
		if err != nil {
			log.Fatalf("evaluation failed: %v", err)
		}

		allow := len(results) > 0 && len(results[0].Expressions) > 0 && results[0].Expressions[0].Value == true
		fmt.Printf("input=%v allow=%v\n", input, allow)
	}
}
