package auth

import _ "embed"

// These the current set of rules we have for auth.
const (
	RuleAuthenticate   = "auth"
	RuleAny            = "ruleAny"
	RuleAdminOnly      = "ruleAdminOnly"
	RuleUserOnly       = "ruleUserOnly"
	RuleAdminOrSubject = "ruleAdminOrSubject"
)

// Package name of our rego code.
const (
	opaPackage string = "service.rego"
)

// Core OPA policies.
var (

	//go:embed rego/authorization.rego
	opaAuthorization string
)
