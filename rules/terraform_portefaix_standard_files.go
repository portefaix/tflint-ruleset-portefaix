// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rules

import (
	"fmt"
	"log"
	"path"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const (
	filenameProvider = "provider.tf"
)

// TerraformPortefaixStandardFilesRule checks whether modules adhere to Terraform Portefaix standard component structure
type TerraformPortefaixStandardFilesRule struct{}

// NewTerraformPortefaixStandardFilesRule returns a new rule
func NewTerraformPortefaixStandardFilesRule() *TerraformPortefaixStandardFilesRule {
	return &TerraformPortefaixStandardFilesRule{}
}

// Name returns the rule name
func (ruleule *TerraformPortefaixStandardFilesRule) Name() string {
	return "terraform_portefaix_standard_files"
}

// Enabled returns whether the rule is enabled by default
func (rule *TerraformPortefaixStandardFilesRule) Enabled() bool {
	return false
}

// Severity returns the rule severity
func (rule *TerraformPortefaixStandardFilesRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (rule *TerraformPortefaixStandardFilesRule) Link() string {
	return "https://github.com/portefaix/tflint-ruleset-portefaix/blob/master/README.md"
}

// Check emits errors for any missing files and any block types that are included in the wrong file
func (rule *TerraformPortefaixStandardFilesRule) Check(runner tflint.Runner) error {
	log.Printf("[DEBUG] Check `%s` rule for runner", rule.Name())
	if err := rule.checkProviderFile(runner); err != nil {
		return err
	}

	return nil
}

func (rule *TerraformPortefaixStandardFilesRule) checkProviderFile(runner tflint.Runner) error {
	files, _ := runner.Files()
	log.Printf("[DEBUG] Files: %d", len(files))

	// allowedFiles := map[string]bool{"providers.tf": true, "main.tf": true}

	for name := range files {
		_, filename := path.Split(name)
		log.Printf("[DEBUG] OK: %s %s %s", name, filename, path.Dir(name))
		if filename == filenameProvider {
			return nil
		}
	}
	message := fmt.Sprintf("Module must include a %s file as the provider configuration.", filenameProvider)
	return runner.EmitIssue(rule, message, hcl.Range{Start: hcl.InitialPos})
}