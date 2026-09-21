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

| | DMN (Decision Model and Notation) | JDM (JSON Decision Model | OPA (Open Policy Agent)|
|---|---|---|---|
| Standardization | OMG standard, multi-vendor | Open source (GoRules/zen-engine) | [Cloud Native Computing Foundation](https://www.cncf.io/), open source (multi-org maintainers) |
| Format | XML | JSON | Rego (own declarative language) |
| Expression language | FEEL | Simple expression syntax (JS-like) | Rego (Datalog/Prolog-style) |
| Tooling maturity | Mature, enterprise-grade (Camunda, Drools, ODM) | Newer, lighter, growing fast | Mature, widely adopted in cloud-native/K8s ecosystem |
| Embeddability | Heavier runtime | Very lightweight, easy to embed (Rust core, WASM-capable) | Lightweight, native Go SDK (pure Go, no CGO) |
| Latency | REST APIs round-trips | In-process | In-process (embedded) or REST/gRPC (server mode) |
| Distributed Point of Failure | Yes | No | No (embedded) / Yes (server mode) |
| Additional Infra required | Yes | No | No (embedded) / Optional (server mode for centralized policy mgmt) |
| Fit for business common use-case (including Commercial Contracts - e.g. Property rentals, Volume Discounts, Penalties, Company policies, RBAC) | Good fit — decision tables/DRDs suit tiered contract logic (discounts, penalties, approval routing) with business-analyst-facing tooling and audit trail; heavier to stand up (JVM service) for simple cases | Strong fit — same tiered/tabular logic (volume discounts, SLA penalties, renewal rules, approval matrices) with lighter footprint; visual JDM editor lets business/legal update rules directly; simple RBAC (role→resource→action, or role+attribute combos) also fits naturally in a decision table | Good fit for RBAC/company-policy access-gating specifically (allow/deny on role, resource, attributes, context) and where policies must integrate with existing K8s/API-gateway/service-mesh authz; usable but not purpose-built for tabular business calculations like discounts/penalties — Rego is developer-authored, less business-analyst-friendly than DMN/JDM tables |
| Complexity - High/Medium/Low (Development, Test, Deployment, Maintain) | Development: Medium (visual decision tables ease authoring, but Quarkus/Kogito project scaffolding needed); Test: Medium (Swagger UI + TCK conformance suite available); Deployment: High (JVM build, container image, separate service to run/scale); Maintain: Medium (stable OMG standard, but multi-vendor governance and JVM ops overhead) | Development: Low (JSON graph, visual JDM editor, native Go/Python bindings — no scaffolding); Test: Low (in-process evaluation, simulator in editor); Deployment: Low (no separate service, ships inside app binary); Maintain: Low-Medium (simple to run, but single-vendor roadmap dependency) | Development: Medium (pure Go SDK is simple to wire in, but Rego's Datalog/Prolog-style syntax has a real learning curve for developers and is not business-analyst-friendly); Test: Low (opa test built-in test framework, Rego Playground); Deployment: Low (embedded) / Medium (server mode needs bundle distribution, policy CI/CD); Maintain: Low-Medium (mature ecosystem, but recent Apple acqui-hire of core maintainers adds a watch item on long-term stewardship) |
