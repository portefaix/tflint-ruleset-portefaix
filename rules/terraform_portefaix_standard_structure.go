package rules

import (
	"fmt"
	"log"
	"path"

	"github.com/hashicorp/go-version"
	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

const (
	backendVarsDirectory = "backend-vars"
	tfVarsDirectory      = "tfvars"
)

// TerraformPortefaixStandardStructureRule checks whether modules adhere to Terraform Portefaix standard component structure
type TerraformPortefaixStandardStructureRule struct{}

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
func (rule *TerraformPortefaixStandardStructureRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (rule *TerraformPortefaixStandardStructureRule) Link() string {
	return "https://github.com/nlamirault/tflint-ruleset-portefaix/blob/master/README.md"
}

// Check emits errors for any missing files and any block types that are included in the wrong file
func (rule *TerraformPortefaixStandardStructureRule) Check(runner tflint.Runner) error {
	log.Printf("[ERR] Check `%s` rule for runner", rule.Name())
	if err := rule.checkDirectories(runner); err != nil {
		return err
	}
	// if err := rule.checkExternalModule(runner); err != nil {
	// 	return err
	// }

	return nil
}

func (rule *TerraformPortefaixStandardStructureRule) checkExternalModule(runner tflint.Runner) error {
	return runner.WalkModuleCalls(func(call *configs.ModuleCall) error {
		log.Printf("[INFO] Source: %v", call.SourceAddr)
		if call.SourceAddr != "terraform-google-modules" {
			return runner.EmitIssue(rule, "unacceptable module source", call.SourceAddrRange)
		}

		if len(call.Providers) != 0 {
			return runner.EmitIssue(rule, "must not pass providers", hcl.Range{})
		}

		expectedVersion, _ := version.NewVersion("0.1.0")
		if !call.Version.Required.Check(expectedVersion) {
			return runner.EmitIssue(rule, "must accept version 0.1.0", call.Version.DeclRange)
		}

		return nil
	})
}

func (rule *TerraformPortefaixStandardStructureRule) checkDirectories(runner tflint.Runner) error {
	files, _ := runner.Files()
	log.Printf("[INFO] Files: %d", len(files))
	allowedFiles := map[string]bool{"providers.tf": true, "main.tf": true}

	for name := range files {
		_, filename := path.Split(name)
		log.Printf("[INFO] OK: %s %s", name, filename)
		if _, exists := allowedFiles[filename]; !exists {
			message := fmt.Sprintf("File %s is not allowed here.", filename)

			return runner.EmitIssue(rule, message, hcl.Range{Start: hcl.InitialPos})
		}
	}

	// f := runner.Files()
	// directories := make(map[string]*hcl.File, len(f))
	// for name, file := range f {
	// 	directories[filepath.Base(name)] = file
	// }

	// log.Printf("[DEBUG] %d directories found: %v", len(directories), directories)

	// if directories[backendVarsDirectory] == nil {
	// 	runner.EmitIssue(
	// 		r,
	// 		fmt.Sprintf("Structure should include a %s directory for the backend configuration", backendVarsDirectory),
	// 		hcl.Range{
	// 			Filename: filepath.Join(config.Module.SourceDir, backendVarsDirectory),
	// 			Start:    hcl.InitialPos,
	// 		},
	// 	)
	// }

	// if directories[tfVarsDirectory] == nil {
	// 	runner.EmitIssue(
	// 		r,
	// 		fmt.Sprintf("Structure should include a %s directory for the configuration", tfVarsDirectory),
	// 		hcl.Range{
	// 			Filename: filepath.Join(config.Module.SourceDir, tfVarsDirectory),
	// 			Start:    hcl.InitialPos,
	// 		},
	// 	)
	// }

	return nil
}
