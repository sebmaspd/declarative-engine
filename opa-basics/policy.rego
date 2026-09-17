package example.authz

import rego.v1

default allow := false

allow if {
	input.method == "GET"
}

allow if {
	input.method == "POST"
	input.role == "admin"
}
