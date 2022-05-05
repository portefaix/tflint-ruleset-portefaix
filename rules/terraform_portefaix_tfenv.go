// Copyright (C) 2021 Nicolas Lamirault <nicolas.lamirault@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rules

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const (
	filenameTfenv = ".terraform-version"
)

// TerraformPortefaixTfenvRule checks whether modules adhere to Terraform Portefaix standard component structure
type TerraformPortefaixTfenvRule struct {
	tflint.DefaultRule
}

// NewTerraformPortefaixTfenvRule returns a new rule
func NewTerraformPortefaixTfenvRule() *TerraformPortefaixTfenvRule {
	return &TerraformPortefaixTfenvRule{}
}

// Name returns the rule name
func (ruleule *TerraformPortefaixTfenvRule) Name() string {
	return "terraform_portefaix_tfenv"
}

// Enabled returns whether the rule is enabled by default
func (rule *TerraformPortefaixTfenvRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (rule *TerraformPortefaixTfenvRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (rule *TerraformPortefaixTfenvRule) Link() string {
	return "https://github.com/portefaix/tflint-ruleset-portefaix/blob/master/README.md"
}

// Check emits errors for any missing files and any block types that are included in the wrong file
func (rule *TerraformPortefaixTfenvRule) Check(runner tflint.Runner) error {
	log.Printf("[DEBUG] Check `%s` rule for runner", rule.Name())
	if err := rule.checkTerraformVersionFile(runner); err != nil {
		return err
	}

	return nil
}

func (rule *TerraformPortefaixTfenvRule) checkTerraformVersionFile(runner tflint.Runner) error {
	files, _ := runner.GetFiles()
	log.Printf("[DEBUG] Files: %d", len(files))

	if len(files) == 0 {
		return nil
	}

	keys := make([]string, 0, len(files))
	for k := range files {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	directory := path.Dir(keys[0])

	tfEnvFilename := fmt.Sprintf("%s/%s", directory, filenameTfenv)
	if _, err := os.Stat(tfEnvFilename); os.IsNotExist(err) {
		message := fmt.Sprintf("Module must include a %s filename for Tfenv to specify which Terraform version use.", filenameTfenv)
		return runner.EmitIssue(rule, message, hcl.Range{Start: hcl.InitialPos})
	}

	return nil
}
