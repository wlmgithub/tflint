// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAthenaWorkgroupInvalidNameRule checks the pattern is valid
type AwsAthenaWorkgroupInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsAthenaWorkgroupInvalidNameRule returns new rule with default attributes
func NewAwsAthenaWorkgroupInvalidNameRule() *AwsAthenaWorkgroupInvalidNameRule {
	return &AwsAthenaWorkgroupInvalidNameRule{
		resourceType:  "aws_athena_workgroup",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-z0-9._-]{1,128}$`),
	}
}

// Name returns the rule name
func (r *AwsAthenaWorkgroupInvalidNameRule) Name() string {
	return "aws_athena_workgroup_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAthenaWorkgroupInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAthenaWorkgroupInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAthenaWorkgroupInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAthenaWorkgroupInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[a-zA-z0-9._-]{1,128}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
