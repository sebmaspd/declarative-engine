package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/open-policy-agent/opa/rego"
)

//go:embed policy.rego
var policyFS embed.FS

type result struct {
	Tier                  string  `json:"tier"`
	UnitPrice             float64 `json:"unit_price"`
	Discount              float64 `json:"discount"`
	NetUnitPrice          float64 `json:"net_unit_price"`
	VolumeDiscountedPrice float64 `json:"volume_discounted_price"`
}

func main() {
	purchaseVolume := flag.Int("purchase-volume", 0, "cumulative purchase volume, in units, for the contract cycle (required)")
	flag.Parse()

	if *purchaseVolume <= 0 {
		fmt.Fprintln(os.Stderr, "error: -purchase-volume must be a positive integer")
		flag.Usage()
		os.Exit(1)
	}

	ctx := context.Background()

	policy, err := policyFS.ReadFile("policy.rego")
	if err != nil {
		log.Fatalf("failed to read policy: %v", err)
	}

	query, err := rego.New(
		rego.Query("data.volumediscounts.result"),
		rego.Module("policy.rego", string(policy)),
	).PrepareForEval(ctx)
	if err != nil {
		log.Fatalf("failed to prepare query: %v", err)
	}

	results, err := query.Eval(ctx, rego.EvalInput(map[string]any{
		"purchase_volume": *purchaseVolume,
	}))
	if err != nil {
		log.Fatalf("evaluation failed: %v", err)
	}
	if len(results) == 0 || len(results[0].Expressions) == 0 {
		log.Fatalf("policy produced no result")
	}

	raw, err := json.Marshal(results[0].Expressions[0].Value)
	if err != nil {
		log.Fatalf("failed to marshal result: %v", err)
	}

	var r result
	if err := json.Unmarshal(raw, &r); err != nil {
		log.Fatalf("failed to unmarshal result: %v", err)
	}

	fmt.Printf("Purchase Volume:          %d units (%s Tier)\n", *purchaseVolume, r.Tier)
	fmt.Printf("Unit-Price:               $%.2f\n", r.UnitPrice)
	fmt.Printf("Discount:                 %.0f%%\n", r.Discount*100)
	fmt.Printf("Net-Unit-Price:           $%.2f\n", r.NetUnitPrice)
	fmt.Printf("Volume-Discounted-Price:  $%.2f\n", r.VolumeDiscountedPrice)
}
