# declarative-engine
Quick evaluations for rules engine the declarative way using OPA (Open Policy Agent).  

## Prompts

1. In directory `opa/`, create a minimal OPA (Open Policy Agent) example in Golang.
2. Rename `opa` directory to `opa-basics`.
3. Based on the sample contract clause in `volume-discounts/contract-clause.md`, create an OPA policy and a Go program that accepts the Purchase Volume as input. The output will show the Unit-Price, Discount, Net-Unit-Price, Volume-Discounted-Price.
4. Update the main program to properly handle input arguments using the `flag` package.
5. Write a Makefile to build the respective binaries; the OPA policy files should be named after the binaries, and saved to the `/bin` directory.
6. List all previous prompts in the README.md "Prompts" section.

## OPA Policy Playground

**Rego** (pronounced "ray-go") is OPA's policy language.  

https://play.openpolicyagent.org/  

Paste the .rego (without `package` name) into the playground, enter the INPUT { "<input.var_name>" : <value> }, then click `Evaluate` to see the OUTPUT.