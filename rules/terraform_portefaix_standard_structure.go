// Copyright (C) Nicolas Lamirault <nicolas.lamirault@gmail.com>
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

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
	backendVarsDirectory = "backend-vars"
	tfVarsDirectory      = "tfvars"
)

// TerraformPortefaixStandardStructureRule checks whether modules adhere to Terraform Portefaix standard component structure
type TerraformPortefaixStandardStructureRule struct {
	tflint.DefaultRule
}

// NewTerraformPortefaixStandardStructureRule returns a new rule
func NewTerraformPortefaixStandardStructureRule() *TerraformPortefaixStandardStructureRule {
	return &TerraformPortefaixStandardStructureRule{}
}

// Name returns the rule name
func (ruleule *TerraformPortefaixStandardStructureRule) Name() string {
	return "terraform_portefaix_standard_structure"
}

// Enabled returns whether the rule is enabled by default
func (rule *TerraformPortefaixStandardStructureRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (rule *TerraformPortefaixStandardStructureRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (rule *TerraformPortefaixStandardStructureRule) Link() string {
	return "https://github.com/portefaix/tflint-ruleset-portefaix/blob/master/README.md"
}

// Check emits errors for any missing files and any block types that are included in the wrong file
func (rule *TerraformPortefaixStandardStructureRule) Check(runner tflint.Runner) error {
	log.Printf("[DEBUG] Check `%s` rule for runner", rule.Name())
	if err := rule.checkDirectories(runner); err != nil {
		return err
	}

	return nil
}

func (rule *TerraformPortefaixStandardStructureRule) checkDirectories(runner tflint.Runner) error {
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

	backendVarsDir := fmt.Sprintf("%s/%s", directory, backendVarsDirectory)
	if _, err := os.Stat(backendVarsDir); os.IsNotExist(err) {
		message := fmt.Sprintf("Module must include a %s directory as the directory for backend configuration.", backendVarsDir)
		return runner.EmitIssue(rule, message, hcl.Range{Start: hcl.InitialPos})
	}
	tfVarsDir := fmt.Sprintf("%s/%s", directory, tfVarsDirectory)
	if _, err := os.Stat(tfVarsDir); os.IsNotExist(err) {
		message := fmt.Sprintf("Module must include a %s directory as the directory for module configuration.", tfVarsDir)
		return runner.EmitIssue(rule, message, hcl.Range{Start: hcl.InitialPos})
	}

	return nil
}
