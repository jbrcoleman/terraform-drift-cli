# terraform-drift-cli
A cli tool used to detect drift in infrastructure deployed through terraform.
# tfdrift

A fast, colour-coded CLI for visualising Terraform plan changes at a glance.

```
  + create  (2)
    + create  aws_instance.web
    + create  module.vpc.aws_subnet.public

  ~ update  (1)
    ~ update  aws_security_group.web_sg

  ± replace  (1)
    ± replace  module.vpc.aws_subnet.private
               module: module.vpc

  - destroy  (1)
    - destroy  aws_s3_bucket.logs

Plan: 2 to add, 1 to change, 1 to replace, 1 to destroy.
```

## Installation

```bash
go install github.com/yourusername/tfdrift@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/tfdrift
cd tfdrift
go build -o tfdrift .
```

## Usage

### Read from a saved plan JSON file
```bash
terraform plan -json -out=plan.bin
terraform show -json plan.bin > plan.json
tfdrift plan -f plan.json
```

### Pipe directly from terraform
```bash
terraform plan -json | tfdrift plan
```

### Run terraform plan automatically
```bash
# Must be run inside a Terraform root module directory
tfdrift plan --run
```

### Show a one-line summary only
```bash
tfdrift plan -f plan.json --summary
```

### Filter by resource type
```bash
tfdrift plan -f plan.json --filter aws_security_group
```

### Output as JSON (useful in CI pipelines)
```bash
tfdrift plan -f plan.json --json | jq '.[] | select(.action == "delete")'
```

## Project Structure

```
tfdrift/
├── main.go
├── cmd/
│   ├── root.go       # cobra root command
│   └── plan.go       # plan subcommand + flags
├── internal/
│   ├── parser/
│   │   ├── parser.go       # TF plan JSON structs + extraction logic
│   │   └── parser_test.go
│   └── output/
│       └── output.go       # colour-coded terminal output
└── testdata/
    └── example_plan.json   # sample plan for local testing
```

## Development

```bash
# Run tests
go test ./...

# Try it with the example plan
go run . plan -f testdata/example_plan.json

# Summary only
go run . plan -f testdata/example_plan.json --summary

# JSON output
go run . plan -f testdata/example_plan.json --json
```

## Ideas for Extension

- `--diff` flag to show attribute-level before/after diffs for updated resources
- `--module` flag to filter by module address
- Slack/webhook notification support
- Compare two plan JSON files to detect changes between runs
- GitHub Actions integration (post plan summary as PR comment)