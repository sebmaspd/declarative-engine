# Rules Engine as a common Architectural Component

## Motivation

Most business rules are "encoded" in applications.
Most business rules are hard-coded.
Most business rules are not published as common knowledge-base.
Most business rules could be duplicated in more than one application.
Changes to business rules may require re-development and re-deployment effort.

Need a Rules Engine as part of our Architectural component:  
- a declarative way of defining business rules and policies
- a common knowledge-base of business rules and policies
- updates to business rules and policies can be shared
- reuse of business rules, knowledge-base by applications and by AI
 

## Rules Engine Comparison

| | DMN (Decision Model and Notation) | Custom DMN | JDM (JSON Decision Model | OPA (Open Policy Agent)|
|---|---|---|---|---|
| Standardization | OMG standard, multi-vendor | Custom build based on OMG specs | Open source (GoRules/zen-engine) | [Cloud Native Computing Foundation](https://www.cncf.io/), open source (multi-org maintainers) |
| Format | XML | XML (standard DMN files retained for portability/authoring compatibility with existing DMN tools) | JSON | Rego (own declarative language) |
| Expression language | FEEL | FEEL subset | Simple expression syntax (JS-like) | Rego (Datalog/Prolog-style) |
| Tooling maturity | Mature, enterprise-grade (Camunda, Drools, ODM) | Low — in-house build, no vendor/community tooling or conformance suite behind it; can still leverage existing DMN visual editors (Camunda Modeler, DMN.new) for authoring since the file format is standard | Newer, lighter, growing fast | Mature, widely adopted in cloud-native/K8s ecosystem |
| Embeddability | Heavier runtime | Lightweight — purpose-built to embed natively in Go/Python, avoiding the JVM dependency of standard DMN engines | Very lightweight, easy to embed (Rust core, WASM-capable) | Lightweight, native Go SDK (pure Go, no CGO) |
| Latency | REST APIs round-trips | In-process | In-process | In-process (embedded) or REST/gRPC (server mode) |
| Distributed Point of Failure | Yes | No | No | No (embedded) / Yes (server mode) |
| Additional Infra required | Yes | No | No | No (embedded) / Optional (server mode for centralized policy mgmt) |
| Fit for business common use-case (including Commercial Contracts - e.g. Property rentals, Volume Discounts, Penalties, Company policies, RBAC) | Good fit — decision tables/DRDs suit tiered contract logic (discounts, penalties, approval routing) with business-analyst-facing tooling and audit trail; heavier to stand up (JVM service) for simple cases | Good fit for the same tiered contract/RBAC logic, with the added benefit of native in-process embedding — but only as strong as the FEEL subset actually implemented; full DMN spec (complex boxed expressions, contexts, DRDs) may not be supported, and there's no independent conformance guarantee (e.g., DMN TCK) that it matches standard engine behavior | Strong fit — same tiered/tabular logic (volume discounts, SLA penalties, renewal rules, approval matrices) with lighter footprint; visual JDM editor lets business/legal update rules directly; simple RBAC (role→resource→action, or role+attribute combos) also fits naturally in a decision table | Good fit for RBAC/company-policy access-gating specifically (allow/deny on role, resource, attributes, context) and where policies must integrate with existing K8s/API-gateway/service-mesh authz; usable but not purpose-built for tabular business calculations like discounts/penalties — Rego is developer-authored, less business-analyst-friendly than DMN/JDM tables |
| Effort - High/Medium/Low (Development, Test, Deployment, Maintain) | Development: Medium (visual decision tables ease authoring, but Quarkus/Kogito project scaffolding needed); Test: Medium (Swagger UI + TCK conformance suite available); Deployment: High (JVM build, container image, separate service to run/scale); Maintain: Medium (stable OMG standard, but multi-vendor governance and JVM ops overhead) | Development: High (building/maintaining a custom FEEL-subset interpreter from scratch is significant engineering effort, even reusing standard XML authoring tools); Test: High (no official DMN TCK coverage — must build and maintain your own correctness/conformance test suite); Deployment: Low (embeds directly, no separate service); Maintain: High (full ownership of interpreter bugs, FEEL-subset gaps, and keeping pace with any DMN spec evolution, with no vendor or community support to lean on) | Development: Low (JSON graph, visual JDM editor, native Go/Python bindings — no scaffolding); Test: Low (in-process evaluation, simulator in editor); Deployment: Low (no separate service, ships inside app binary); Maintain: Low-Medium (simple to run, but single-vendor roadmap dependency) | Development: Medium (pure Go SDK is simple to wire in, but Rego's Datalog/Prolog-style syntax has a real learning curve for developers and is not business-analyst-friendly); Test: Low (opa test built-in test framework, Rego Playground); Deployment: Low (embedded) / Medium (server mode needs bundle distribution, policy CI/CD); Maintain: Low-Medium (mature ecosystem, but recent Apple acqui-hire of core maintainers adds a watch item on long-term stewardship) |
| Engine Transparency (Black-box vs White-box) |  |  |  |  |
| Observability (by 3rd Party tools e.g. DataDog, ElasticSearch) |  |  |  |  |
| Configurability (Run-time vs Compile-time) |  |  |  |  |
| Complexity |  |  |  |  |
