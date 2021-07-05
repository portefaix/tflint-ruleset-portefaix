# TFLint Ruleset Portefaix

## Requirements

- TFLint v0.24+
- Go v1.16

## Installation

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "portefaix" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/terraform-linters/tflint-ruleset-portefaix"
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|terraform_portefaix_standard_structure|Check module respect Skale-5 directories recommendations|ERROR|✔||
|terraform_portefaix_standard_files|Check module respect Portefaix files recommendations|ERROR|✔||
|terraform_portefaix_tfenv|Check module use a file for Tfenv|ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make build
```

You can easily install the built plugin with the following:

```
$ make install
```