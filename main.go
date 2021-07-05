// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"

	"github.com/portefaix/tflint-ruleset-portefaix/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "portefaix",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewTerraformPortefaixStandardStructureRule(),
				rules.NewTerraformPortefaixStandardFilesRule(),
			},
		},
	})
}
