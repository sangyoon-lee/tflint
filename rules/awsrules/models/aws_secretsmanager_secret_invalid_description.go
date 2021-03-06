// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSecretsmanagerSecretInvalidDescriptionRule checks the pattern is valid
type AwsSecretsmanagerSecretInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsSecretsmanagerSecretInvalidDescriptionRule returns new rule with default attributes
func NewAwsSecretsmanagerSecretInvalidDescriptionRule() *AwsSecretsmanagerSecretInvalidDescriptionRule {
	return &AwsSecretsmanagerSecretInvalidDescriptionRule{
		resourceType:  "aws_secretsmanager_secret",
		attributeName: "description",
		max:           2048,
	}
}

// Name returns the rule name
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Name() string {
	return "aws_secretsmanager_secret_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSecretsmanagerSecretInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
